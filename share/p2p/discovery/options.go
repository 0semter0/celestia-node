package discovery

import (
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
)

// Parameters is the set of Parameters that must be configured for the Discovery module
type Parameters struct {
	// PeersLimit defines the soft limit of FNs to connect to via discovery.
	// Set 0 to disable.
	PeersLimit uint
	// AdvertiseInterval is a interval between advertising sessions.
	// NOTE: only full and bridge can advertise themselves.
	AdvertiseInterval time.Duration

	// AdvertiseRetryTimeout defines time interval between advertise attempts.
	AdvertiseRetryTimeout time.Duration

	// DiscoveryRetryTimeout defines time interval between discovery attempts
	// this is set independently for tests in discover_test.go
	DiscoveryRetryTimeout time.Duration
}

// options is the set of options that can be configured for the Discovery module
type options struct {
	// onUpdatedPeers will be called on peer set changes
	onUpdatedPeers OnUpdatedPeers
	// advertise indicates whether the node should also
	// advertise to the discovery instance's topic
	advertise bool
}

// Option is a function that configures Discovery Parameters
type Option func(*options)

// DefaultParameters returns the default Parameters' configuration values
// for the Discovery module
func DefaultParameters() *Parameters {
	return &Parameters{
		PeersLimit:            5,
		AdvertiseInterval:     time.Hour,
		AdvertiseRetryTimeout: time.Second,
		DiscoveryRetryTimeout: time.Second * 60,
	}
}

// TestParameters returns the default Parameters' configuration values
// for the Discovery module, with some changes for configuration
// during tests
func TestParameters() *Parameters {
	p := DefaultParameters()
	p.AdvertiseInterval = time.Second * 1
	p.DiscoveryRetryTimeout = time.Millisecond * 50
	return p
}

// Validate validates the values in Parameters
func (p *Parameters) Validate() error {
	if p.AdvertiseRetryTimeout <= 0 {
		return fmt.Errorf("discovery: advertise retry timeout cannot be zero or negative")
	}

	if p.DiscoveryRetryTimeout <= 0 {
		return fmt.Errorf("discovery: discovery retry timeout cannot be zero or negative")
	}

	if p.PeersLimit <= 0 {
		return fmt.Errorf("discovery: peers limit cannot be zero or negative")
	}

	if p.AdvertiseInterval <= 0 {
		return fmt.Errorf("discovery: advertise interval cannot be zero or negative")
	}
	return nil
}

// WithOnPeersUpdate chains OnPeersUpdate callbacks on every update of discovered peers list.
func WithOnPeersUpdate(f OnUpdatedPeers) Option {
	return func(p *options) {
		p.onUpdatedPeers = p.onUpdatedPeers.add(f)
	}
}

func WithAdvertise() Option {
	return func(p *options) {
		p.advertise = true
	}
}

func newOptions(opts ...Option) *options {
	defaults := &options{
		onUpdatedPeers: func(peer.ID, bool) {},
		advertise:      false,
	}

	for _, opt := range opts {
		opt(defaults)
	}
	return defaults
}
