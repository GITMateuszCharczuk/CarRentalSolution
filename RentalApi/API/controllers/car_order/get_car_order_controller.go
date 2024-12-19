package controllers

import (
	"rental-api/API/mappers"
	"rental-api/API/services"
	contract "rental-api/Application.contract/car_orders/GetCarOrder"
	queries "rental-api/Application/query_handlers/car_order/get_car_order"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GetCarOrderController struct {
	validator *validator.Validate
}

func NewGetCarOrderController(validator *validator.Validate) *GetCarOrderController {
	return &GetCarOrderController{validator: validator}
}

// Handle godoc
// @Summary Get a car order by ID
// @Description Retrieves a specific car order by its ID
// @Tags car-orders
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Param id path string true "Car Order ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Success 200 {object} contract.GetCarOrderResponse200 "Car order retrieved successfully"
// @Failure 401 {object} contract.GetCarOrderResponse401 "Unauthorized"
// @Failure 404 {object} contract.GetCarOrderResponse404 "Car order not found"
// @Failure 500 {object} contract.GetCarOrderResponse500 "Server error during retrieval"
// @Router /rental-api/api/car-orders/{id} [get]
func (h *GetCarOrderController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	req := contract.GetCarOrderRequest{
		CarOrderId: services.ExtractFromPath(c, "id"),
		JwtToken:   services.GetJwtTokenFromQuery(c),
	}
	if validateResponse := services.ValidateRequest[contract.GetCarOrderResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	query := mappers.MapToGetCarOrderQuery(&req)
	resp := services.SendToMediator[*queries.GetCarOrderQuery, *contract.GetCarOrderResponse](c.Request.Context(), &query)
	responseSender.Send(resp)
}

func (h *GetCarOrderController) Route() string {
	return "/car-orders/:id"
}

func (h *GetCarOrderController) Methods() []string {
	return []string{"GET"}
}
