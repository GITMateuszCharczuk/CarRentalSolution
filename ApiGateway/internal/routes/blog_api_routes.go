package routes

import (
	"api-gateway/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterBlogApiRoutes(router *gin.Engine, serviceURL string, mainApiRoute string) {
	// Posts routes
	router.POST(mainApiRoute+"/posts", services.ReverseProxy(serviceURL, "/blog-api/api", mainApiRoute))
	router.GET(mainApiRoute+"/posts/:id", services.ReverseProxy(serviceURL, "/blog-api/api", mainApiRoute))
	router.GET(mainApiRoute+"/posts", services.ReverseProxy(serviceURL, "/blog-api/api", mainApiRoute))
	router.PUT(mainApiRoute+"/posts/:id", services.ReverseProxy(serviceURL, "/blog-api/api", mainApiRoute))
	router.DELETE(mainApiRoute+"/posts/:id", services.ReverseProxy(serviceURL, "/blog-api/api", mainApiRoute))

	// Comments routes
	router.POST(mainApiRoute+"/posts/:id/comments", services.ReverseProxy(serviceURL, "/blog-api/api", mainApiRoute))
	router.DELETE(mainApiRoute+"/comments/:id", services.ReverseProxy(serviceURL, "/blog-api/api", mainApiRoute))
	router.GET(mainApiRoute+"/posts/:id/comments", services.ReverseProxy(serviceURL, "/blog-api/api", mainApiRoute))

	// Likes routes
	router.POST(mainApiRoute+"/posts/:id/likes", services.ReverseProxy(serviceURL, "/blog-api/api", mainApiRoute))
	router.DELETE(mainApiRoute+"/posts/:id/likes", services.ReverseProxy(serviceURL, "/blog-api/api", mainApiRoute))
	router.GET(mainApiRoute+"/posts/:id/likes", services.ReverseProxy(serviceURL, "/blog-api/api", mainApiRoute))

	// Tags routes
	router.GET(mainApiRoute+"/posts/tags/:id", services.ReverseProxy(serviceURL, "/blog-api/api", mainApiRoute))
}
