package controllers

import (
	"log"
	"rental-api/API/mappers"
	"rental-api/API/services"
	contract "rental-api/Application.contract/images/add_image"
	commands "rental-api/Application/command_handlers/car_image/add_image"
	"rental-api/Domain/responses"

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
// @Param request body contract.AddUrlToCarOfferRequest true "Image details"
// @Success 200 {object} contract.AddUrlToCarOfferResponse200 "Image added successfully"
// @Failure 400 {object} contract.AddUrlToCarOfferResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.AddUrlToCarOfferResponse401 "Unauthorized"
// @Failure 403 {object} contract.AddUrlToCarOfferResponse403 "Forbidden - Not authorized"
// @Failure 404 {object} contract.AddUrlToCarOfferResponse404 "Car offer not found"
// @Failure 500 {object} contract.AddUrlToCarOfferResponse500 "Server error during addition"
// @Router /rental-api/api/car-offers/{carOfferId}/images [post]
func (h *AddImageController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	token := services.GetJwtTokenFromQuery(c)
	var req contract.AddUrlToCarOfferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		response := responses.NewResponse[contract.AddUrlToCarOfferResponse](400, "Invalid request format")
		responseSender.Send(response)
		return
	}

	req.CarOfferId = c.Param("carOfferId")
	req.JwtToken = token
	if validateResponse := services.ValidateRequest[contract.AddUrlToCarOfferResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}

	command := mappers.MapToAddImageCommand(&req)
	resp := services.SendToMediator[*commands.AddImageCommand, *contract.AddUrlToCarOfferResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *AddImageController) Route() string {
	return "/car-offers/:carOfferId/images"
}

func (h *AddImageController) Methods() []string {
	return []string{"POST"}
}
