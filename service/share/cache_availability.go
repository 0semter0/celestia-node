package share

import (
	"bytes"
	"context"
	"sync"

	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/autobatch"
	"github.com/ipfs/go-datastore/namespace"
	"github.com/tendermint/tendermint/pkg/da"
)

var (
	// DefaultWriteBatchSize defines the size of the batched header write.
	// Headers are written in batches not to thrash the underlying Datastore with writes.
	// TODO(@Wondertan, @renaynay): Those values must be configurable and proper defaults should be set for specific node
	//  type. (#709)
	DefaultWriteBatchSize   = 2048
	cacheAvailabilityPrefix = datastore.NewKey("sampling_result")

	minRoot = da.MinDataAvailabilityHeader()
)

// CacheAvailability wraps a given Availability (whether it's light or full)
// and stores the results of a successful sampling routine over a given Root's hash
// to disk.
type CacheAvailability struct {
	avail Availability

	// TODO(@Wondertan): Once we come to parallelized DASer, this lock becomes a contention point
	//  Related to #483
	dsLk sync.RWMutex
	ds   *autobatch.Datastore
}

// NewCacheAvailability wraps the given Availability with an additional datastore
// for sampling result caching.
func NewCacheAvailability(avail Availability, ds datastore.Batching) *CacheAvailability {
	ds = namespace.Wrap(ds, cacheAvailabilityPrefix)
	autoDS := autobatch.NewAutoBatching(ds, DefaultWriteBatchSize)

	return &CacheAvailability{
		avail: avail,
		ds:    autoDS,
	}
}

// Start starts the underlying Availability.
func (ca *CacheAvailability) Start(ctx context.Context) error {
	return ca.avail.Start(ctx)
}

// Stop is an alias for Close to conform to the Availability interface.
func (ca *CacheAvailability) Stop(ctx context.Context) error {
	return ca.Close(ctx)
}

// SharesAvailable will store, upon success, the hash of the given Root to disk.
func (ca *CacheAvailability) SharesAvailable(ctx context.Context, root *Root) error {
	// short-circuit if the given root is minimum DAH of an empty data square
	if isMinRoot(root) {
		return nil
	}
	// do not sample over Root that has already been sampled
	key := rootKey(root)

	ca.dsLk.RLock()
	exists, err := ca.ds.Has(ctx, key)
	ca.dsLk.RUnlock()
	if err != nil || exists {
		return err
	}

	err = ca.avail.SharesAvailable(ctx, root)
	if err != nil {
		return err
	}

	ca.dsLk.Lock()
	err = ca.ds.Put(ctx, key, []byte{})
	ca.dsLk.Unlock()
	if err != nil {
		log.Errorw("storing root of successful SharesAvailable request to disk", "err", err)
	}
	return err
}

func (ca *CacheAvailability) ProbabilityOfAvailability() float64 {
	return ca.avail.ProbabilityOfAvailability()
}

// Close flushes all queued writes to disk.
func (ca *CacheAvailability) Close(ctx context.Context) error {
	return ca.ds.Flush(ctx)
}

func rootKey(root *Root) datastore.Key {
	return datastore.NewKey(root.String())
}

// isMinRoot returns whether the given root is a minimum (empty)
// DataAvailabilityHeader (DAH).
func isMinRoot(root *Root) bool {
	return bytes.Equal(minRoot.Hash(), root.Hash())
}
