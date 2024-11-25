package receiver

import (
	"file-storage/Domain/event"
	"file-storage/Infrastructure/processor"

	"github.com/google/wire"
	"github.com/nats-io/nats.go"
)

func ProvideEventReceiver(jsContext nats.JetStreamContext, eventProcessor *processor.EventProcessorImpl) (event.EventReceiver, error) {
	return NewJetStreamReceiver(jsContext, eventProcessor)
}

var WireSet = wire.NewSet(NewJetStreamReceiver)
