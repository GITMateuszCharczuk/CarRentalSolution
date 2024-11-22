package routes

import (
	"api-gateway/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterEmailRoutes(router *gin.Engine, serviceURL string) {
	router.GET("/emails/:id", services.ReverseProxy(serviceURL))
	router.GET("/emails", services.ReverseProxy(serviceURL))
	router.POST("/send-email", services.ReverseProxy(serviceURL))
}
