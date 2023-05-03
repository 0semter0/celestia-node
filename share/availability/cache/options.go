package cache

import (
	"fmt"
)

const (
	// DefaultWriteBatchSize defines the size of the batched header write.
	// Headers are written in batches not to thrash the underlying Datastore with writes.
	// TODO(@Wondertan, @renaynay): proper defaults should be set for specific node type. (#709)
	DefaultWriteBatchSize          = 2048
	DefaultCacheAvailabilityPrefix = "sampling_result"
)

// Parameters is the set of Parameters that must be configured for cache
// availability implementation
type Parameters struct {
	// WriteBatchSize defines the size of the batched header write.
	WriteBatchSize uint
}

// Option is a function that configures cache availability Parameters
type Option func(*Parameters)

// DefaultParameters returns the default Parameters' configuration values
// for the light availability implementation
func DefaultParameters() Parameters {
	return Parameters{
		WriteBatchSize: DefaultWriteBatchSize,
	}
}

// Validate validates the values in Parameters
func (ca *Parameters) Validate() error {
	if ca.WriteBatchSize <= 0 {
		return fmt.Errorf(
			"cache availability: invalid option: value for DefaultWriteBatchSize, %s, %s",
			"is negative or 0.",         // current value
			"value must greater than 0", // what the valueshould be
		)
	}

	return nil
}

// WithWriteBatchSize is a functional option that the Availability interface
// implementers use to set the WriteBatchSize configuration param
func WithWriteBatchSize(writeBatchSize uint) Option {
	return func(p *Parameters) {
		p.WriteBatchSize = writeBatchSize
	}
}
