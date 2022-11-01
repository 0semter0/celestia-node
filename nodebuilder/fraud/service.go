package fraud

import (
	"context"
	"encoding/json"

	"github.com/celestiaorg/celestia-node/fraud"
)

var _ Module = (*Service)(nil)

// Service is an implementation of Module that uses fraud.Service as a backend. It is used to
// provide fraud proofs as a non-interface type to the API, and wrap fraud.Subscriber with a
// channel of Proofs.
type Service struct {
	fraud.Service
}

func (s *Service) Subscribe(ctx context.Context, proofType fraud.ProofType) (chan Proof, error) {
	subscription, err := s.Service.Subscribe(proofType)
	if err != nil {
		return nil, err
	}
	concreteProofs := make(chan Proof)
	go func() {
		proof, _ := subscription.Proof(ctx)
		concreteProofs <- Proof{Proof: proof}
	}()
	return concreteProofs, nil
}

func (s *Service) Get(ctx context.Context, proofType fraud.ProofType) ([]Proof, error) {
	proofs, err := s.Service.Get(ctx, proofType)
	if err != nil {
		return nil, err
	}
	concreteProofs := make([]Proof, len(proofs))
	for i, proof := range proofs {
		concreteProofs[i].Proof = proof
	}
	return concreteProofs, nil
}

// Proof embeds the fraud.Proof interface type to provide a concrete type for JSON serialization.
type Proof struct {
	fraud.Proof
}

func (f *Proof) UnmarshalJSON(data []byte) error {
	type fraudProof struct {
		ProofType fraud.ProofType `json:"proof_type"`
		Data      []byte          `json:"data"`
	}
	var fp fraudProof
	err := json.Unmarshal(data, &fp)
	if err != nil {
		return err
	}
	f.Proof, err = fraud.Unmarshal(fp.ProofType, fp.Data)
	if err != nil {
		return err
	}
	return nil
}

func (f *Proof) MarshalJSON() ([]byte, error) {
	marshaledProof, err := f.MarshalBinary()
	if err != nil {
		return nil, err
	}
	fraudProof := &struct {
		ProofType fraud.ProofType `json:"proof_type"`
		Data      []byte          `json:"data"`
	}{
		ProofType: f.Type(),
		Data:      marshaledProof,
	}
	return json.Marshal(fraudProof)
}
