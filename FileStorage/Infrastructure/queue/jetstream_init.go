package queue

import (
	"log"

	"github.com/nats-io/nats.go"
)

func InitializeJetStream(natsURL string) (nats.JetStreamContext, error) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		return nil, err
	}
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}
	log.Println("Connected to JetStream")
	return js, nil
}
