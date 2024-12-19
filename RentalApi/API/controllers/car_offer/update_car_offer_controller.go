package controllers

import (
	"rental-api/API/mappers"
	"rental-api/API/services"
	contract "rental-api/Application.contract/car_offers/update_car_offer"
	commands "rental-api/Application/command_handlers/car_offer/update_car_offer"
	"rental-api/Domain/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UpdateCarOfferController struct {
	validator *validator.Validate
}

func NewUpdateCarOfferController(validator *validator.Validate) *UpdateCarOfferController {
	return &UpdateCarOfferController{validator: validator}
}

// Handle godoc
// @Summary Update a car offer
// @Description Updates an existing car offer with the provided details
// @Tags car-offers
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Param id path string true "Car Offer ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Param request body contract.UpdateCarOfferRequest true "Updated car offer details"
// @Success 200 {object} contract.UpdateCarOfferResponse200 "Car offer updated successfully"
// @Failure 400 {object} contract.UpdateCarOfferResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.UpdateCarOfferResponse401 "Unauthorized"
// @Failure 403 {object} contract.UpdateCarOfferResponse403 "Forbidden - Not authorized"
// @Failure 404 {object} contract.UpdateCarOfferResponse404 "Car offer not found"
// @Failure 500 {object} contract.UpdateCarOfferResponse500 "Server error during update"
// @Router /rental-api/api/car-offers/{id} [put]
func (h *UpdateCarOfferController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	token := services.GetJwtTokenFromQuery(c)
	id := services.ExtractFromPath(c, "id")
	var req contract.UpdateCarOfferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response := responses.NewResponse[contract.UpdateCarOfferResponse](400, "Invalid request format")
		responseSender.Send(response)
		return
	}
	req.Id = id
	req.JwtToken = token
	if validateResponse := services.ValidateRequest[contract.UpdateCarOfferResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	command := mappers.MapToUpdateCarOfferCommand(&req)
	resp := services.SendToMediator[*commands.UpdateCarOfferCommand, *contract.UpdateCarOfferResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *UpdateCarOfferController) Route() string {
	return "/car-offers/:id"
}

func (h *UpdateCarOfferController) Methods() []string {
	return []string{"PUT"}
}
