package discovery

import (
	"errors"
	"sync"

	"github.com/libp2p/go-libp2p-core/peer"
)

// LimitedSet is a thread safe set of peers with given limit.
// Inspired by libp2p peer.Set but extended with Remove method.
type LimitedSet struct {
	lk sync.RWMutex
	ps map[peer.ID]struct{}

	limit int
}

// NewLimitedSet constructs a set with the maximum peers amount.
func NewLimitedSet(size int) *LimitedSet {
	ps := new(LimitedSet)
	ps.ps = make(map[peer.ID]struct{})
	ps.limit = size
	return ps
}

func (ps *LimitedSet) Contains(p peer.ID) bool {
	ps.lk.RLock()
	_, ok := ps.ps[p]
	ps.lk.RUnlock()
	return ok
}

func (ps *LimitedSet) Size() int {
	ps.lk.RLock()
	defer ps.lk.RUnlock()
	return len(ps.ps)
}

// TryAdd Attempts to add the given peer into the set.
// This operation will fail if the number of peers in the set is equal to size
func (ps *LimitedSet) TryAdd(p peer.ID) error {
	ps.lk.Lock()
	defer ps.lk.Unlock()
	if len(ps.ps) < ps.limit {
		ps.ps[p] = struct{}{}
		return nil
	}

	return errors.New("discovery: peers limit reached")
}

func (ps *LimitedSet) Remove(id peer.ID) {
	ps.lk.Lock()
	defer ps.lk.Unlock()
	if ps.limit > 0 {
		delete(ps.ps, id)
	}
}

func (ps *LimitedSet) Peers() []peer.ID {
	ps.lk.Lock()
	out := make([]peer.ID, 0, len(ps.ps))
	for p := range ps.ps {
		out = append(out, p)
	}
	ps.lk.Unlock()
	return out
}
