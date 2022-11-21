package fraud

import (
	"context"

	"github.com/celestiaorg/celestia-node/fraud"
)

var _ Module = (*API)(nil)

// Module encompasses the behavior necessary to subscribe and broadcast fraud proofs within the
// network. Any method signature changed here needs to also be changed in the API struct.
//
//go:generate mockgen -destination=mocks/api.go -package=mocks . Module
type Module interface {
	// Subscribe allows to subscribe on a Proof pub sub topic by its type.
	Subscribe(context.Context, fraud.ProofType) (<-chan Proof, error)
	// Get fetches fraud proofs from the disk by its type.
	Get(context.Context, fraud.ProofType) ([]Proof, error)
}

// API is a wrapper around Module for the RPC.
// TODO(@distractedm1nd): These structs need to be autogenerated.
type API struct {
	Internal struct {
		Subscribe func(context.Context, fraud.ProofType) (<-chan Proof, error) `perm:"read"`
		Get       func(context.Context, fraud.ProofType) ([]Proof, error)      `perm:"read"`
	}
}

func (api *API) Subscribe(ctx context.Context, proofType fraud.ProofType) (<-chan Proof, error) {
	return api.Internal.Subscribe(ctx, proofType)
}

func (api *API) Get(ctx context.Context, proofType fraud.ProofType) ([]Proof, error) {
	return api.Internal.Get(ctx, proofType)
}
