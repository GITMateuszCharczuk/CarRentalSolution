package receiver

import (
	"file-storage/Domain/event"
	"log"

	"github.com/nats-io/nats.go"
)

type JetStreamReceiver struct {
	subscription *nats.Subscription
}

func NewJetStreamReceiver(js nats.JetStreamContext) (event.EventReceiver, error) {
	receiver := &JetStreamReceiver{}

	sub, err := js.Subscribe("events.*", receiver.ProcessEvent)
	if err != nil {
		return nil, err
	}

	receiver.subscription = sub
	return receiver, nil
}

func (r *JetStreamReceiver) ProcessEvent(m *nats.Msg) {
	var event event.Event
	if err := event.UnmarshalJSON(m.Data); err != nil {
		log.Printf("Failed to unmarshal event: %v", err)
		return
	}

	switch event.Type {
	case "save":
		log.Printf("Processing 'save' event with data: %v", event.Data)
	case "delete":
		log.Printf("Processing 'delete' event with data: %v", event.Data)
	default:
		log.Printf("Unknown event type: %s", event.Type)
	}
}

func (r *JetStreamReceiver) StartReceiving() error {
	log.Println("JetStream receiver is now listening for events...")
	return nil
}

func (r *JetStreamReceiver) StopReceiving() error {
	if r.subscription != nil {
		return r.subscription.Unsubscribe()
	}
	return nil
}
