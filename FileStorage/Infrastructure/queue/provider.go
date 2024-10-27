package queue

import (
	"file-storage/Domain/event"
	"file-storage/Infrastructure/config"

	"github.com/google/wire"
)

func ProvideEventPublisher(cfg *config.Config) (event.EventPublisher, error) {
	js, err := InitializeJetStream(cfg.StreamName)
	if err != nil {
		return nil, err
	}
	return NewJetStreamPublisher(js), nil
}

var WireSet = wire.NewSet(ProvideEventPublisher)
