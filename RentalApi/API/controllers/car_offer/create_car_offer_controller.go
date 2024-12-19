package controllers

import (
	"log"
	"rental-api/API/mappers"
	"rental-api/API/services"
	contract "rental-api/Application.contract/car_offers/create_car_offer"
	commands "rental-api/Application/command_handlers/car_offer/create_car_offer"
	"rental-api/Domain/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateCarOfferController struct {
	validator *validator.Validate
}

func NewCreateCarOfferController(validator *validator.Validate) *CreateCarOfferController {
	return &CreateCarOfferController{validator: validator}
}

// Handle godoc
// @Summary Create a new car offer
// @Description Creates a new car offer with the provided details
// @Tags car-offers
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Param request body contract.CreateCarOfferRequest true "Car offer details"
// @Success 200 {object} contract.CreateCarOfferResponse200 "Car offer created successfully"
// @Failure 400 {object} contract.CreateCarOfferResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.CreateCarOfferResponse401 "Unauthorized"
// @Failure 500 {object} contract.CreateCarOfferResponse500 "Server error during creation"
// @Router /rental-api/api/car-offers [post]
func (h *CreateCarOfferController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	token := services.GetJwtTokenFromQuery(c)
	var req contract.CreateCarOfferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		response := responses.NewResponse[contract.CreateCarOfferResponse](400, "Invalid request format")
		responseSender.Send(response)
		return
	}

	req.JwtToken = token
	if validateResponse := services.ValidateRequest[contract.CreateCarOfferResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}

	command := mappers.MapToCreateCarOfferCommand(&req)
	resp := services.SendToMediator[*commands.CreateCarOfferCommand, *contract.CreateCarOfferResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *CreateCarOfferController) Route() string {
	return "/car-offers"
}

func (h *CreateCarOfferController) Methods() []string {
	return []string{"POST"}
}
