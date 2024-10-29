package publisher

import (
	"encoding/json"
	"file-storage/Domain/event"
	"log"

	"github.com/nats-io/nats.go"
)

// JetStreamPublisher publishes events using JetStream
type JetStreamPublisher struct {
	js nats.JetStreamContext
}

// NewJetStreamPublisher creates a new JetStreamPublisher instance with the provided JetStream context
func NewJetStreamPublisher(js nats.JetStreamContext) event.EventPublisher {
	return &JetStreamPublisher{js: js}
}

func (p *JetStreamPublisher) PublishEvent(subject string, data interface{}) error {
	message, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = p.js.Publish(subject, message)
	if err == nil {
		log.Printf("Published event to %s", subject)
	}
	return err
}
