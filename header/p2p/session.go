package p2p

import (
	"context"
	"errors"
	"io"
	"sort"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"

	"github.com/celestiaorg/go-libp2p-messenger/serde"

	"github.com/celestiaorg/celestia-node/header"
	p2p_pb "github.com/celestiaorg/celestia-node/header/p2p/pb"
)

// TODO(@vgonkivs): make it configurable
var (
	// headersPerPeer is a maximum amount of headers that will be requested per peer.
	headersPerPeer uint64 = 64
)

// session aims to divide a range of headers
// into several smaller requests among different peers.
type session struct {
	ctx        context.Context
	cancel     context.CancelFunc
	host       host.Host
	protocolID protocol.ID
	// peerTracker contains discovered peers with records that describes their activity.
	queue *peerQueue

	reqCh chan *p2p_pb.ExtendedHeaderRequest
	errCh chan error
}

func newSession(ctx context.Context, h host.Host, peerTracker []*peerStat, protocolID protocol.ID) *session {
	ctx, cancel := context.WithCancel(ctx)
	return &session{
		ctx:        ctx,
		cancel:     cancel,
		protocolID: protocolID,
		host:       h,
		queue:      newPeerQueue(ctx, peerTracker),
		errCh:      make(chan error),
	}
}

// GetRangeByHeight requests headers from different peers.
func (s *session) getRangeByHeight(ctx context.Context, from, amount uint64) ([]*header.ExtendedHeader, error) {
	log.Debugw("requesting headers", "from", from, "to", from+amount)

	requests := prepareRequests(from, amount, headersPerPeer)
	result := make(chan []*header.ExtendedHeader, len(requests))
	s.reqCh = make(chan *p2p_pb.ExtendedHeaderRequest, len(requests))

	go s.handleOutgoingRequests(ctx, result)
	for _, req := range requests {
		s.reqCh <- req
	}

	headers := make([]*header.ExtendedHeader, 0, amount)
	for i := 0; i < cap(result); i++ {
		select {
		case <-s.ctx.Done():
			return nil, errors.New("header/p2p: exchange is closed")
		case <-ctx.Done():
			return nil, ctx.Err()
		case err := <-s.errCh:
			return nil, err
		case res := <-result:
			headers = append(headers, res...)
		}
	}
	sort.Slice(headers, func(i, j int) bool {
		return headers[i].Height < headers[j].Height
	})
	return headers, nil
}

// close stops the session.
func (s *session) close() {
	if s.cancel != nil {
		s.cancel()
		s.cancel = nil
	}
}

// handleOutgoingRequests pops a peer from the queue and sends a prepared request to the peer.
func (s *session) handleOutgoingRequests(ctx context.Context, result chan []*header.ExtendedHeader) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-s.ctx.Done():
			return
		case req := <-s.reqCh:
			stats := s.queue.waitPop(ctx)
			if stats.peerID == "" {
				return
			}
			go s.doRequest(ctx, stats, req, result)
		}
	}
}

// doRequest chooses the best peer to fetch headers and sends a request in range of available maxRetryAttempts.
func (s *session) doRequest(
	ctx context.Context,
	stat *peerStat,
	req *p2p_pb.ExtendedHeaderRequest,
	headers chan []*header.ExtendedHeader,
) {
	r, size, duration, err := s.requestHeaders(ctx, stat.peerID, req)
	if err != nil {
		if err == context.Canceled || err == context.DeadlineExceeded {
			return
		}
		log.Errorw("requesting headers from peer failed."+
			"Retrying the request from different peer", "failed peer", stat.peerID, "err", err)
		select {
		case <-ctx.Done():
		case <-s.ctx.Done():
		case s.reqCh <- req:
		}
		return
	}

	h, err := s.processResponse(r)
	if err != nil {
		s.errCh <- err
		return
	}
	log.Debugw("request headers from peer succeed", "from", s.host.ID(), "pID", stat.peerID, "amount", req.Amount)
	// send headers to the channel, update peer stats and return peer to the queue, so it can be re-used in case
	// if there are other requests awaiting
	headers <- h
	stat.updateStats(size, duration)
	s.queue.push(stat)
}

// requestHeaders sends the ExtendedHeaderRequest to a remote peer.
func (s *session) requestHeaders(
	ctx context.Context,
	to peer.ID,
	req *p2p_pb.ExtendedHeaderRequest,
) ([]*p2p_pb.ExtendedHeaderResponse, uint64, uint64, error) {
	startTime := time.Now()
	stream, err := s.host.NewStream(ctx, to, s.protocolID)
	if err != nil {
		return nil, 0, 0, err
	}
	// send request
	_, err = serde.Write(stream, req)
	if err != nil {
		stream.Reset() //nolint:errcheck
		return nil, 0, 0, err
	}
	err = stream.CloseWrite()
	if err != nil {
		log.Error(err)
	}
	headers := make([]*p2p_pb.ExtendedHeaderResponse, 0)
	totalRequestSize := uint64(0)
	for i := 0; i < int(req.Amount); i++ {
		resp := new(p2p_pb.ExtendedHeaderResponse)
		msgSize, err := serde.Read(stream, resp)
		if err != nil {
			if err == io.EOF {
				break
			}
			stream.Reset() //nolint:errcheck
			return nil, 0, 0, err
		}
		totalRequestSize += uint64(msgSize)
		headers = append(headers, resp)
	}
	duration := time.Since(startTime).Milliseconds()
	if err = stream.Close(); err != nil {
		log.Errorw("closing stream", "err", err)
	}
	return headers, totalRequestSize, uint64(duration), nil
}

// processResponse converts ExtendedHeaderResponse to ExtendedHeader.
func (s *session) processResponse(responses []*p2p_pb.ExtendedHeaderResponse) ([]*header.ExtendedHeader, error) {
	headers := make([]*header.ExtendedHeader, 0)
	for _, resp := range responses {
		err := convertStatusCodeToError(resp.StatusCode)
		if err != nil {
			return nil, err
		}
		header, err := header.UnmarshalExtendedHeader(resp.Body)
		if err != nil {
			return nil, err
		}
		headers = append(headers, header)
	}

	return headers, nil
}

// prepareRequests converts incoming range into separate ExtendedHeaderRequest.
func prepareRequests(from, amount, headersPerPeer uint64) []*p2p_pb.ExtendedHeaderRequest {
	requests := make([]*p2p_pb.ExtendedHeaderRequest, 0, amount/headersPerPeer)
	for amount > uint64(0) {
		var requestSize uint64
		request := &p2p_pb.ExtendedHeaderRequest{
			Data: &p2p_pb.ExtendedHeaderRequest_Origin{Origin: from},
		}
		if amount < headersPerPeer {
			requestSize = amount
			amount = 0
		} else {
			amount -= headersPerPeer
			from += headersPerPeer
			requestSize = headersPerPeer
		}
		request.Amount = requestSize
		requests = append(requests, request)
	}
	return requests
}
