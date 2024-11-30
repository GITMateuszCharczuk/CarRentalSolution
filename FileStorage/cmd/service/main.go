package main

import (
	"context"
	"file-storage/API/server"
	"file-storage/Application/commands"
	"file-storage/Application/queries"
	"file-storage/Domain/event"
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

	commands.RegisterCommandHandlers(components.EventPublisher)
	queries.RegisterQueryHandlers(components.FileRepository)

	server, err := InitializeApi(components.FileRepository, components.EventPublisher, components.Config)
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
