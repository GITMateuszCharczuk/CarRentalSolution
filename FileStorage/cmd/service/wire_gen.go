// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"file-storage/API/controllers"
	"file-storage/API/server"
	"file-storage/Domain/event"
	"file-storage/Domain/repository_interfaces"
	"file-storage/Infrastructure/config"
	"file-storage/Infrastructure/db"
	"file-storage/Infrastructure/processor"
	"file-storage/Infrastructure/publisher"
	"file-storage/Infrastructure/queue"
	"file-storage/Infrastructure/receiver"
	"file-storage/Infrastructure/repository"
)

// Injectors from wire.go:

func InitializeInfrastructureComponents() (*InfrastructureComponents, error) {
	configConfig := config.ProvideConfig()
	database := db.ProvideMongoDB(configConfig)
	collection := db.ProvideMongoCollection(database, configConfig)
	bucket := db.ProvideBucket(database)
	fileRepository := repository.ProvideFileRepository(collection, bucket)
	jetStreamContext, err := queue.ProvideJetStreamContext(configConfig)
	if err != nil {
		return nil, err
	}
	eventPublisher, err := publisher.ProvideEventPublisher(jetStreamContext)
	if err != nil {
		return nil, err
	}
	eventProcessorImpl := processor.NewEventProcessorImpl(fileRepository)
	eventReceiver, err := receiver.NewJetStreamReceiver(jetStreamContext, eventProcessorImpl)
	if err != nil {
		return nil, err
	}
	infrastructureComponents := &InfrastructureComponents{
		Config:         configConfig,
		FileRepository: fileRepository,
		EventPublisher: eventPublisher,
		EventReceiver:  eventReceiver,
	}
	return infrastructureComponents, nil
}

func InitializeApi(FileRepository repository_interfaces.FileRepository, EventPublisher event.EventPublisher, Config *config.Config) (*server.Server, error) {
	saveFileController := controllers.NewSaveFileController()
	getFileController := controllers.NewGetFileController()
	deleteFileController := controllers.NewDeleteFileController()
	v := controllers.ProvideControllers(saveFileController, getFileController, deleteFileController)
	controllersControllers := controllers.NewControllers(v)
	serverServer := server.ProvideRoutes(controllersControllers, Config)
	return serverServer, nil
}

// wire.go:

type InfrastructureComponents struct {
	Config         *config.Config
	FileRepository repository_interfaces.FileRepository
	EventPublisher event.EventPublisher
	EventReceiver  event.EventReceiver
}
