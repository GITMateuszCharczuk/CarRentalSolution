package queue

import (
	"file-storage/Domain/event"
	"log"
)

type JetStreamPublisher struct{}

func NewJetStreamPublisher() event.EventPublisher {
	return &JetStreamPublisher{}
}

func (p *JetStreamPublisher) PublishEvent(subject string, data interface{}) error {
	log.Printf("Publishing event: %s", subject)
	return PublishEvent(subject, data)
}
