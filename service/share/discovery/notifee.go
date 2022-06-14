package discovery

import (
	"context"
	"errors"
	"time"

	logging "github.com/ipfs/go-log/v2"

	"github.com/libp2p/go-libp2p-core/event"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
)

const (
	// peersLimit is max amount of peers that will be discovered
	peersLimit = 5
	// peerWeight total weight of discovered peers
	peerWeight = 1000
	// connectTimeout is timeout given to connect to discovered peer
	connectTimeout = time.Minute
)

var log = logging.Logger("discovery")

// Notifee allows to receive and store discovered peers.
type Notifee struct {
	set  *peer.Set
	host host.Host
}

// NewNotifee constructs new Notifee.
func NewNotifee(cache *peer.Set, h host.Host) *Notifee {
	return &Notifee{
		cache,
		h,
	}
}

// HandlePeersFound receives peers and tries to establish a connection with them.
// Peer will be added to PeerCache if connection succeeds.
func (n *Notifee) HandlePeersFound(topic string, peers []peer.AddrInfo) error {
	for _, peer := range peers {
		if n.set.Size() == peersLimit {
			return errors.New("amount of peers reaches the limit")
		}

		if peer.ID == n.host.ID() || len(peer.Addrs) == 0 || n.set.Contains(peer.ID) {
			continue
		}
		if n.host.Network().Connectedness(peer.ID) != network.Connected {
			ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
			err := n.host.Connect(ctx, peer)
			if err != nil {
				cancel()
				log.Warn(err)
				continue
			}
			cancel()
		}
		log.Debugw("adding peer to cache", "id", peer.ID)
		n.host.ConnManager().TagPeer(peer.ID, topic, peerWeight)
		n.set.Add(peer.ID)
		go n.emit(peer.ID, network.Connected)
	}

	return nil
}

func (n *Notifee) emit(id peer.ID, state network.Connectedness) {
	emitter, err := n.host.EventBus().Emitter(&event.EvtPeerConnectednessChanged{})
	if err != nil {
		log.Warn(err)
		return
	}
	err = emitter.Emit(event.EvtPeerConnectednessChanged{Peer: id, Connectedness: state})
	if err != nil {
		log.Warn(err)
	}
}
