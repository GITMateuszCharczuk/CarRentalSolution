package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"rental-api/API/server"
	command_handlers "rental-api/Application/command_handlers"
	query_handlers "rental-api/Application/query_handlers"
	"syscall"
	"time"
)

func main() {
	components, err := InitializeInfrastructureComponents()
	if err != nil {
		log.Fatalf("Failed to initialize Infrastructure: %v", err)
	}
	command_handlers.RegisterCommandHandlers(
		components.CarOrderCommandRepo,
		components.CarOrderQueryRepo,
		components.CarOfferCommandRepo,
		components.CarOfferQueryRepo,
		components.CarImageCommandRepo,
		components.connector,
	)
	query_handlers.RegisterQueryHandlers(
		components.CarOrderQueryRepo,
		components.CarOfferQueryRepo,
		components.connector,
		components.CarTagQueryRepo,
		components.CarImageQueryRepo,
	)
	server, err := InitializeApi(
		components.CarOfferQueryRepo,
		components.CarOfferCommandRepo,
		components.CarOrderQueryRepo,
		components.CarOrderCommandRepo,
		components.CarImageQueryRepo,
		components.CarImageCommandRepo,
		components.CarTagQueryRepo,
		components.CarTagCommandRepo,
		components.connector,
		components.orderManagementSystem,
		components.Config,
	)
	if err != nil {
		log.Fatalf("Failed to initialize API: %v", err)
	}
	server.RegisterRoutes()
	go func() {
		if err := server.StartServer(); err != nil {
			log.Fatalf("Failed to start API server: %v", err)
		}
	}()
	go func() {
		components.orderManagementSystem.StartPeriodicCheck(context.Background())
	}()

	waitForShutdown(server)
}

func waitForShutdown(server *server.Server) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Shutting down application...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.StopServer(ctx); err != nil {
		log.Printf("Error stopping API server: %v", err)
	}
	log.Println("Api gracefully stopped.")
}
