package controllers

import (
	"rental-api/API/mappers"
	"rental-api/API/services"
	contract "rental-api/Application.contract/images/add_image"
	commands "rental-api/Application/command_handlers/car_image/add_image"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AddImageController struct {
	validator *validator.Validate
}

func NewAddImageController(validator *validator.Validate) *AddImageController {
	return &AddImageController{validator: validator}
}

// Handle godoc
// @Summary Add image to car offer
// @Description Adds a new image URL to an existing car offer
// @Tags car-images
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Param offerId path string true "Car Offer ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Param imageId path string true "Image ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Success 200 {object} contract.AddUrlToCarOfferResponse200 "Image added successfully"
// @Failure 400 {object} contract.AddUrlToCarOfferResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.AddUrlToCarOfferResponse401 "Unauthorized"
// @Failure 403 {object} contract.AddUrlToCarOfferResponse403 "Forbidden - Not authorized"
// @Failure 404 {object} contract.AddUrlToCarOfferResponse404 "Car offer not found"
// @Failure 500 {object} contract.AddUrlToCarOfferResponse500 "Server error during addition"
// @Router /rental-api/api/car-offers/images/{offerId}/{imageId} [post]
func (h *AddImageController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	req := contract.AddUrlToCarOfferRequest{
		CarOfferId: services.ExtractFromPath(c, "offerId"),
		ImageId:    services.ExtractFromPath(c, "imageId"),
		JwtToken:   services.GetJwtTokenFromQuery(c),
	}
	if validateResponse := services.ValidateRequest[contract.AddUrlToCarOfferResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}

	command := mappers.MapToAddImageCommand(&req)
	resp := services.SendToMediator[*commands.AddImageCommand, *contract.AddUrlToCarOfferResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *AddImageController) Route() string {
	return "/car-offers/images/:offerId/:imageId"
}

func (h *AddImageController) Methods() []string {
	return []string{"POST"}
}
