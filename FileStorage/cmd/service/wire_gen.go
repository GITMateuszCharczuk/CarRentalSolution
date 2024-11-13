// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"file-storage/API/controllers"
	"file-storage/API/routes"
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
	eventProcessor := processor.NewEventProcessor(fileRepository)
	eventReceiver, err := receiver.NewJetStreamReceiver(jetStreamContext, eventProcessor)
	if err != nil {
		return nil, err
	}
	infrastructureComponents := &InfrastructureComponents{
		FileRepository: fileRepository,
		EventPublisher: eventPublisher,
		EventReceiver:  eventReceiver,
	}
	return infrastructureComponents, nil
}

func InitializeApi(FileRepository repository_interfaces.FileRepository, EventPublisher event.EventPublisher) (*routes.Router, error) {
	saveFileCommandHandler := commands.ProvideSaveFileCommandHandler(FileRepository, EventPublisher)
	saveFileController := controllers.NewSaveFileController(saveFileCommandHandler)
	getFileQueryHandler := queries.ProvideGetFileQueryHandler(FileRepository)
	getFileController := controllers.NewGetFileController(getFileQueryHandler)
	deleteFileCommandHandler := commands.ProvideDeleteFileCommandHandler(FileRepository, EventPublisher)
	deleteFileController := controllers.NewDeleteFileController(deleteFileCommandHandler)
	v := controllers.ProvideControllers(saveFileController, getFileController, deleteFileController)
	controllersControllers := controllers.NewControllers(v)
	router := routes.NewRouter(controllersControllers)
	return router, nil
}

// wire.go:

type InfrastructureComponents struct {
	FileRepository repository_interfaces.FileRepository
	EventPublisher event.EventPublisher
	EventReceiver  event.EventReceiver
}
