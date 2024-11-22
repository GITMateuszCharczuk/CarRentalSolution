package receiver

import (
	"email-service/Domain/event"
	"email-service/Domain/models"
	processor "email-service/Infrastructure/event_processor"
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type JetStreamReceiver struct {
	subscription   *nats.Subscription
	eventProcessor *processor.EventProcessorImpl
}

func NewJetStreamReceiver(js nats.JetStreamContext, eventProcessor *processor.EventProcessorImpl) (event.EventReceiver, error) {
	receiver := &JetStreamReceiver{eventProcessor: eventProcessor}

	sub, err := js.Subscribe("email-events.*", receiver.ProcessEvent,
		nats.Durable("email-consumer"),
		nats.AckWait(30*time.Second),
		nats.ManualAck(),
	)

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
		m.Nak()
		return
	}

	switch evt.Type {
	case models.EventTypeSendEmail:
		log.Printf("Processing 'send_email' event with data: %v", evt.Data)
		if err := r.eventProcessor.ProcessSendEmailEvent(evt.Data); err != nil {
			log.Printf("Error processing send_email event: %v", err)
			m.Nak()
			return
		}
		m.Ack()

	default:
		log.Printf("Unknown event type: %s", evt.Type)
		m.Nak()
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
