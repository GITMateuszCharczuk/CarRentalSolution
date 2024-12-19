package controllers

import (
	"rental-api/API/mappers"
	"rental-api/API/services"
	contract "rental-api/Application.contract/car_offers/delete_car_offer"
	commands "rental-api/Application/command_handlers/car_offer/delete_car_offer"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DeleteCarOfferController struct {
	validator *validator.Validate
}

func NewDeleteCarOfferController(validator *validator.Validate) *DeleteCarOfferController {
	return &DeleteCarOfferController{validator: validator}
}

// Handle godoc
// @Summary Delete a car offer
// @Description Deletes an existing car offer
// @Tags car-offers
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Param id path string true "Car Offer ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Success 200 {object} contract.DeleteCarOfferResponse200 "Car offer deleted successfully"
// @Failure 400 {object} contract.DeleteCarOfferResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.DeleteCarOfferResponse401 "Unauthorized"
// @Failure 403 {object} contract.DeleteCarOfferResponse403 "Forbidden - Not authorized"
// @Failure 404 {object} contract.DeleteCarOfferResponse404 "Car offer not found"
// @Failure 500 {object} contract.DeleteCarOfferResponse500 "Server error during deletion"
// @Router /rental-api/api/car-offers/{id} [delete]
func (h *DeleteCarOfferController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	req := contract.DeleteCarOfferRequest{
		CarOfferId: services.ExtractFromPath(c, "id"),
		JwtToken:   services.GetJwtTokenFromQuery(c),
	}
	if validateResponse := services.ValidateRequest[contract.DeleteCarOfferResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	command := mappers.MapToDeleteCarOfferCommand(&req)
	resp := services.SendToMediator[*commands.DeleteCarOfferCommand, *contract.DeleteCarOfferResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *DeleteCarOfferController) Route() string {
	return "/car-offers/:id"
}

func (h *DeleteCarOfferController) Methods() []string {
	return []string{"DELETE"}
}
