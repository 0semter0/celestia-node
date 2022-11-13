package eds

import (
	"context"
	"errors"
	"sync"

	"github.com/filecoin-project/dagstore"
	"github.com/filecoin-project/dagstore/shard"
	lru "github.com/hashicorp/golang-lru"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	ipld "github.com/ipfs/go-ipld-format"
)

var _ blockstore.Blockstore = (*Blockstore)(nil)

var (
	maxCacheSize            = 100
	ErrUnsupportedOperation = errors.New("unsupported operation")
	ErrInvalidShardCount    = errors.New("the provided cid does not map to exactly one shard")
)

type accessorWithBlockstore struct {
	sa *dagstore.ShardAccessor
	bs dagstore.ReadBlockstore
}

// Blockstore implements the blockstore.Blockstore interface on an EDSStore.
// The lru cache approach is heavily inspired by the open PR filecoin-project/dagstore/116.
// The main differences to the implementation here are that we do not support multiple shards per key,
// call GetSize directly on the underlying RO blockstore, and do not throw errors on Put/PutMany.
// Also, we do not abstract away the blockstore operations.
type Blockstore struct {
	store *EDSStore

	// bsStripedLocks prevents simultaneous RW access to the blockstore cache for a shard. Instead of using only one
	// lock or one lock per key, we stripe the shard keys across 256 locks
	bsStripedLocks [256]sync.Mutex
	// caches the blockstore for a given shard for shard read affinity i.e.
	// further reads will likely be from the same shard. Maps (shard key -> blockstore).
	blockstoreCache *lru.Cache
}

func NewEDSBlockstore(s *EDSStore) (*Blockstore, error) {
	// instantiate the blockstore cache
	bslru, err := lru.NewWithEvict(maxCacheSize, func(_ interface{}, val interface{}) {
		// ensure we close the blockstore for a shard when it's evicted from the cache so dagstore can gc it.
		abs, ok := val.(*accessorWithBlockstore)
		if ok {
			abs.sa.Close()
		}
	})
	if err != nil {
		return nil, err
	}
	return &Blockstore{
		store:           s,
		blockstoreCache: bslru,
	}, nil
}

func (bs *Blockstore) DeleteBlock(ctx context.Context, cid cid.Cid) error {
	// TODO: Should we delete the CAR containing this cid, or should it remain a noop?
	return ErrUnsupportedOperation
}

func (bs *Blockstore) Has(ctx context.Context, cid cid.Cid) (bool, error) {
	keys, err := bs.store.dgstr.ShardsContainingMultihash(ctx, cid.Hash())
	if err != nil {
		return false, err
	}
	return len(keys) > 0, nil
}

func (bs *Blockstore) Get(ctx context.Context, cid cid.Cid) (blocks.Block, error) {
	blockstr, err := bs.getReadOnlyBlockstore(ctx, cid)
	if err != nil {
		return nil, ipld.ErrNotFound{Cid: cid}
	}
	// TODO: if bs.Get returns an error and it is from the cache, we should remove it from the cache
	return blockstr.Get(ctx, cid)
}

func (bs *Blockstore) GetSize(ctx context.Context, cid cid.Cid) (int, error) {
	blockstr, err := bs.getReadOnlyBlockstore(ctx, cid)
	if err != nil {
		return 0, err
	}
	return blockstr.GetSize(ctx, cid)
}

// Put is a noop on the EDS blockstore, but it does not return an error because it is called by bitswap. For
// clarification, an implementation of Put does not make sense in this context because it is unclear which CAR file the
// block should be written to.
func (bs *Blockstore) Put(ctx context.Context, block blocks.Block) error {
	return nil
}

// PutMany is a noop on the EDS blockstore, but it does not return an error because it is called by bitswap. For
// clarification, an implementation of PutMany does not make sense in this context because it is unclear which CAR file
// the blocks should be written to.
func (bs *Blockstore) PutMany(ctx context.Context, blocks []blocks.Block) error {
	return nil
}

// AllKeysChan is a noop on the EDS blockstore because the keys are not stored in a single CAR file.
func (bs *Blockstore) AllKeysChan(ctx context.Context) (<-chan cid.Cid, error) {
	return nil, ErrUnsupportedOperation
}

// HashOnRead is a noop on the EDS blockstore but an error cannot be returned due to the method signature from the
// blockstore interface.
func (bs *Blockstore) HashOnRead(bool) {
	log.Warnf("HashOnRead is a noop on the EDS blockstore")
}

func (bs *Blockstore) getReadOnlyBlockstore(ctx context.Context, cid cid.Cid) (dagstore.ReadBlockstore, error) {
	keys, err := bs.store.dgstr.ShardsContainingMultihash(ctx, cid.Hash())
	if err != nil {
		return nil, err
	}
	if len(keys) != 1 {
		return nil, ErrInvalidShardCount
	}

	// try to fetch from cache
	shardKey := keys[0]
	blockstr, err := bs.readFromBSCache(shardKey)
	if err == nil && blockstr != nil {
		return blockstr, nil
	}

	// wasn't found in cache, so acquire it and add to cache
	ch := make(chan dagstore.ShardResult, 1)
	err = bs.store.dgstr.AcquireShard(ctx, shardKey, ch, dagstore.AcquireOpts{})
	if err != nil {
		return nil, err
	}
	result := <-ch
	if result.Error != nil {
		return nil, result.Error
	}

	return bs.addToBSCache(shardKey, result.Accessor)
}

func (bs *Blockstore) readFromBSCache(shardContainingCid shard.Key) (dagstore.ReadBlockstore, error) {
	lk := &bs.bsStripedLocks[shardKeyToStriped(shardContainingCid)]
	lk.Lock()
	defer lk.Unlock()

	// We've already ensured that the given shard has the cid/multihash we are looking for.
	val, ok := bs.blockstoreCache.Get(shardContainingCid)
	if !ok {
		return nil, errors.New("not found in cache")
	}

	rbs := val.(*accessorWithBlockstore).bs
	return rbs, nil
}

func (bs *Blockstore) addToBSCache(
	shardContainingCid shard.Key,
	accessor *dagstore.ShardAccessor,
) (dagstore.ReadBlockstore, error) {
	lk := &bs.bsStripedLocks[shardKeyToStriped(shardContainingCid)]
	lk.Lock()
	defer lk.Unlock()

	blockStore, err := accessor.Blockstore()
	if err != nil {
		return nil, err
	}

	bs.blockstoreCache.Add(shardContainingCid, &accessorWithBlockstore{
		bs: blockStore,
		sa: accessor,
	})
	return blockStore, nil
}

// shardKeyToStriped returns the index of the lock to use for a given shard key. We use the last byte of the
// shard key as the pseudo-random index.
func shardKeyToStriped(sk shard.Key) byte {
	return sk.String()[len(sk.String())-1]
}
