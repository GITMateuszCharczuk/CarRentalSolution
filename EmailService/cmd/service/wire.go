// main/wire.go

// go:build wireinject
//go:build wireinject
// +build wireinject

package main

import (
	"email-service/API/controllers"
	"email-service/API/routes"
	"email-service/Application/commands"
	"email-service/Application/queries"
	"email-service/Domain/event"
	"email-service/Domain/repository_interfaces"
	"email-service/Infrastructure/config"
	"email-service/Infrastructure/processor"
	"email-service/Infrastructure/publisher"
	"email-service/Infrastructure/queue"
	"email-service/Infrastructure/receiver"

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
		publisher.WireSet,
		processor.WireSet,
		receiver.WireSet,
		wire.Struct(new(InfrastructureComponents), "*"),
	)
	return &InfrastructureComponents{}, nil
}

func InitializeApi(FileRepository repository_interfaces.FileRepository, EventPublisher event.EventPublisher) (*routes.Router, error) {
	wire.Build(
		commands.WireSet,
		queries.WireSet,
		controllers.WireSet,
		routes.NewRouter,
	)
	return &routes.Router{}, nil
}
