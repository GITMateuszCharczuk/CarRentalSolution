package controllers

import (
	"rental-api/API/mappers"
	"rental-api/API/services"
	contract "rental-api/Application.contract/images/delete_image"
	commands "rental-api/Application/command_handlers/car_image/delete_image"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DeleteImageController struct {
	validator *validator.Validate
}

func NewDeleteImageController(validator *validator.Validate) *DeleteImageController {
	return &DeleteImageController{validator: validator}
}

// Handle godoc
// @Summary Delete image from car offer
// @Description Deletes an image from a car offer
// @Tags car-images
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Param carOfferId path string true "Car Offer ID" example:"371eb93a-054c-44db-b429-0a0ebe87c3b9"
// @Param imageId path string true "Image ID (UUID)" example:"550e8400-e29b-41d4-a716-446655440000"
// @Success 200 {object} contract.DeleteImageFromCarOfferResponse200 "Image deleted successfully"
// @Failure 400 {object} contract.DeleteImageFromCarOfferResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.DeleteImageFromCarOfferResponse401 "Unauthorized"
// @Failure 403 {object} contract.DeleteImageFromCarOfferResponse403 "Forbidden - Not authorized"
// @Failure 404 {object} contract.DeleteImageFromCarOfferResponse404 "Image or car offer not found"
// @Failure 500 {object} contract.DeleteImageFromCarOfferResponse500 "Server error during deletion"
// @Router /rental-api/api/car-offers/images/{carOfferId}/{imageId} [delete]
func (h *DeleteImageController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	req := contract.DeleteImageFromCarOfferRequest{
		Id:         services.ExtractFromPath(c, "imageId"),
		CarOfferId: services.ExtractFromPath(c, "carOfferId"),
		JwtToken:   services.GetJwtTokenFromQuery(c),
	}
	if validateResponse := services.ValidateRequest[contract.DeleteImageFromCarOfferResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	command := mappers.MapToDeleteImageCommand(&req)
	resp := services.SendToMediator[*commands.DeleteImageCommand, *contract.DeleteImageFromCarOfferResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *DeleteImageController) Route() string {
	return "/car-offers/images/:carOfferId/:imageId"
}

func (h *DeleteImageController) Methods() []string {
	return []string{"DELETE"}
}
