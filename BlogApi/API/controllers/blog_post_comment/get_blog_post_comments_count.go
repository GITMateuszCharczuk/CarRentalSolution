package controllers

import (
	"blog-api/API/mappers"
	"blog-api/API/services"
	contract "blog-api/Application.contract/BlogPostComments/GetBlogPostCommentsCount"
	queries "blog-api/Application/query_handlers/blog_post_comment/get_blog_post_comments_count"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GetBlogPostCommentsCountController struct {
	validator *validator.Validate
}

func NewGetBlogPostCommentsCountController(validator *validator.Validate) *GetBlogPostCommentsCountController {
	return &GetBlogPostCommentsCountController{validator: validator}
}

// Handle godoc
// @Summary Get comments for a blog post
// @Description Retrieves all comments for a specific blog post with pagination
// @Tags comments
// @Accept json
// @Produce json
// @Param blog_post_id query string false "Filter by blog post ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Success 200 {object} contract.GetBlogPostCommentsCountResponse200 "Comments retrieved successfully"
// @Failure 404 {object} contract.GetBlogPostCommentsCountResponse404 "Blog post not found"
// @Failure 500 {object} contract.GetBlogPostCommentsResponse500 "Server error during retrieval"
// @Router /blog-api/api/posts/comments/count [get]
func (h *GetBlogPostCommentsCountController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	req := contract.GetBlogPostCommentsCountRequest{
		BlogPostId: c.Query("blog_post_id"),
	}
	if validateResponse := services.ValidateRequest[contract.GetBlogPostCommentsCountResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	query := mappers.MapToGetBlogPostCommentsCountQuery(&req)

	resp := services.SendToMediator[*queries.GetBlogPostCommentsCountQuery, *contract.GetBlogPostCommentsCountResponse](c.Request.Context(), &query)
	responseSender.Send(resp)
}

func (h *GetBlogPostCommentsCountController) Route() string {
	return "/posts/comments/count"
}

func (h *GetBlogPostCommentsCountController) Methods() []string {
	return []string{"GET"}
}
