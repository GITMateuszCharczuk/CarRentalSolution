package main

import (
	"context"
	"identity-api/API/server"
	commandHandlers "identity-api/Application/commmand_handlers"
	queryHandlers "identity-api/Application/query_handlers"
	"identity-api/Domain/event"
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

	commandHandlers.RegisterCommandHandlers(components.EventPublisher, components.Config)
	queryHandlers.RegisterQueryHandlers(components.DataFetcher)

	server, err := InitializeApi(components.DataFetcher, components.EventPublisher, components.Config)
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
