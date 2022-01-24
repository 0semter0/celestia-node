package header

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"

	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

// Syncer implements efficient synchronization for headers.
type Syncer struct {
	sub      *P2PSubscriber
	exchange Exchange
	store    Store
	trusted  tmbytes.HexBytes

	// inProgress is set to 1 once syncing commences and
	// is set to 0 once syncing is either finished or
	// not currently in progress
	inProgress uint64
	// signals to start syncing
	triggerSync chan struct{}
	// pending keeps ranges of valid headers rcvd from the network awaiting to be appended to store
	pending ranges

	cancel context.CancelFunc
}

// NewSyncer creates a new instance of Syncer.
func NewSyncer(exchange Exchange, store Store, sub *P2PSubscriber, trusted tmbytes.HexBytes) *Syncer {
	return &Syncer{
		sub:         sub,
		exchange:    exchange,
		store:       store,
		trusted:     trusted,
		triggerSync: make(chan struct{}, 1), // should be buffered
	}
}

// Start starts the syncing routine.
func (s *Syncer) Start(ctx context.Context) error {
	if s.cancel != nil {
		return fmt.Errorf("header: Syncer already started")
	}

	if s.sub != nil {
		err := s.sub.AddValidator(s.validateMsg)
		if err != nil {
			return err
		}
	}

	// TODO(@Wondertan): Ideally, this initialization should be part of Init process
	err := s.initStore(ctx)
	if err != nil {
		log.Error(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	go s.syncLoop(ctx)
	s.wantSync()
	s.cancel = cancel
	return nil
}

// Stop stops Syncer.
func (s *Syncer) Stop(context.Context) error {
	s.cancel()
	s.cancel = nil
	return nil
}

// IsSyncing returns the current sync status of the Syncer.
func (s *Syncer) IsSyncing() bool {
	return atomic.LoadUint64(&s.inProgress) == 1
}

// init initializes if it's empty
func (s *Syncer) initStore(ctx context.Context) error {
	_, err := s.store.Head(ctx)
	switch err {
	case ErrNoHead:
		// if there is no head - request header at trusted hash.
		trusted, err := s.exchange.RequestByHash(ctx, s.trusted)
		if err != nil {
			return fmt.Errorf("header: requesting header at trusted hash during init: %w", err)
		}

		err = s.store.Append(ctx, trusted)
		if err != nil {
			return fmt.Errorf("header: appending header during init: %w", err)
		}
	case nil:
	}

	return nil
}

// trustedHead returns the highest known trusted header which is a valid header with time within the trusting period.
func (s *Syncer) trustedHead(ctx context.Context) (*ExtendedHeader, error) {
	// check pending for trusted header and return it if applicable
	// NOTE: Pending cannot be expired, guaranteed
	pendHead := s.pending.Head()
	if pendHead != nil {
		return pendHead, nil
	}

	sbj, err := s.store.Head(ctx)
	if err != nil {
		return nil, err
	}

	// check if our subjective header is not expired and use it
	if !sbj.IsExpired() {
		return sbj, nil
	}

	// otherwise, request head from a trustedPeer or, in other words, do automatic subjective initialization
	objHead, err := s.exchange.RequestHead(ctx)
	if err != nil {
		return nil, err
	}

	if objHead.Height > sbj.Height {
		s.pending.Add(objHead)
	}
	return objHead, nil
}

// wantSync will trigger the syncing loop (non-blocking).
func (s *Syncer) wantSync() {
	select {
	case s.triggerSync <- struct{}{}:
	default:
	}
}

// syncLoop controls syncing process.
func (s *Syncer) syncLoop(ctx context.Context) {
	for {
		select {
		case <-s.triggerSync:
			s.sync(ctx)
		case <-ctx.Done():
			return
		}
	}
}

// sync ensures we are synced up to any trusted header.
func (s *Syncer) sync(ctx context.Context) {
	// indicate syncing
	atomic.StoreUint64(&s.inProgress, 1)
	// indicate syncing is stopped
	defer atomic.StoreUint64(&s.inProgress, 0)

	trstHead, err := s.trustedHead(ctx)
	if err != nil {
		log.Errorw("getting trusted head", "err", err)
		return
	}

	err = s.syncTo(ctx, trstHead)
	if err != nil {
		log.Errorw("syncing headers", "err", err)
		return
	}
}

// validateMsg implements validation of incoming Headers as PubSub msg validator.
func (s *Syncer) validateMsg(ctx context.Context, p peer.ID, msg *pubsub.Message) pubsub.ValidationResult {
	maybeHead, err := UnmarshalExtendedHeader(msg.Data)
	if err != nil {
		log.Errorw("unmarshalling header",
			"from", p.ShortString(),
			"err", err)
		return pubsub.ValidationReject
	}

	return s.incoming(ctx, p, maybeHead)
}

// incoming process new incoming Headers, validates them and applies/caches if applicable.
func (s *Syncer) incoming(ctx context.Context, p peer.ID, maybeHead *ExtendedHeader) pubsub.ValidationResult {
	// 1. Try to append. If header is not adjacent/from future - try it for pending cache below
	err := s.store.Append(ctx, maybeHead)
	if err == nil {
		return pubsub.ValidationAccept
	} else if err != ErrNonAdjacent {
		var verErr *VerifyError
		if errors.As(err, &verErr) {
			log.Errorw("invalid header",
				"height", maybeHead.Height,
				"hash", maybeHead.Hash(),
				"from", p.ShortString(),
				"reason", verErr.Reason)
			return pubsub.ValidationReject
		}

		log.Errorw("appending header",
			"height", maybeHead.Height,
			"hash", maybeHead.Hash().String(),
			"from", p.ShortString())
		// might be a storage error or something else, so ignore
		return pubsub.ValidationIgnore
	}

	// 2. Get known trusted head pending to be synced
	trstHead, err := s.trustedHead(ctx)
	if err != nil {
		log.Errorw("getting trusted head", "err", err)
		return pubsub.ValidationIgnore // we don't know if header is invalid so ignore
	}

	// 3. Filter out maybeHead if behind trusted
	if maybeHead.Height <= trstHead.Height {
		log.Warnw("received known header",
			"height", maybeHead.Height,
			"hash", maybeHead.Hash(),
			"from", p.ShortString())

		// TODO(@Wondertan): Remove once duplicates are fully fixed
		log.Warnf("Ignore the warn above - there is a known issue with duplicate headers on the network.")
		return pubsub.ValidationIgnore // we don't know if header is invalid so ignore
	}

	// 4. Verify maybeHead against trusted
	err = trstHead.VerifyNonAdjacent(maybeHead)
	var verErr *VerifyError
	if errors.As(err, &verErr) {
		log.Errorw("invalid header",
			"height", maybeHead.Height,
			"hash", maybeHead.Hash(),
			"from", p.ShortString(),
			"reason", verErr.Reason)
		return pubsub.ValidationReject
	}

	// 5. Save verified header to pending cache
	// NOTE: Pending cache can't be DOSed as we verify above each header against a trusted one.
	s.pending.Add(maybeHead)
	// and trigger sync to catch-up
	s.wantSync()
	log.Infow("new pending head",
		"height", maybeHead.Height,
		"hash", maybeHead.Hash(),
		"from", p.ShortString())
	return pubsub.ValidationAccept
}

// TODO(@Wondertan): Number of headers that can be requested at once. Either make this configurable or,
//  find a proper rationale for constant.
var requestSize uint64 = 512

// syncTo requests headers from locally stored head up to the new head.
func (s *Syncer) syncTo(ctx context.Context, newHead *ExtendedHeader) error {
	head, err := s.store.Head(ctx)
	if err != nil {
		return err
	}

	if head.Height == newHead.Height {
		return nil
	}

	log.Infow("syncing headers", "from", head.Height, "to", newHead.Height)
	defer log.Info("synced headers")

	start, end := uint64(head.Height)+1, uint64(newHead.Height)
	for start <= end {
		amount := end - start + 1
		if amount > requestSize {
			amount = requestSize
		}

		headers, err := s.getHeaders(ctx, start, amount)
		if err != nil && len(headers) == 0 {
			return err
		}

		err = s.store.Append(ctx, headers...)
		if err != nil {
			return err
		}

		start += uint64(len(headers))
	}

	return nil
}

// getHeaders gets headers from either remote peers or from local cached of headers rcvd by PubSub
func (s *Syncer) getHeaders(ctx context.Context, start, amount uint64) ([]*ExtendedHeader, error) {
	// short-circuit if nothing in pending cache to avoid unnecessary allocation below
	if _, ok := s.pending.BackWithin(start, start+amount); !ok {
		return s.exchange.RequestHeaders(ctx, start, amount)
	}

	end, out := start+amount, make([]*ExtendedHeader, 0, amount)
	for start < end {
		// if we have some range cached - use it
		if r, ok := s.pending.BackWithin(start, end); ok {
			// first, request everything between start and found range
			hs, err := s.exchange.RequestHeaders(ctx, start, r.Start-start)
			if err != nil {
				return nil, err
			}
			out = append(out, hs...)
			start += uint64(len(hs))

			// than, apply cached range
			cached := r.Before(end)
			out = append(out, cached...)
			start += uint64(len(cached))

			// repeat, as there might be multiple cache ranges
			continue
		}

		// fetch the leftovers
		hs, err := s.exchange.RequestHeaders(ctx, start, end-start)
		if err != nil {
			// still return what was successfully gotten
			return out, err
		}

		return append(out, hs...), nil
	}

	return out, nil
}
