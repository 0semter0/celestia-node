package core

import (
	"context"
	"testing"
	"time"

	mdutils "github.com/ipfs/go-merkledag/test"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/event"
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/celestiaorg/celestia-node/header"
	"github.com/celestiaorg/celestia-node/libs/header/p2p"
	"github.com/celestiaorg/celestia-node/share/p2p/shrexsub"
)

// TestListener tests the lifecycle of the core listener.
func TestListener(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	t.Cleanup(cancel)

	// create mocknet with two pubsub endpoints
	ps0, ps1 := createMocknetWithTwoPubsubEndpoints(ctx, t)
	subscriber := p2p.NewSubscriber[*header.ExtendedHeader]("private", ps1, header.MsgID)
	err := subscriber.AddValidator(func(context.Context, *header.ExtendedHeader) pubsub.ValidationResult {
		return pubsub.ValidationAccept
	})
	require.NoError(t, err)
	require.NoError(t, subscriber.Start(ctx))
	subs, err := subscriber.Subscribe()
	require.NoError(t, err)

	// create one block to store as Head in local store and then unsubscribe from block events
	fetcher := createCoreFetcher(t)
	eds := createEdsPubSub(ctx, t)
	// create Listener and start listening
	cl := createListener(ctx, t, fetcher, ps0, eds)
	err = cl.Start(ctx)
	require.NoError(t, err)

	edsSubs, err := eds.Subscribe()
	require.NoError(t, err)
	defer edsSubs.Cancel()

	// ensure headers and dataHash are getting broadcasted to the relevant topics
	for i := 1; i < 6; i++ {
		h, err := subs.NextHeader(ctx)
		require.NoError(t, err)

		dataHash, err := edsSubs.Next(ctx)
		require.NoError(t, err)

		require.Equal(t, h.DataHash.Bytes(), []byte(dataHash))
	}

	err = cl.Stop(ctx)
	require.NoError(t, err)
	require.Nil(t, cl.cancel)
}

func createMocknetWithTwoPubsubEndpoints(ctx context.Context, t *testing.T) (*pubsub.PubSub, *pubsub.PubSub) {
	net, err := mocknet.FullMeshLinked(2)
	require.NoError(t, err)
	host0, host1 := net.Hosts()[0], net.Hosts()[1]

	// create pubsub for host
	ps0, err := pubsub.NewGossipSub(context.Background(), host0,
		pubsub.WithMessageSignaturePolicy(pubsub.StrictNoSign))
	require.NoError(t, err)
	// create pubsub for peer-side (to test broadcast comes through network)
	ps1, err := pubsub.NewGossipSub(context.Background(), host1,
		pubsub.WithMessageSignaturePolicy(pubsub.StrictNoSign))
	require.NoError(t, err)

	sub0, err := host0.EventBus().Subscribe(&event.EvtPeerIdentificationCompleted{})
	require.NoError(t, err)
	sub1, err := host1.EventBus().Subscribe(&event.EvtPeerIdentificationCompleted{})
	require.NoError(t, err)

	err = net.ConnectAllButSelf()
	require.NoError(t, err)

	// wait on both peer identification events
	for i := 0; i < 2; i++ {
		select {
		case <-sub0.Out():
		case <-sub1.Out():
		case <-ctx.Done():
			assert.FailNow(t, "timeout waiting for peers to connect")
		}
	}

	return ps0, ps1
}

func createListener(
	ctx context.Context,
	t *testing.T,
	fetcher *BlockFetcher,
	ps *pubsub.PubSub,
	edsSub *shrexsub.PubSub,
) *Listener {
	p2pSub := p2p.NewSubscriber[*header.ExtendedHeader]("private", ps, header.MsgID)
	err := p2pSub.Start(ctx)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, p2pSub.Stop(ctx))
	})

	return NewListener(p2pSub, fetcher, edsSub.Broadcast, mdutils.Bserv(), header.MakeExtendedHeader)
}

func createEdsPubSub(ctx context.Context, t *testing.T) *shrexsub.PubSub {
	net, err := mocknet.FullMeshLinked(1)
	require.NoError(t, err)
	edsSub, err := shrexsub.NewPubSub(ctx, net.Hosts()[0], "eds-test")
	require.NoError(t, err)
	require.NoError(t, edsSub.Start(ctx))
	t.Cleanup(func() {
		require.NoError(t, edsSub.Stop(ctx))
	})
	return edsSub
}
