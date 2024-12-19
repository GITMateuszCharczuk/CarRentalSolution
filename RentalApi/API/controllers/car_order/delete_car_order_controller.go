package controllers

import (
	"rental-api/API/mappers"
	"rental-api/API/services"
	contract "rental-api/Application.contract/car_orders/DeleteCarOrder"
	commands "rental-api/Application/command_handlers/car_order/delete_car_order"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DeleteCarOrderController struct {
	validator *validator.Validate
}

func NewDeleteCarOrderController(validator *validator.Validate) *DeleteCarOrderController {
	return &DeleteCarOrderController{validator: validator}
}

// Handle godoc
// @Summary Delete a car order
// @Description Deletes an existing car order
// @Tags car-orders
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Param id path string true "Car Order ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Success 200 {object} contract.DeleteCarOrderResponse200 "Car order deleted successfully"
// @Failure 400 {object} contract.DeleteCarOrderResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.DeleteCarOrderResponse401 "Unauthorized"
// @Failure 403 {object} contract.DeleteCarOrderResponse403 "Forbidden - Not authorized"
// @Failure 404 {object} contract.DeleteCarOrderResponse404 "Car order not found"
// @Failure 500 {object} contract.DeleteCarOrderResponse500 "Server error during deletion"
// @Router /rental-api/api/car-orders/{id} [delete]
func (h *DeleteCarOrderController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	req := contract.DeleteCarOrderRequest{
		CarOrderId: services.ExtractFromPath(c, "id"),
		JwtToken:   services.GetJwtTokenFromQuery(c),
	}
	if validateResponse := services.ValidateRequest[contract.DeleteCarOrderResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	command := mappers.MapToDeleteCarOrderCommand(&req)
	resp := services.SendToMediator[*commands.DeleteCarOrderCommand, *contract.DeleteCarOrderResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *DeleteCarOrderController) Route() string {
	return "/car-orders/:id"
}

func (h *DeleteCarOrderController) Methods() []string {
	return []string{"DELETE"}
}
