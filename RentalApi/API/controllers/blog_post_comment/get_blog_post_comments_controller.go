package controllers

import (
	"rental-api/API/mappers"
	"rental-api/API/services"
	contract "rental-api/Application.contract/BlogPostComments/GetBlogPostComments"
	queries "rental-api/Application/query_handlers/blog_post_comment/get_blog_post_comments"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GetBlogPostCommentsController struct {
	validator *validator.Validate
}

func NewGetBlogPostCommentsController(validator *validator.Validate) *GetBlogPostCommentsController {
	return &GetBlogPostCommentsController{validator: validator}
}

// Handle godoc
// @Summary Get comments for a blog post
// @Description Retrieves all comments for a specific blog post with pagination
// @Tags comments
// @Accept json
// @Produce json
// @Param ids query []string false "Filter by blog post IDs" example:"123e4567-e89b-12d3-a456-426614174000"
// @Param user_ids query []string false "Filter by user IDs" example:"123e4567-e89b-12d3-a456-426614174000"
// @Param date_time_from query string false "Filter from date" example:"2023-01-01T00:00:00Z"
// @Param date_time_to query string false "Filter to date" example:"2024-01-01T00:00:00Z"
// @Param sort_fields query []string false "Sort fields (field:asc|desc)" example:"created_at:desc"
// @Param page_size query int false "Page size" minimum:1 example:10
// @Param current_page query int false "Current page" minimum:1 example:1
// @Success 200 {object} contract.GetBlogPostCommentsResponse200 "Comments retrieved successfully"
// @Failure 404 {object} contract.GetBlogPostCommentsResponse404 "Blog post not found"
// @Failure 500 {object} contract.GetBlogPostCommentsResponse500 "Server error during retrieval"
// @Router /rental-api/api/posts/{id}/comments [get]
func (h *GetBlogPostCommentsController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	req := contract.GetBlogPostCommentsRequest{
		BlogPostIds:  services.ExtractQueryArray(c, "ids"),
		DateTimeFrom: c.Query("date_time_from"),
		DateTimeTo:   c.Query("date_time_to"),
		UserIds:      services.ExtractQueryArray(c, "user_ids"),
		SortQuery:    services.ExtractSortQuery(c),
		Pagination:   services.ExtractPagination(c),
	}
	if validateResponse := services.ValidateRequest[contract.GetBlogPostCommentsResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	query := mappers.MapToGetBlogPostCommentsQuery(&req)

	resp := services.SendToMediator[*queries.GetBlogPostCommentsQuery, *contract.GetBlogPostCommentsResponse](c.Request.Context(), &query)
	responseSender.Send(resp)
}

func (h *GetBlogPostCommentsController) Route() string {
	return "/posts/:id/comments"
}

func (h *GetBlogPostCommentsController) Methods() []string {
	return []string{"GET"}
}
