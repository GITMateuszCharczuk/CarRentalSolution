// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"email-service/API/controllers"
	"email-service/API/routes"
	"email-service/Application/commands"
	"email-service/Application/queries"
	"email-service/Domain/event"
	"email-service/Domain/fetcher"
	"email-service/Infrastructure/config"
	"email-service/Infrastructure/data_fetcher"
	"email-service/Infrastructure/email_sender"
	"email-service/Infrastructure/event_processor"
	"email-service/Infrastructure/event_publisher"
	"email-service/Infrastructure/event_receiver"
	"email-service/Infrastructure/queue"
)

// Injectors from wire.go:

func InitializeInfrastructureComponents() (*InfrastructureComponents, error) {
	configConfig := config.ProvideConfig()
	jetStreamContext, err := queue.ProvideJetStreamContext(configConfig)
	if err != nil {
		return nil, err
	}
	eventPublisher, err := publisher.ProvideEventPublisher(jetStreamContext)
	if err != nil {
		return nil, err
	}
	emailSender, err := smtp.ProvideEmailService(configConfig)
	if err != nil {
		return nil, err
	}
	eventProcessorImpl := processor.NewEventProcessor(emailSender)
	eventReceiver, err := receiver.NewJetStreamReceiver(jetStreamContext, eventProcessorImpl)
	if err != nil {
		return nil, err
	}
	dataFetcher := datafetcher.ProvideDataFetcherImpl(configConfig)
	infrastructureComponents := &InfrastructureComponents{
		EventPublisher: eventPublisher,
		EventReceiver:  eventReceiver,
		DataFetcher:    dataFetcher,
		Config:         configConfig,
	}
	return infrastructureComponents, nil
}

func InitializeApi(DataFetcher fetcher.DataFetcher, EventPublisher event.EventPublisher, cfg *config.Config) (*routes.Router, error) {
	getEmailQueryHandler := queries.ProvideGetEmailQueryHandler(DataFetcher)
	getEmailController := controllers.NewGetEmailController(getEmailQueryHandler)
	getEmailsQueryHandler := queries.ProvideGetEmailsQueryHandler(DataFetcher)
	getEmailsController := controllers.NewGetEmailsController(getEmailsQueryHandler)
	sendEmailCommandHandler := commands.ProvideSendEmailCommandHandler(EventPublisher, cfg)
	sendEmailController := controllers.NewSendEmailController(sendEmailCommandHandler)
	v := controllers.ProvideControllers(getEmailController, getEmailsController, sendEmailController)
	controllersControllers := controllers.NewControllers(v)
	router := routes.ProvideRouter(controllersControllers, cfg)
	return router, nil
}

// wire.go:

type InfrastructureComponents struct {
	EventPublisher event.EventPublisher
	EventReceiver  event.EventReceiver
	DataFetcher    fetcher.DataFetcher
	Config         *config.Config
}
