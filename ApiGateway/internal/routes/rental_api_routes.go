package routes

import (
	"api-gateway/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRentalApiRoutes(router *gin.Engine, serviceURL string, mainApiRoute string) {
	// Car offers routes
	router.POST(mainApiRoute+"/car-offers", services.ReverseProxy(serviceURL, "/rental-api/api", mainApiRoute))
	router.PUT(mainApiRoute+"/car-offers/:id", services.ReverseProxy(serviceURL, "/rental-api/api", mainApiRoute))
	router.DELETE(mainApiRoute+"/car-offers/:id", services.ReverseProxy(serviceURL, "/rental-api/api", mainApiRoute))
	router.GET(mainApiRoute+"/car-offers/:id", services.ReverseProxy(serviceURL, "/rental-api/api", mainApiRoute))
	router.GET(mainApiRoute+"/car-offers", services.ReverseProxy(serviceURL, "/rental-api/api", mainApiRoute))

	// Car orders routes
	router.POST(mainApiRoute+"/car-orders", services.ReverseProxy(serviceURL, "/rental-api/api", mainApiRoute))
	router.PUT(mainApiRoute+"/car-orders/:id", services.ReverseProxy(serviceURL, "/rental-api/api", mainApiRoute))
	router.DELETE(mainApiRoute+"/car-orders/:id", services.ReverseProxy(serviceURL, "/rental-api/api", mainApiRoute))
	router.GET(mainApiRoute+"/car-orders/:id", services.ReverseProxy(serviceURL, "/rental-api/api", mainApiRoute))
	router.GET(mainApiRoute+"/car-orders", services.ReverseProxy(serviceURL, "/rental-api/api", mainApiRoute))

	// Car offer images routes
	router.POST(mainApiRoute+"/car-offers/images/:offerId/:imageId", services.ReverseProxy(serviceURL, "/rental-api/api", mainApiRoute))
	router.DELETE(mainApiRoute+"/car-offers/images/:carOfferId/:imageId", services.ReverseProxy(serviceURL, "/rental-api/api", mainApiRoute))

	// Tags routes
	router.GET(mainApiRoute+"/car-offers/tags/:id", services.ReverseProxy(serviceURL, "/rental-api/api", mainApiRoute))

}
