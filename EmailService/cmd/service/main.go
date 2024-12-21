package main

import (
	"context"
	"email-service/API/server"
	commandHandlers "email-service/Application/commmand_handlers"
	queryHandlers "email-service/Application/query_handlers"
	"email-service/Domain/event"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	components, err := InitializeInfrastructureComponents()
	if err != nil {
		log.Fatalf("Failed to initialize Infrastructure: %v", err)
	}

	commandHandlers.RegisterCommandHandlers(components.EventPublisher, components.Config, components.MicroserviceConnector)
	queryHandlers.RegisterQueryHandlers(components.DataFetcher, components.MicroserviceConnector)

	server, err := InitializeApi(components.DataFetcher, components.EventPublisher, components.Config, components.MicroserviceConnector)
	if err != nil {
		log.Fatalf("Failed to initialize API: %v", err)
	}

	go func() {
		if err := components.EventReceiver.StartReceiving(); err != nil {
			log.Fatalf("Failed to start JetStream receiver: %v", err)
		}
	}()

	server.RegisterRoutes()

	go func() {
		if err := server.StartServer(); err != nil {
			log.Fatalf("Failed to start API server: %v", err)
		}
	}()

	waitForShutdown(server, components.EventReceiver)
}

func waitForShutdown(server *server.Server, receiver event.EventReceiver) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Println("Shutting down application...")

	if err := receiver.StopReceiving(); err != nil {
		log.Printf("Error stopping JetStream receiver: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.StopServer(ctx); err != nil {
		log.Printf("Error stopping API server: %v", err)
	}

	log.Println("Api gracefully stopped.")
}
