// main/wire.go

// go:build wireinject
//go:build wireinject
// +build wireinject

package main

import (
	"email-service/API/controllers"
	"email-service/API/server"
	"email-service/Domain/event"
	fetcher "email-service/Domain/fetcher"
	"email-service/Infrastructure/config"
	data_fetcher "email-service/Infrastructure/data_fetcher"
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
