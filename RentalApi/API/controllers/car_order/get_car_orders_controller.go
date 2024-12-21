package controllers

import (
	"rental-api/API/mappers"
	"rental-api/API/services"
	contract "rental-api/Application.contract/car_orders/GetCarOrders"
	queries "rental-api/Application/query_handlers/car_order/get_car_orders"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GetCarOrdersController struct {
	validator *validator.Validate
}

func NewGetCarOrdersController(validator *validator.Validate) *GetCarOrdersController {
	return &GetCarOrdersController{validator: validator}
}

// Handle godoc
// @Summary Get car orders
// @Description Retrieves a list of car orders with optional filtering, pagination and sorting
// @Tags car-orders
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Param page_size query int false "Page size" example:"10"
// @Param current_page query int false "Current page" example:"1"
// @Param start_date query string false "Filter by start date" example:"2023-01-01T00:00:00Z"
// @Param end_date query string false "Filter by end date" example:"2024-01-01T00:00:00Z"
// @Param user_id query string false "Filter by user ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Param car_offer_id query string false "Filter by car offer ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Param statuses query []string false "Filter by statuses" example:"pending,active,overdue"
// @Param date_filter_type query string false "Type of date filter" example:"created"
// @Param sort_fields query []string false "Sort fields (field:asc|desc)" example:"created_at:desc"
// @Success 200 {object} contract.GetCarOrdersResponse200 "Car orders retrieved successfully"
// @Failure 401 {object} contract.GetCarOrdersResponse401 "Unauthorized"
// @Failure 500 {object} contract.GetCarOrdersResponse500 "Server error during retrieval"
// @Router /rental-api/api/car-orders [get]
func (h *GetCarOrdersController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	token := services.GetJwtTokenFromQuery(c)
	req := contract.GetCarOrdersRequest{ //todo add another variables
		StartDate:      c.Query("start_date"),
		EndDate:        c.Query("end_date"),
		UserId:         c.Query("user_id"),
		CarOfferId:     c.Query("car_offer_id"),
		Statuses:       services.ExtractQueryArray(c, "statuses"),
		DateFilterType: c.Query("date_filter_type"),
		SortQuery:      services.ExtractSortQuery(c),
		Pagination:     services.ExtractPagination(c),
		JwtToken:       token,
	}

	if validateResponse := services.ValidateRequest[contract.GetCarOrdersResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}

	query := mappers.MapToGetCarOrdersQuery(&req)
	resp := services.SendToMediator[*queries.GetCarOrdersQuery, *contract.GetCarOrdersResponse](c.Request.Context(), &query)
	responseSender.Send(resp)
}

func (h *GetCarOrdersController) Route() string {
	return "/car-orders"
}

func (h *GetCarOrdersController) Methods() []string {
	return []string{"GET"}
}
