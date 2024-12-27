package main

import (
	"api-gateway/internal/config"
	"api-gateway/internal/middleware"
	"api-gateway/internal/routes"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig("../../.env")

	r := gin.Default()

	serviceURLs := []string{
		cfg.EmailServiceURL + "/email-service/api/swagger/doc.json",
		cfg.FileServiceURL + "/file-storage/api/swagger/doc.json",
		cfg.BlogApiURL + "/blog-api/api/swagger/doc.json",
		cfg.RentalApiURL + "/rental-api/api/swagger/doc.json",
		cfg.IdentityApiURL + "/identity-api/api/swagger/doc.json",
	}

	r.Use(cors.Default())
	r.Use(middleware.RequestLogger())
	r.Use(middleware.RecoveryMiddleware())
	r.Use(middleware.RateLimiter(cfg.RequestSentLimit, cfg.RequestSentTimeWindow))
	r.Use(middleware.RequestSizeLimiter(cfg.RequestSizeLimit * 1024 * 1024)) // x * 1MB

	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	routes.RegisterEmailRoutes(r, cfg.EmailServiceURL, cfg.MainApiRoute)
	routes.RegisterFileRoutes(r, cfg.FileServiceURL, cfg.MainApiRoute)
	routes.RegisterBlogApiRoutes(r, cfg.BlogApiURL, cfg.MainApiRoute)
	routes.RegisterRentalApiRoutes(r, cfg.RentalApiURL, cfg.MainApiRoute)
	routes.RegisterIdentityApiRoutes(r, cfg.IdentityApiURL, cfg.MainApiRoute)
	routes.RegisterSwaggerRoutes(r, serviceURLs, cfg.MainApiRoute)

	r.Run(cfg.ServicePort)
}
