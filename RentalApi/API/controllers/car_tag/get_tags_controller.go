package controllers

import (
	"rental-api/API/mappers"
	"rental-api/API/services"
	contract "rental-api/Application.contract/tags/get_tags"
	queries "rental-api/Application/query_handlers/car_tag/get_tags"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GetTagsController struct {
	validator *validator.Validate
}

func NewGetTagsController(validator *validator.Validate) *GetTagsController {
	return &GetTagsController{validator: validator}
}

// Handle godoc
// @Summary Get all unique tags
// @Description Retrieves a list of all unique tags used in car offer
// @Tags car-offers
// @Accept json
// @Produce json
// @Param carOfferId query string false "Car Offer ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Param sort_fields query []string false "Sort fields (field:asc|desc)" example:"created_at:desc"
// @Success 200 {object} contract.GetTagsResponse200 "Tags retrieved successfully"
// @Failure 500 {object} contract.GetTagsResponse500 "Server error during retrieval"
// @Router /rental-api/api/car-offers/tags [get]
func (h *GetTagsController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	req := contract.GetTagsRequest{
		CarOfferId: c.Query("carOfferId"),
		SortQuery:  services.ExtractSortQuery(c),
	}
	if validateResponse := services.ValidateRequest[contract.GetTagsResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	query := mappers.MapToGetTagsQuery(&req)
	resp := services.SendToMediator[*queries.GetTagsQuery, *contract.GetTagsResponse](c.Request.Context(), &query)
	responseSender.Send(resp)
}

func (h *GetTagsController) Route() string {
	return "/car-offers/tags"
}

func (h *GetTagsController) Methods() []string {
	return []string{"GET"}
}
