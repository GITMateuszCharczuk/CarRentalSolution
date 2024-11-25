// main/wire.go

// go:build wireinject
//go:build wireinject
// +build wireinject

package main

import (
	"file-storage/API/controllers"
	"file-storage/API/server"
	"file-storage/Application/commands"
	"file-storage/Application/queries"
	"file-storage/Domain/event"
	"file-storage/Domain/repository_interfaces"
	"file-storage/Infrastructure/config"
	"file-storage/Infrastructure/db"
	"file-storage/Infrastructure/processor"
	"file-storage/Infrastructure/publisher"
	"file-storage/Infrastructure/queue"
	"file-storage/Infrastructure/receiver"
	"file-storage/Infrastructure/repository"

	"github.com/google/wire"
)

type InfrastructureComponents struct {
	Config         *config.Config
	FileRepository repository_interfaces.FileRepository
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
		processor.WireSet,
		receiver.WireSet,
		wire.Struct(new(InfrastructureComponents), "*"),
	)
	return &InfrastructureComponents{}, nil
}

func InitializeApi(FileRepository repository_interfaces.FileRepository, EventPublisher event.EventPublisher, Config *config.Config) (*server.Server, error) {
	wire.Build(
		commands.WireSet,
		queries.WireSet,
		controllers.WireSet,
		server.WireSet,
	)
	return &server.Server{}, nil
}
