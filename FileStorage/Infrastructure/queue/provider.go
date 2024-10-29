package queue

import (
	"file-storage/Infrastructure/config"

	"github.com/google/wire"
	"github.com/nats-io/nats.go"
)

func ProvideJetStreamContext(cfg *config.Config) (nats.JetStreamContext, error) {
	js, err := InitializeJetStream(cfg.StreamName)
	if err != nil {
		return nil, err
	}
	return js, nil
}

var WireSet = wire.NewSet(ProvideJetStreamContext)
