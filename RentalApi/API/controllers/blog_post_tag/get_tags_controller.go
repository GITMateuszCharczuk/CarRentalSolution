package controllers

import (
	"rental-api/API/mappers"
	"rental-api/API/services"
	contract "rental-api/Application.contract/Tags/GetTags"
	queries "rental-api/Application/query_handlers/blog_post_tag/get_tags"

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
// @Description Retrieves a list of all unique tags used in blog posts
// @Tags tags
// @Accept json
// @Produce json
// @Param id path string false "Blog Post ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Param sort_fields query []string false "Sort fields (field:asc|desc)" example:"created_at:desc"
// @Success 200 {object} contract.GetTagsResponse200 "Tags retrieved successfully"
// @Failure 500 {object} contract.GetTagsResponse500 "Server error during retrieval"
// @Router /rental-api/api/tags/{id} [get]
func (h *GetTagsController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	req := contract.GetTagsRequest{
		BlogPostId: services.ExtractFromPath(c, "id"),
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
	return "/tags/:id"
}

func (h *GetTagsController) Methods() []string {
	return []string{"GET"}
}
