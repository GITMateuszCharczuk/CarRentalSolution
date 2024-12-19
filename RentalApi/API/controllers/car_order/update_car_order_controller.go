package controllers

import (
	"rental-api/API/mappers"
	"rental-api/API/services"
	contract "rental-api/Application.contract/car_orders/UpdateCarOrder"
	commands "rental-api/Application/command_handlers/car_order/update_car_order"
	"rental-api/Domain/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UpdateCarOrderController struct {
	validator *validator.Validate
}

func NewUpdateCarOrderController(validator *validator.Validate) *UpdateCarOrderController {
	return &UpdateCarOrderController{validator: validator}
}

// Handle godoc
// @Summary Update a car order
// @Description Updates an existing car order with the provided details
// @Tags car-orders
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Param id path string true "Car Order ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Param request body contract.UpdateCarOrderRequest true "Updated car order details"
// @Success 200 {object} contract.UpdateCarOrderResponse200 "Car order updated successfully"
// @Failure 400 {object} contract.UpdateCarOrderResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.UpdateCarOrderResponse401 "Unauthorized"
// @Failure 403 {object} contract.UpdateCarOrderResponse403 "Forbidden - Not authorized"
// @Failure 404 {object} contract.UpdateCarOrderResponse404 "Car order not found"
// @Failure 500 {object} contract.UpdateCarOrderResponse500 "Server error during update"
// @Router /rental-api/api/car-orders/{id} [put]
func (h *UpdateCarOrderController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	token := services.GetJwtTokenFromQuery(c)
	id := services.ExtractFromPath(c, "id")
	var req contract.UpdateCarOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response := responses.NewResponse[contract.UpdateCarOrderResponse](400, "Invalid request format")
		responseSender.Send(response)
		return
	}
	req.Id = id
	req.JwtToken = token
	if validateResponse := services.ValidateRequest[contract.UpdateCarOrderResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	command := mappers.MapToUpdateCarOrderCommand(&req)
	resp := services.SendToMediator[*commands.UpdateCarOrderCommand, *contract.UpdateCarOrderResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *UpdateCarOrderController) Route() string {
	return "/car-orders/:id"
}

func (h *UpdateCarOrderController) Methods() []string {
	return []string{"PUT"}
}
