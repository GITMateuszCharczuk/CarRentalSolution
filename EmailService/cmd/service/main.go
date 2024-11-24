package main

import (
	"context"
	"email-service/API/routes"
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
	router, err := InitializeApi(components.DataFetcher, components.EventPublisher, components.Config)
	if err != nil {
		log.Fatalf("Failed to initialize API: %v", err)
	}

	go func() {
		if err := components.EventReceiver.StartReceiving(); err != nil {
			log.Fatalf("Failed to start JetStream receiver: %v", err)
		}
	}()

	router.RegisterRoutes()

	go func() {
		if err := router.StartServer(); err != nil {
			log.Fatalf("Failed to start API server: %v", err)
		}
	}()

	waitForShutdown(router, components.EventReceiver)
}

func waitForShutdown(router *routes.Router, receiver event.EventReceiver) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Println("Shutting down application...")

	if err := receiver.StopReceiving(); err != nil {
		log.Printf("Error stopping JetStream receiver: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := router.StopServer(ctx); err != nil {
		log.Printf("Error stopping API server: %v", err)
	}

	log.Println("Api gracefully stopped.")
}
