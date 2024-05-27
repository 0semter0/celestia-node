package bitswap

import (
	"context"
	"fmt"
	"testing"

	"github.com/ipfs/boxo/bitswap"
	"github.com/ipfs/boxo/bitswap/network"
	"github.com/ipfs/boxo/blockstore"
	"github.com/ipfs/boxo/exchange"
	"github.com/ipfs/boxo/routing/offline"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	ds "github.com/ipfs/go-datastore"
	dssync "github.com/ipfs/go-datastore/sync"
	record "github.com/libp2p/go-libp2p-record"
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
	"github.com/stretchr/testify/require"

	"github.com/celestiaorg/rsmt2d"
)

func remoteClient(ctx context.Context, t *testing.T, bstore blockstore.Blockstore) exchange.Fetcher {
	net, err := mocknet.FullMeshLinked(2)
	require.NoError(t, err)

	dstore := dssync.MutexWrap(ds.NewMapDatastore())
	routing := offline.NewOfflineRouter(dstore, record.NamespacedValidator{})
	_ = bitswap.New(
		ctx,
		network.NewFromIpfsHost(net.Hosts()[0], routing),
		bstore,
	)

	dstoreClient := dssync.MutexWrap(ds.NewMapDatastore())
	bstoreClient := blockstore.NewBlockstore(dstoreClient)
	routingClient := offline.NewOfflineRouter(dstoreClient, record.NamespacedValidator{})

	bitswapClient := bitswap.New(
		ctx,
		network.NewFromIpfsHost(net.Hosts()[1], routingClient),
		bstoreClient,
	)

	err = net.ConnectAllButSelf()
	require.NoError(t, err)

	return bitswapClient
}

type testBlockstore struct {
	eds *rsmt2d.ExtendedDataSquare
}

func newTestBlockstore(eds *rsmt2d.ExtendedDataSquare) *testBlockstore {
	return &testBlockstore{eds: eds}
}

func (b *testBlockstore) Get(_ context.Context, cid cid.Cid) (blocks.Block, error) {
	spec, ok := specRegistry[cid.Prefix().MhType]
	if !ok {
		return nil, fmt.Errorf("unsupported codec")
	}

	bldr, err := spec.builder(cid)
	if err != nil {
		return nil, err
	}

	return bldr.BlockFromEDS(b.eds)
}

func (b *testBlockstore) GetSize(ctx context.Context, cid cid.Cid) (int, error) {
	blk, err := b.Get(ctx, cid)
	if err != nil {
		return 0, err
	}
	return len(blk.RawData()), nil
}

func (b *testBlockstore) Has(ctx context.Context, cid cid.Cid) (bool, error) {
	_, err := b.Get(ctx, cid)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (b *testBlockstore) Put(context.Context, blocks.Block) error {
	panic("not implemented")
}

func (b *testBlockstore) PutMany(context.Context, []blocks.Block) error {
	panic("not implemented")
}

func (b *testBlockstore) DeleteBlock(context.Context, cid.Cid) error {
	panic("not implemented")
}

func (b *testBlockstore) AllKeysChan(context.Context) (<-chan cid.Cid, error) {
	panic("not implemented")
}

func (b *testBlockstore) HashOnRead(bool) {
	panic("not implemented")
}
