package controllers

import (
	"identity-api/API/services"
	contract "identity-api/Application.contract/BlogPosts/GetTags"
	queries "identity-api/Application/query_handlers/get_tags"

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
// @Tags blog
// @Accept json
// @Produce json
// @Success 200 {object} contract.GetTagsResponse200 "Tags retrieved successfully"
// @Failure 500 {object} contract.GetTagsResponse500 "Server error during retrieval"
// @Router /blog-api/api/tags [get]
func (h *GetTagsController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	query := queries.GetTagsQuery{}
	resp := services.SendToMediator[*queries.GetTagsQuery, *contract.GetTagsResponse](c.Request.Context(), &query)
	responseSender.Send(resp)
}

func (h *GetTagsController) Route() string {
	return "/tags"
}

func (h *GetTagsController) Methods() []string {
	return []string{"GET"}
}
