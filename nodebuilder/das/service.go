package das

import (
	"context"

	"github.com/celestiaorg/celestia-node/das"
)

//go:generate mockgen -destination=mocks/api.go -package=mocks . Module
type Module interface {
	// SamplingStats returns the current statistics over the DA sampling process.
	SamplingStats(ctx context.Context) (das.SamplingStats, error)
	WaitCatchUp(ctx context.Context) error
}

// API is a wrapper around Module for the RPC.
// TODO(@distractedm1nd): These structs need to be autogenerated.
type API struct {
	SamplingStats func(ctx context.Context) (das.SamplingStats, error)
	WaitCatchUp   func(ctx context.Context) error
}
