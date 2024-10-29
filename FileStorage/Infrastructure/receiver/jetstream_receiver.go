package receiver

import (
	"encoding/json"
	"file-storage/Domain/event"
	"file-storage/Domain/models"
	"file-storage/Infrastructure/processor"
	"log"

	"github.com/nats-io/nats.go"
)

type JetStreamReceiver struct {
	subscription   *nats.Subscription
	eventProcessor *processor.EventProcessor
}

func NewJetStreamReceiver(js nats.JetStreamContext, eventProcessor *processor.EventProcessor) (event.EventReceiver, error) {
	receiver := &JetStreamReceiver{eventProcessor: eventProcessor}

	sub, err := js.Subscribe("events.*", receiver.ProcessEvent)
	if err != nil {
		return nil, err
	}

	receiver.subscription = sub
	return receiver, nil
}

func (r *JetStreamReceiver) ProcessEvent(m *nats.Msg) {
	var evt models.Event
	if err := json.Unmarshal(m.Data, &evt); err != nil {
		log.Printf("Failed to unmarshal event: %v", err)
		return
	}

	switch evt.Type {
	case models.EventTypeUpload:
		log.Printf("Processing 'upload' event with data: %v", evt.Data)
		if err := r.eventProcessor.ProcessUploadEvent(evt.Data); err != nil {
			log.Printf("Error processing upload event: %v", err)
		}

	case models.EventTypeDelete:
		log.Printf("Processing 'delete' event with data: %v", evt.Data)
		if err := r.eventProcessor.ProcessDeleteEvent(evt.Data); err != nil {
			log.Printf("Error processing delete event: %v", err)
		}

	default:
		log.Printf("Unknown event type: %s", evt.Type)
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
