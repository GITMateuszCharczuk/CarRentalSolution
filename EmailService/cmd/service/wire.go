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
	"email-service/Infrastructure/config"
	fetcher "email-service/Infrastructure/data_fetcher"
	smtp "email-service/Infrastructure/email_sender"
	processor "email-service/Infrastructure/event_processor"
	publisher "email-service/Infrastructure/event_publisher"
	receiver "email-service/Infrastructure/event_receiver"
	"email-service/Infrastructure/queue"

	"github.com/google/wire"
)

type InfrastructureComponents struct {
	EventPublisher event.EventPublisher
	EventReceiver  event.EventReceiver
	DataFetcher    fetcher.DataFetcher
}

func InitializeInfrastructureComponents() (*InfrastructureComponents, error) {
	wire.Build(
		config.WireSet,
		queue.WireSet,
		fetcher.WireSet,
		smtp.WireSet,
		publisher.WireSet,
		processor.WireSet,
		receiver.WireSet,
		wire.Struct(new(InfrastructureComponents), "*"),
	)
	return &InfrastructureComponents{}, nil
}

func InitializeApi(DataFetcher fetcher.DataFetcher, EventPublisher event.EventPublisher) (*routes.Router, error) {
	wire.Build(
		commands.WireSet,
		queries.WireSet,
		controllers.WireSet,
		routes.NewRouter,
	)
	return &routes.Router{}, nil
}
