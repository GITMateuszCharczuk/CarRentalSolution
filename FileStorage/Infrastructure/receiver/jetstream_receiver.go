package receiver

import (
	"encoding/json"
	"file-storage/Domain/event"
	"file-storage/Domain/models"
	"file-storage/Infrastructure/processor"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type JetStreamReceiverImpl struct {
	subscription   *nats.Subscription
	eventProcessor *processor.EventProcessorImpl
}

func NewJetStreamReceiver(js nats.JetStreamContext, eventProcessor *processor.EventProcessorImpl) (event.EventReceiver, error) {
	receiver := &JetStreamReceiverImpl{eventProcessor: eventProcessor}

	sub, err := js.Subscribe("file-events.*", receiver.ProcessEvent,
		nats.Durable("durable-consumer"),
		nats.AckWait(30*time.Second),
		nats.ManualAck(),
	)

	if err != nil {
		return nil, err
	}

	receiver.subscription = sub
	return receiver, nil
}

func (r *JetStreamReceiverImpl) ProcessEvent(m *nats.Msg) {
	var evt models.Event
	if err := json.Unmarshal(m.Data, &evt); err != nil {
		log.Printf("Failed to unmarshal event: %v", err)
		m.Nak()
		return
	}

	switch evt.Type {
	case models.EventTypeUpload:
		log.Printf("Processing 'upload' event")
		if err := r.eventProcessor.ProcessUploadEvent(evt.Data); err != nil {
			log.Printf("Error processing upload event: %v", err)
			m.Nak()
		}
		m.Ack()

	case models.EventTypeDelete:
		log.Printf("Processing 'delete' event")
		if err := r.eventProcessor.ProcessDeleteEvent(evt.Data); err != nil {
			log.Printf("Error processing delete event: %v", err)
			m.Nak()
		}
		m.Ack()

	default:
		log.Printf("Unknown event type: %s", evt.Type)
		m.Nak()
	}
}

func (r *JetStreamReceiverImpl) StartReceiving() error {
	log.Println("JetStream receiver is now listening for events...")
	return nil
}

func (r *JetStreamReceiverImpl) StopReceiving() error {
	if r.subscription != nil {
		return r.subscription.Unsubscribe()
	}
	return nil
}
