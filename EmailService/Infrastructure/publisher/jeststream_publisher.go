// publisher/jetstream_publisher.go

package publisher

import (
	"encoding/json"
	"file-storage/Domain/event"
	"file-storage/Domain/models"
	"log"

	"github.com/nats-io/nats.go"
)

type JetStreamPublisher struct {
	js nats.JetStreamContext
}

func NewJetStreamPublisher(js nats.JetStreamContext) event.EventPublisher {
	return &JetStreamPublisher{js: js}
}

func (p *JetStreamPublisher) PublishEvent(subject string, data interface{}, eventType models.EventType) error {
	event := models.Event{
		Type: eventType,
		Data: data,
	}

	message, err := json.Marshal(event)
	if err != nil {
		return err
	}
	_, err = p.js.Publish(subject, message)
	if err == nil {
		log.Printf("Published event %s to %s", eventType, subject)
	}
	return err
}
