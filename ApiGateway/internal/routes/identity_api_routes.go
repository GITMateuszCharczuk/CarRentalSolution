package routes

import (
	"api-gateway/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterIdentityApiRoutes(router *gin.Engine, serviceURL string, mainApiRoute string) {
	// User management routes
	router.GET(mainApiRoute+"/users", services.ReverseProxy(serviceURL, "/identity-api/api", mainApiRoute))
	router.GET(mainApiRoute+"/user/info", services.ReverseProxy(serviceURL, "/identity-api/api", mainApiRoute))
	router.PUT(mainApiRoute+"/user", services.ReverseProxy(serviceURL, "/identity-api/api", mainApiRoute))
	router.DELETE(mainApiRoute+"/user/:id", services.ReverseProxy(serviceURL, "/identity-api/api", mainApiRoute))

	// Authentication routes
	router.POST(mainApiRoute+"/register", services.ReverseProxy(serviceURL, "/identity-api/api", mainApiRoute))
	router.POST(mainApiRoute+"/login", services.ReverseProxy(serviceURL, "/identity-api/api", mainApiRoute))

	// Token management routes
	router.GET(mainApiRoute+"/token/validate", services.ReverseProxy(serviceURL, "/identity-api/api", mainApiRoute))
	router.POST(mainApiRoute+"/token/refresh", services.ReverseProxy(serviceURL, "/identity-api/api", mainApiRoute))
}
