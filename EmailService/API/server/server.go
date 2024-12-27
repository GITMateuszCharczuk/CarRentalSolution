package server

import (
	"context"
	"email-service/API/controllers"
	"email-service/Infrastructure/config"
	"log"
	"net/http"
	"time"

	_ "email-service/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	Config      *config.Config
	Controllers *controllers.Controllers
	server      *http.Server
}

func NewServer(Controllers *controllers.Controllers, Config *config.Config) *Server {
	return &Server{
		Controllers: Controllers,
		Config:      Config,
	}
}

func (r *Server) RegisterRoutes() {
	router := gin.Default()

	apiGroup := router.Group(r.Config.ServiceAddress)
	{
		for _, handler := range r.Controllers.All {
			route := handler.Route()
			for _, method := range handler.Methods() {
				switch method {
				case "GET":
					apiGroup.GET(route, func(c *gin.Context) { handler.Handle(c) })
				case "POST":
					apiGroup.POST(route, func(c *gin.Context) { handler.Handle(c) })
				case "DELETE":
					apiGroup.DELETE(route, func(c *gin.Context) { handler.Handle(c) })
				case "PUT":
					apiGroup.PUT(route, func(c *gin.Context) { handler.Handle(c) })
				case "PATCH":
					apiGroup.PATCH(route, func(c *gin.Context) { handler.Handle(c) })
				}
			}
		}

		apiGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.server = &http.Server{
		Addr:    r.Config.ServicePort,
		Handler: router,
	}
}

func (r *Server) StartServer() error {
	log.Println("Starting server on address:", r.server.Addr)
	if err := r.server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (r *Server) StopServer(ctx context.Context) error {
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
