package main

import (
	"api-gateway/internal/config"
	"api-gateway/internal/middleware"
	"api-gateway/internal/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig("../../.env")

	r := gin.Default()

	r.Use(middleware.RequestLogger())
	r.Use(middleware.RecoveryMiddleware())
	r.Use(middleware.RateLimiter(cfg.RequestSentLimit, cfg.RequestSentTimeWindow))
	r.Use(middleware.RequestSizeLimiter(cfg.RequestSizeLimit * 1024 * 1024)) // x * 1MB

	routes.RegisterEmailRoutes(r, cfg.EmailServiceURL)
	routes.RegisterFileRoutes(r, cfg.FileServiceURL)

	r.Run(fmt.Sprintf(":%s", cfg.ServicePort))
}
