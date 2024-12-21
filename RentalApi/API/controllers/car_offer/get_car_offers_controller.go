package controllers

import (
	"rental-api/API/mappers"
	"rental-api/API/services"
	contract "rental-api/Application.contract/car_offers/get_car_offers"
	queries "rental-api/Application/query_handlers/car_offer/get_car_offers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GetCarOffersController struct {
	validator *validator.Validate
}

func NewGetCarOffersController(validator *validator.Validate) *GetCarOffersController {
	return &GetCarOffersController{validator: validator}
}

// Handle godoc
// @Summary Get car offers
// @Description Retrieves a list of car offers with optional filtering, pagination and sorting
// @Tags car-offers
// @Accept json
// @Produce json
// @Param page_size query int false "Page size" example:"10"
// @Param current_page query int false "Current page" example:"1"
// @Param ids query []string false "Filter by car offer IDs" example:"123e4567-e89b-12d3-a456-426614174000"
// @Param date_time_from query string false "Filter from date" example:"2023-01-01T00:00:00Z"
// @Param date_time_to query string false "Filter to date" example:"2024-01-01T00:00:00Z"
// @Param tags query []string false "Filter by tags" example:"luxury,sports"
// @Param visible query string false "Filter by visibility" example:"true"
// @Param sort_fields query []string false "Sort fields (field:asc|desc)" example:"created_at:desc"
// @Success 200 {object} contract.GetCarOffersResponse200 "Car offers retrieved successfully"
// @Failure 400 {object} contract.GetCarOffersResponse400 "Invalid request parameters"
// @Failure 500 {object} contract.GetCarOffersResponse500 "Server error during retrieval"
// @Router /rental-api/api/car-offers [get]
func (h *GetCarOffersController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	req := contract.GetCarOffersRequest{
		Ids:          services.ExtractQueryArray(c, "ids"),
		DateTimeFrom: c.Query("date_time_from"),
		DateTimeTo:   c.Query("date_time_to"),
		Tags:         services.ExtractQueryArray(c, "tags"),
		Visible:      c.Query("visible"),
		SortQuery:    services.ExtractSortQuery(c),
		Pagination:   services.ExtractPagination(c),
	}

	if validateResponse := services.ValidateRequest[contract.GetCarOffersResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}

	query := mappers.MapToGetCarOffersQuery(&req)
	resp := services.SendToMediator[*queries.GetCarOffersQuery, *contract.GetCarOffersResponse](c.Request.Context(), &query)
	responseSender.Send(resp)
}

func (h *GetCarOffersController) Route() string {
	return "/car-offers"
}

func (h *GetCarOffersController) Methods() []string {
	return []string{"GET"}
}
