package fraud

import (
	"context"
	"encoding"
	"fmt"

	pubsub "github.com/libp2p/go-libp2p-pubsub"

	"github.com/celestiaorg/celestia-node/header"
)

type ErrFraudExists struct {
	Proof []Proof
}

func (e *ErrFraudExists) Error() string {
	return fmt.Sprintf("fraud: %s proof exists\n", e.Proof[0].Type())
}

type ProofType int

const (
	BadEncoding ProofType = iota
)

func (p ProofType) String() string {
	switch p {
	case BadEncoding:
		return "badencoding"
	default:
		panic(fmt.Sprintf("fraud: invalid proof type: %d", p))
	}
}

// Proof is a generic interface that will be used for all types of fraud proofs in the network.
type Proof interface {
	// Type returns the exact type of fraud proof.
	Type() ProofType
	// HeaderHash returns the block hash.
	HeaderHash() []byte
	// Height returns the block height corresponding to the Proof.
	Height() uint64
	// Validate check the validity of fraud proof.
	// Validate throws an error if some conditions don't pass and thus fraud proof is not valid.
	// NOTE: header.ExtendedHeader should pass basic validation otherwise it will panic if it's malformed.
	Validate(*header.ExtendedHeader) error

	encoding.BinaryMarshaler
}

// GetProof listens for the next Fraud Proof and unmarshal received binary into Proof.
func GetProof(ctx context.Context, s *pubsub.Subscription) (Proof, error) {
	data, err := s.Next(ctx)
	if err != nil {
		return nil, err
	}
	topic := s.Topic()
	unmarshaler := DefaultUnmarshalers[getProofTypeFromTopic(topic)]
	return unmarshaler(data.Data)
}

// OnProof subscribes on a single Fraud Proof.
// In case a Fraud Proof is received, then the given handle function will be invoked.
func OnProof(ctx context.Context, subscriber Subscriber, p ProofType, handle func(proof Proof)) {
	subscription, err := subscriber.Subscribe(p)
	if err != nil {
		log.Error(err)
		return
	}
	defer subscription.Cancel()

	// At this point we receive already verified fraud proof,
	// so there are no needs to call Validate.
	proof, err := GetProof(ctx, subscription)
	if err != nil {
		if err != context.Canceled {
			log.Errorw("reading next proof failed", "err", err)
		}
		return
	}

	handle(proof)
}
