package routes

import (
	"context"
	"file-storage/Application/handlers"
	"log"
	"net/http"
	"time"
)

type Router struct {
	fileHandlers *handlers.Handlers
	server       *http.Server
}

func NewRouter(fileHandlers *handlers.Handlers) *Router {
	return &Router{fileHandlers: fileHandlers}
}

func (r *Router) RegisterRoutes() {
	mux := http.NewServeMux()
	for _, handler := range r.fileHandlers.All {
		mux.HandleFunc(handler.Route(), handler.Handle)
	}
	r.server = &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}

func (r *Router) StartServer() error {
	log.Println("Starting server on port")
	if err := r.server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (r *Router) StopServer(ctx context.Context) error {
	if r.server == nil {
		return nil
	}
	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	log.Println("Stopping server gracefully...")
	err := r.server.Shutdown(shutdownCtx)
	if err != nil {
		log.Printf("Error shutting down server: %v", err)
		return err
	}
	log.Println("Server gracefully stopped.")
	return nil
}
