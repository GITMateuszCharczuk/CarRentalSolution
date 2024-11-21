package infrastructure

import (
	"file-storage/Domain/event"
	smtp "file-storage/Infrastructure/MailHog"
	"file-storage/Infrastructure/config"
	"file-storage/Infrastructure/publisher"
	"file-storage/Infrastructure/queue"
	"file-storage/Infrastructure/receiver"

	"github.com/google/wire"
)

type InfrastructureComponents struct {
	EventPublisher event.EventPublisher
	EventReceiver  event.EventReceiver
}

func InitializeInfrastructureComponents() (*InfrastructureComponents, error) {
	wire.Build(
		config.WireSet,
		queue.WireSet,
		smtp.WireSet,
		receiver.WireSet,
		publisher.WireSet,
		wire.Struct(new(InfrastructureComponents), "*"),
	)
	return &InfrastructureComponents{}, nil
}
