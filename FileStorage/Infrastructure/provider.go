package infrastructure

import (
	"file-storage/Domain/event"
	"file-storage/Domain/repository_interfaces"
	"file-storage/Infrastructure/config"
	"file-storage/Infrastructure/db"
	"file-storage/Infrastructure/publisher"
	"file-storage/Infrastructure/queue"
	"file-storage/Infrastructure/receiver"
	"file-storage/Infrastructure/repository"

	"github.com/google/wire"
)

type InfrastructureComponents struct {
	FileRepository *repository_interfaces.FileRepository
	EventPublisher event.EventPublisher
	EventReceiver  event.EventReceiver
}

func InitializeInfrastructureComponents() (*InfrastructureComponents, error) {
	wire.Build(
		config.WireSet,
		db.WireSet,
		repository.WireSet,
		queue.WireSet,
		publisher.WireSet,
		receiver.WireSet,
		wire.Struct(new(InfrastructureComponents), "*"),
	)
	return &InfrastructureComponents{}, nil
}
