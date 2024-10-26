package queue

import (
	"log"

	"github.com/nats-io/nats.go"
)

var js nats.JetStreamContext

func InitializeJetStream() {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ = nc.JetStream()
}

func PublishEvent(event string) {
	js.Publish("files.events", []byte(event))
}

func SubscribeToQueue(queue string) {
	_, err := js.QueueSubscribe("files.events", queue, func(msg *nats.Msg) {
		log.Printf("Received message: %s", string(msg.Data))
	})
	if err != nil {
		log.Fatalf("Failed to subscribe to queue: %v", err)
	}
}
