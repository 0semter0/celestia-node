package das

import (
	"fmt"
	"time"
)

var (
	ErrInvalidOption = fmt.Errorf("das: invalid option")
)

func errInvalidOptionValue(optionName string, value string) error {
	return fmt.Errorf("%w: value %s cannot be %s", ErrInvalidOption, optionName, value)
}

// Type Option is the functional option that is applied to the daser instance for parameters configuration.
type Option func(*DASer)

// Type parameters is the set of parameters that must be configured for the daser
type Parameters struct {
	//  SamplingRange is the maximum amount of headers processed in one job.
	SamplingRange uint64

	// ConcurrencyLimit defines the maximum amount of sampling workers running in parallel.
	ConcurrencyLimit int

	// BackgroundStoreInterval is the period of time for background checkpointStore to perform a checkpoint backup.
	BackgroundStoreInterval time.Duration

	// PriorityQueueSize defines the size limit of the priority queue
	PriorityQueueSize int

	// SampleFrom is the height sampling will start from
	SampleFrom uint64
}

// DefaultParameters returns the default configuration values for the daser parameters
func DefaultParameters() Parameters {
	// TODO(@derrandz): parameters needs performance testing on real network to define optimal values
	return Parameters{
		SamplingRange:           100,
		ConcurrencyLimit:        16,
		BackgroundStoreInterval: 10 * time.Minute,
		PriorityQueueSize:       16 * 4,
		SampleFrom:              1,
	}
}

// Validate validates the values in Parameters
//
// All parameters must be positive and non-zero, except:
//
// BackgroundStoreInterval = 0 disables background storer,
// PriorityQueueSize = 0 disables prioritization of recently produced blocks for sampling
func (p *Parameters) Validate() error {
	// SamplingRange = 0 will cause the jobs' queue to be empty
	// Therefore no sampling jobs will be reserved and more importantly the DASer will break
	if p.SamplingRange <= 0 {
		return errInvalidOptionValue(
			"SamplingRange",
			"negative or 0",
		)
	}

	// ConcurrencyLimit = 0 will cause the number of workers to be 0 and
	// Thus no threads will be assigned to the waiting jobs therefore breaking the DASer
	if p.ConcurrencyLimit <= 0 {
		return errInvalidOptionValue(
			"ConcurrencyLimit",
			"negative or 0",
		)
	}

	// SampleFrom = 0 would tell the DASer to start sampling from block height 0
	// which does not exist therefore breaking the DASer.
	if p.SampleFrom <= 0 {
		return errInvalidOptionValue(
			"SampleFrome",
			"negative or 0",
		)
	}

	return nil
}

// WithSamplingRange is a functional option to configure the daser's `SamplingRange` parameter
// ```
//
//	WithSamplingRange(10)(daser)
//
// ```
//
// or
// ```
//
//	option := WithSamplingRange(10)
//	// shenanigans to create daser
//	option(daser)
//
// ```
func WithSamplingRange(samplingRange uint64) Option {
	return func(d *DASer) {
		d.params.SamplingRange = samplingRange
	}
}

// WithConcurrencyLimit is a functional option to configure the daser's `ConcurrencyLimit` parameter
// Refer to WithSamplingRange documentation to see an example of how to use this
func WithConcurrencyLimit(concurrencyLimit int) Option {
	return func(d *DASer) {
		d.params.ConcurrencyLimit = concurrencyLimit
	}
}

// WithBackgroundStoreInterval is a functional option to configure the daser's `backgroundStoreInterval` parameter
// Refer to WithSamplingRange documentation to see an example of how to use this
func WithBackgroundStoreInterval(backgroundStoreInterval time.Duration) Option {
	return func(d *DASer) {
		d.params.BackgroundStoreInterval = backgroundStoreInterval
	}
}

// WithPriorityQueueSize is a functional option to configure the daser's `priorityQueuSize` parameter
// Refer to WithSamplingRange documentation to see an example of how to use this
func WithPriorityQueueSize(priorityQueueSize int) Option {
	return func(d *DASer) {
		d.params.PriorityQueueSize = priorityQueueSize
	}
}

// WithSampleFrom is a functional option to configure the daser's `SampleFrom` parameter
// Refer to WithSamplingRange documentation to see an example of how to use this
func WithSampleFrom(sampleFrom uint64) Option {
	return func(d *DASer) {
		d.params.SampleFrom = sampleFrom
	}
}
