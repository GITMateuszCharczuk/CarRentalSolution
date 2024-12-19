package controllers

import (
	"log"
	"rental-api/API/mappers"
	"rental-api/API/services"
	contract "rental-api/Application.contract/car_orders/CreateCarOrder"
	commands "rental-api/Application/command_handlers/car_order/create_car_order"
	"rental-api/Domain/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateCarOrderController struct {
	validator *validator.Validate
}

func NewCreateCarOrderController(validator *validator.Validate) *CreateCarOrderController {
	return &CreateCarOrderController{validator: validator}
}

// Handle godoc
// @Summary Create a new car order
// @Description Creates a new car order with the provided details
// @Tags car-orders
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Param request body contract.CreateCarOrderRequest true "Car order details"
// @Success 200 {object} contract.CreateCarOrderResponse200 "Car order created successfully"
// @Failure 400 {object} contract.CreateCarOrderResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.CreateCarOrderResponse401 "Unauthorized"
// @Failure 500 {object} contract.CreateCarOrderResponse500 "Server error during creation"
// @Router /rental-api/api/car-orders [post]
func (h *CreateCarOrderController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	token := services.GetJwtTokenFromQuery(c)
	var req contract.CreateCarOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		response := responses.NewResponse[contract.CreateCarOrderResponse](400, "Invalid request format")
		responseSender.Send(response)
		return
	}

	req.JwtToken = token
	if validateResponse := services.ValidateRequest[contract.CreateCarOrderResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}

	command := mappers.MapToCreateCarOrderCommand(&req)
	resp := services.SendToMediator[*commands.CreateCarOrderCommand, *contract.CreateCarOrderResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *CreateCarOrderController) Route() string {
	return "/car-orders"
}

func (h *CreateCarOrderController) Methods() []string {
	return []string{"POST"}
}
