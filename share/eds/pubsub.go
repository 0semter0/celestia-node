package eds

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"

	"github.com/celestiaorg/celestia-node/share"
)

// PubSubTopic hardcodes the name of the EDS floodsub topic.
const PubSubTopic = "eds-sub"

// Validator is an injectable func and governs EDS notification or DataHash validity.
// It receives the notification and sender peer and expects the validation result.
// Validator is allowed to be blocking for an indefinite time or until the context is canceled.
type Validator func(context.Context, peer.ID, share.DataHash) pubsub.ValidationResult

// PubSub manages receiving and propagating the EDS from/to the network
// over "eds-sub" subscription.
type PubSub struct {
	pubSub *pubsub.PubSub
	topic  *pubsub.Topic
}

// NewPubSub creates a libp2p.PubSub wrapper.
func NewPubSub(ctx context.Context, h host.Host) (*PubSub, error) {
	pubsub, err := pubsub.NewFloodSub(ctx, h)
	if err != nil {
		return nil, err
	}
	return &PubSub{
		pubSub: pubsub,
	}, nil
}

// Start creates an instances of FloodSub and joins specified topic.
func (s *PubSub) Start(context.Context) error {
	topic, err := s.pubSub.Join(PubSubTopic)
	if err != nil {
		return err
	}

	s.topic = topic
	return nil
}

// Stop completely stops the PubSub:
// * Unregisters all the added Validators
// * Closes the `ShrEx/Sub` topic
func (s *PubSub) Stop(context.Context) error {
	err := s.pubSub.UnregisterTopicValidator(PubSubTopic)
	if err != nil {
		return err
	}

	return s.topic.Close()
}

// AddValidator registers given Validator for EDS notifications (DataHash).
// Any amount of Validators can be registered.
func (s *PubSub) AddValidator(validate Validator) error {
	return s.pubSub.RegisterTopicValidator(PubSubTopic,
		func(ctx context.Context, p peer.ID, msg *pubsub.Message) pubsub.ValidationResult {
			return validate(ctx, p, msg.Data)
		})
}

// Subscribe provides a new Subscription for EDS notifications.
func (s *PubSub) Subscribe() (*Subscription, error) {
	if s.topic == nil {
		return nil, fmt.Errorf("share/eds: topic is not started")
	}
	return newSubscription(s.topic)
}

// Broadcast sends the EDS notification (DataHash) to every connected peer.
func (s *PubSub) Broadcast(ctx context.Context, data share.DataHash) error {
	return s.topic.Publish(ctx, data)
}
