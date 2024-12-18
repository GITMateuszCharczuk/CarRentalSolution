package main

import (
	"blog-api/API/server"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	command_handlers "blog-api/Application/command_handlers"
	query_handlers "blog-api/Application/query_handlers"
)

func main() {
	components, err := InitializeInfrastructureComponents()
	if err != nil {
		log.Fatalf("Failed to initialize Infrastructure: %v", err)
	}

	command_handlers.RegisterCommandHandlers(
		components.BlogCommandRepo,
		components.BlogQueryRepo,
		components.CommentCommandRepo,
		components.CommentQueryRepo,
		components.LikeCommandRepo,
		components.DataFetcher,
	)

	query_handlers.RegisterQueryHandlers(
		components.BlogQueryRepo,
		components.CommentQueryRepo,
		components.LikeQueryRepo,
		components.TagQueryRepo,
		components.DataFetcher,
	)

	server, err := InitializeApi(
		components.BlogQueryRepo,
		components.BlogCommandRepo,
		components.CommentQueryRepo,
		components.CommentCommandRepo,
		components.LikeQueryRepo,
		components.LikeCommandRepo,
		components.TagQueryRepo,
		components.DataFetcher,
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
