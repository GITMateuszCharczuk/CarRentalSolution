// routes/router.go

package routes

import (
	"context"
	"file-storage/API/controllers"
	"file-storage/API/middleware"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Router struct {
	fileControllers *controllers.Controllers
	server          *http.Server
}

func NewRouter(fileControllers *controllers.Controllers) *Router {
	return &Router{fileControllers: fileControllers}
}

func (r *Router) RegisterRoutes() {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	for _, handler := range r.fileControllers.All {
		route := handler.Route()
		for _, method := range handler.Methods() {
			switch method {
			case "GET":
				router.GET(route, func(c *gin.Context) { handler.Handle(c) })
			case "POST":
				router.POST(route, func(c *gin.Context) { handler.Handle(c) })
			case "DELETE":
				router.DELETE(route, func(c *gin.Context) { handler.Handle(c) })
			case "PUT":
				router.PUT(route, func(c *gin.Context) { handler.Handle(c) })
			case "PATCH":
				router.PATCH(route, func(c *gin.Context) { handler.Handle(c) })
			}
		}
	}

	r.server = &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
}

func (r *Router) StartServer() error {
	log.Println("Starting server on port", r.server.Addr)
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
