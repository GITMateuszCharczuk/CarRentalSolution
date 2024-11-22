package publisher

import (
	"email-service/Domain/event"

	"github.com/google/wire"
	"github.com/nats-io/nats.go"
)

func ProvideEventPublisher(jsContext nats.JetStreamContext) (event.EventPublisher, error) {
	return NewJetStreamPublisher(jsContext), nil
}

var WireSet = wire.NewSet(ProvideEventPublisher)
