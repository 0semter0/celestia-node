package fraud

import (
	"context"
	"time"

	"github.com/ipfs/go-datastore"
	"github.com/libp2p/go-libp2p-core/event"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"

	"github.com/celestiaorg/go-libp2p-messenger/serde"

	pb "github.com/celestiaorg/celestia-node/fraud/pb"
)

func (f *ProofService) syncFraudProofs(ctx context.Context) {
	log.Debug("start fetching Fraud Proofs")
	sub, err := f.host.EventBus().Subscribe(&event.EvtPeerIdentificationCompleted{})
	if err != nil {
		log.Error(err)
		return
	}
	defer sub.Close()
	f.topicsLk.RLock()
	proofTypes := make([]uint32, 0, len(f.topics))
	for proofType := range f.topics {
		proofTypes = append(proofTypes, uint32(proofType))
	}
	f.topicsLk.RUnlock()
	connStatus := event.EvtPeerIdentificationCompleted{}
	for i := 0; i < fraudRequests; i++ {
		select {
		case <-ctx.Done():
			return
		case e := <-sub.Out():
			connStatus = e.(event.EvtPeerIdentificationCompleted)
		}

		if connStatus.Peer == f.host.ID() {
			i--
			continue
		}
		go func(pid peer.ID) {
			log.Debugw("requesting proofs from peer", "pid", pid)
			respProofs, err := requestProofs(ctx, f.host, pid, proofTypes)
			if err != nil {
				log.Errorw("error while requesting fraud proofs", "err", err, "peer", pid)
				return
			}
			if len(respProofs) == 0 {
				log.Debugw("peer did not return any proofs", "pid", pid)
				return
			}
			log.Debugw("got fraud proofs from peer", "pid", connStatus.Peer)
			for _, data := range respProofs {
				if err != nil {
					log.Warn(err)
					continue
				}
				f.topicsLk.RLock()
				topic, ok := f.topics[ProofType(data.Type)]
				f.topicsLk.RUnlock()
				if !ok {
					log.Errorf("topic for %s does not exist", ProofType(data.Type))
					continue
				}
				for _, val := range data.Value {
					err = topic.Publish(
						ctx,
						val,
						// broadcast across all local subscriptions in order to verify fraud proof and to stop services
						pubsub.WithLocalPublication(true),
						// key can be nil because it will not be verified in this case
						pubsub.WithSecretKeyAndPeerId(nil, pid),
					)
					if err != nil {
						log.Error(err)
					}
				}
			}
		}(connStatus.Peer)
	}
}

func (f *ProofService) handleFraudMessageRequest(stream network.Stream) {
	req := &pb.FraudMessageRequest{}
	if err := stream.SetReadDeadline(time.Now().Add(readDeadline)); err != nil {
		log.Warn(err)
	}
	_, err := serde.Read(stream, req)
	if err != nil {
		stream.Reset() //nolint:errcheck
		log.Errorw("handling fraud message request failed", "err", err)
		return
	}
	if err = stream.CloseRead(); err != nil {
		log.Warn(err)
	}

	resp := &pb.FraudMessageResponse{}
	resp.Proofs = make([]*pb.ProofResponse, 0, len(req.RequestedProofType))
	for _, p := range req.RequestedProofType {
		proofs, err := f.Get(f.ctx, ProofType(p))
		if err != nil {
			if err != datastore.ErrNotFound {
				log.Error(err)
			}
			continue
		}
		pbProofs := &pb.ProofResponse{Type: p, Value: make([][]byte, 0, len(proofs))}
		for _, proof := range proofs {
			bin, err := proof.MarshalBinary()
			if err != nil {
				log.Error(err)
				continue
			}
			pbProofs.Value = append(pbProofs.Value, bin)
		}
		resp.Proofs = append(resp.Proofs, pbProofs)
	}

	if err = stream.SetWriteDeadline(time.Now().Add(writeDeadline)); err != nil {
		log.Warn(err)
	}
	_, err = serde.Write(stream, resp)
	if err != nil {
		stream.Reset() //nolint:errcheck
		log.Errorw("error while writing a response", "err", err)
		return
	}
	if err = stream.CloseWrite(); err != nil {
		log.Warnw("error while closing a writer in stream", "err", err)
	}
}
