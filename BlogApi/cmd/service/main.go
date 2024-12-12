package main

import (
	"context"
	"identity-api/API/server"
	command_handlers "identity-api/Application/command_handlers"
	query_handlers "identity-api/Application/query_handlers"
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

	command_handlers.RegisterCommandHandlers(
		components.PasswordHasher,
		components.UserQueryRepo,
		components.UserCommandRepo,
		components.TokenService,
	)
	query_handlers.RegisterQueryHandlers(
		components.UserQueryRepo,
		components.TokenService,
	)

	server, err := InitializeApi(
		components.UserQueryRepo,
		components.UserCommandRepo,
		components.TokenService,
		components.PasswordHasher,
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
