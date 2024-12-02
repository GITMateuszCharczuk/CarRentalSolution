// main/wire.go

// go:build wireinject
//go:build wireinject
// +build wireinject

package main

import (
	"identity-api/API/controllers"
	"identity-api/API/server"
	"identity-api/Domain/event"
	fetcher "identity-api/Domain/fetcher"
	"identity-api/Infrastructure/config"
	data_fetcher "identity-api/Infrastructure/data_fetcher"
	smtp "identity-api/Infrastructure/email_sender"
	processor "identity-api/Infrastructure/event_processor"
	publisher "identity-api/Infrastructure/event_publisher"
	receiver "identity-api/Infrastructure/event_receiver"
	"identity-api/Infrastructure/queue"

	"github.com/google/wire"
)

type InfrastructureComponents struct {
	EventPublisher event.EventPublisher
	EventReceiver  event.EventReceiver
	DataFetcher    fetcher.DataFetcher
	Config         *config.Config
}

func InitializeInfrastructureComponents() (*InfrastructureComponents, error) {
	wire.Build(
		config.WireSet,
		queue.WireSet,
		data_fetcher.WireSet,
		smtp.WireSet,
		publisher.WireSet,
		processor.WireSet,
		receiver.WireSet,
		wire.Struct(new(InfrastructureComponents), "*"),
	)
	return &InfrastructureComponents{}, nil
}

func InitializeApi(DataFetcher fetcher.DataFetcher, EventPublisher event.EventPublisher, cfg *config.Config) (*server.Server, error) {
	wire.Build(
		controllers.WireSet,
		server.WireSet,
	)
	return &server.Server{}, nil
}
