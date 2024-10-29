package receiver

import (
	"file-storage/Domain/event"

	"github.com/google/wire"
	"github.com/nats-io/nats.go"
)

func ProvideEventReceiver(jsContext nats.JetStreamContext) (event.EventReceiver, error) {
	return NewJetStreamReceiver(jsContext)
}

var WireSet = wire.NewSet(NewJetStreamReceiver)
