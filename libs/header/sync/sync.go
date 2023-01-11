package sync

import (
	"context"
	"errors"
	"sync"
	"time"

	logging "github.com/ipfs/go-log/v2"

	"github.com/celestiaorg/celestia-node/libs/header"
)

var log = logging.Logger("header/sync")

// Syncer implements efficient synchronization for headers.
//
// Subjective header - the latest known header that is not expired (within trusting period)
// Network header - the latest header received from the network
//
// There are two main processes running in Syncer:
// 1. Main syncing loop(s.syncLoop)
//   - Performs syncing from the subjective(local chain view) header up to the latest known trusted header
//   - Syncs by requesting missing headers from Exchange or
//   - By accessing cache of pending and verified headers
//
// 2. Receives new headers from PubSub subnetwork (s.processIncoming)
//   - Usually, a new header is adjacent to the trusted head and if so, it is simply appended to the local store,
//     incrementing the subjective height and making it the new latest known trusted header.
//   - Or, if it receives a header further in the future,
//     verifies against the latest known trusted header
//     adds the header to pending cache(making it the latest known trusted header)
//     and triggers syncing loop to catch up to that point.
type Syncer[H header.Header] struct {
	sub      header.Subscriber[H]
	exchange header.Exchange[H]
	store    header.Store[H]

	// stateLk protects state which represents the current or latest sync
	stateLk sync.RWMutex
	state   State
	// signals to start syncing
	triggerSync chan struct{}
	// pending keeps ranges of valid new network headers awaiting to be appended to store
	pending ranges[H]
	// netReqLk ensures only one network head is requested at any moment
	netReqLk sync.RWMutex

	// controls lifecycle for syncLoop
	ctx    context.Context
	cancel context.CancelFunc

	Params *Parameters
}

// NewSyncer creates a new instance of Syncer.
func NewSyncer[H header.Header](
	exchange header.Exchange[H],
	store header.Store[H],
	sub header.Subscriber[H],
	opts ...Options,
) (*Syncer[H], error) {
	params := DefaultParameters()
	for _, opt := range opts {
		opt(&params)
	}
	if err := params.Validate(); err != nil {
		return nil, err
	}

	return &Syncer[H]{
		sub:         sub,
		exchange:    exchange,
		store:       store,
		triggerSync: make(chan struct{}, 1), // should be buffered
		Params:      &params,
	}, nil
}

// Start starts the syncing routine.
func (s *Syncer[H]) Start(ctx context.Context) error {
	s.ctx, s.cancel = context.WithCancel(context.Background())
	// register validator for header subscriptions
	// syncer does not subscribe itself and syncs headers together with validation
	err := s.sub.AddValidator(s.incomingNetHead)
	if err != nil {
		return err
	}
	// get the latest head and set it as syncing target
	_, err = s.networkHead(ctx)
	if err != nil {
		return err
	}
	// start syncLoop only if Start is errorless
	go s.syncLoop()
	return nil
}

// Stop stops Syncer.
func (s *Syncer[H]) Stop(context.Context) error {
	s.cancel()
	return nil
}

// WaitSync blocks until ongoing sync is done.
func (s *Syncer[H]) WaitSync(ctx context.Context) error {
	state := s.State()
	if state.Finished() {
		return nil
	}

	// this store method blocks until header is available
	_, err := s.store.GetByHeight(ctx, state.ToHeight)
	return err
}

// State collects all the information about a sync.
type State struct {
	ID                   uint64 // incrementing ID of a sync
	Height               uint64 // height at the moment when State is requested for a sync
	FromHeight, ToHeight uint64 // the starting and the ending point of a sync
	FromHash, ToHash     header.Hash
	Start, End           time.Time
	Error                error // the error that might happen within a sync
}

// Finished returns true if sync is done, false otherwise.
func (s State) Finished() bool {
	return s.ToHeight <= s.Height
}

// Duration returns the duration of the sync.
func (s State) Duration() time.Duration {
	return s.End.Sub(s.Start)
}

// State reports state of the current (if in progress), or last sync (if finished).
// Note that throughout the whole Syncer lifetime there might an initial sync and multiple
// catch-ups. All of them are treated as different syncs with different state IDs and other
// information.
func (s *Syncer[H]) State() State {
	s.stateLk.RLock()
	state := s.state
	s.stateLk.RUnlock()
	state.Height = s.store.Height()
	return state
}

// wantSync will trigger the syncing loop (non-blocking).
func (s *Syncer[H]) wantSync() {
	select {
	case s.triggerSync <- struct{}{}:
	default:
	}
}

// syncLoop controls syncing process.
func (s *Syncer[H]) syncLoop() {
	for {
		select {
		case <-s.triggerSync:
			s.sync(s.ctx)
		case <-s.ctx.Done():
			return
		}
	}
}

// sync ensures we are synced from the Store's head up to the new subjective head
func (s *Syncer[H]) sync(ctx context.Context) {
	newHead := s.pending.Head()
	if newHead.IsZero() {
		return
	}

	head, err := s.store.Head(ctx)
	if err != nil {
		log.Errorw("getting head during sync", "err", err)
		return
	}

	if head.Height() >= newHead.Height() {
		log.Warnw("sync attempt to an already synced header",
			"synced_height", head.Height(),
			"attempted_height", newHead.Height(),
		)
		log.Warn("PLEASE REPORT THIS AS A BUG")
		return // should never happen, but just in case
	}

	log.Infow("syncing headers",
		"from", head.Height(),
		"to", newHead.Height())
	err = s.doSync(ctx, head, newHead)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			// don't log this error as it is normal case of Syncer being stopped
			return
		}

		log.Errorw("syncing headers",
			"from", head.Height(),
			"to", newHead.Height(),
			"err", err)
		return
	}

	log.Infow("finished syncing",
		"from", head.Height(),
		"to", newHead.Height(),
		"elapsed time", s.state.End.Sub(s.state.Start))
}

// doSync performs actual syncing updating the internal State
func (s *Syncer[H]) doSync(ctx context.Context, fromHead, toHead H) (err error) {
	from, to := uint64(fromHead.Height())+1, uint64(toHead.Height())

	s.stateLk.Lock()
	s.state.ID++
	s.state.FromHeight = from
	s.state.ToHeight = to
	s.state.FromHash = fromHead.Hash()
	s.state.ToHash = toHead.Hash()
	s.state.Start = time.Now()
	s.stateLk.Unlock()

	err = s.processHeaders(ctx, fromHead, to)

	s.stateLk.Lock()
	s.state.End = time.Now()
	s.state.Error = err
	s.stateLk.Unlock()
	return err
}

// processHeaders gets and stores headers starting at the given 'from' height up to 'to' height -
// [from:to]
func (s *Syncer[H]) processHeaders(ctx context.Context, fromHeader H, to uint64) error {
	headers, err := s.requestHeaders(ctx, fromHeader, to)
	if err != nil {
		return err
	}
	_, err = s.store.Append(ctx, headers...)
	return err
}

// requestHeaders checks headers in pending cache that apply to the requested range.
// If some headers are missing, it starts requesting them from the network.
// All headers that will be received are verified to be contiguous.
func (s *Syncer[H]) requestHeaders(
	ctx context.Context,
	fromHeader H,
	to uint64,
) ([]H, error) {
	amount := to - uint64(fromHeader.Height())
	cached, ok := s.checkCache(fromHeader, to)
	if !ok {
		// request full range if cache is empty
		return s.findHeaders(ctx, fromHeader, to)
	}

	out := make([]H, 0, amount)
	for _, headers := range cached {
		if fromHeader.Height()+1 == headers[0].Height() {
			// apply cache
			out = append(out, headers...)
			// set new header to count from
			fromHeader = out[len(out)-1]
			continue
		}
		// make an external request
		h, err := s.findHeaders(ctx, fromHeader, uint64(headers[0].Height())-1)
		if err != nil {
			return nil, err
		}
		// apply received headers + headers from the range
		out = append(out, append(h, headers...)...)
		fromHeader = out[len(out)-1]
	}

	// ensure that we have all requested headers
	if uint64(len(out)) == amount {
		return out, nil
	}
	// make one more external request in case if `to` is bigger than the
	// last cached header
	h, err := s.findHeaders(ctx, fromHeader, to)
	if err != nil {
		return nil, err
	}
	return append(out, h...), nil
}

// checkCache returns all headers sub-ranges of headers that could be found in between the requested range.
func (s *Syncer[H]) checkCache(fromHeader H, to uint64) ([][]H, bool) {
	r := s.pending.FindAllRangesWithin(uint64(fromHeader.Height()+1), to) // start looking for headers for the next height
	if len(r) == 0 {
		return nil, false
	}

	out := make([][]H, len(r))
	for i, headersRange := range r {
		cached, _ := headersRange.Before(to)
		out[i] = append(out[i], cached...)
	}
	return out, len(out) > 0
}

// findHeaders requests headers from the network -> (fromHeader.Height;to].
func (s *Syncer[H]) findHeaders(
	ctx context.Context,
	fromHeader H,
	to uint64,
) ([]H, error) {
	amount := to - uint64(fromHeader.Height())
	// request full range if the amount of headers is less than the MaxRequestSize.
	if amount <= s.Params.MaxRequestSize {
		return s.exchange.GetVerifiedRange(ctx, fromHeader, amount)
	}
	out := make([]H, 0, amount)
	// start requesting headers until amount will be 0
	for amount > 0 {
		size := s.Params.MaxRequestSize
		if amount < size {
			size = amount
		}

		headers, err := s.exchange.GetVerifiedRange(ctx, fromHeader, size)
		if err != nil {
			return nil, err
		}
		out = append(out, headers...)
		fromHeader = out[len(out)-1]
		amount -= size
	}

	return out, nil
}
