package routes

import (
	"api-gateway/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterFileRoutes(router *gin.Engine, serviceURL string) {
	router.GET("/files/:id", services.ReverseProxy(serviceURL))
	router.POST("/files/upload", services.ReverseProxy(serviceURL))
	router.DELETE("/files/:id", services.ReverseProxy(serviceURL))
}
