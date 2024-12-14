package controllers

import (
	"identity-api/API/services"
	contract "identity-api/Application.contract/BlogPostComments/GetBlogPostComments"
	queries "identity-api/Application/query_handlers/get_blog_post_comments"

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
// @Tags blog,comments
// @Accept json
// @Produce json
// @Param blogPostId path string true "Blog Post ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Param page_size query int false "Page size" example:"10"
// @Param current_page query int false "Current page" example:"1"
// @Success 200 {object} contract.GetBlogPostCommentsResponse200 "Comments retrieved successfully"
// @Failure 404 {object} contract.GetBlogPostCommentsResponse404 "Blog post not found"
// @Failure 500 {object} contract.GetBlogPostCommentsResponse500 "Server error during retrieval"
// @Router /blog-api/api/posts/{blogPostId}/comments [get]
func (h *GetBlogPostCommentsController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)

	blogPostId := c.Param("blogPostId")
	if blogPostId == "" {
		response := contract.GetBlogPostCommentsResponse{
			BaseResponse: services.NewErrorResponse(400, "Blog post ID is required"),
		}
		responseSender.Send(&response)
		return
	}

	query := queries.GetBlogPostCommentsQuery{
		BlogPostId: blogPostId,
		Pagination: services.ExtractPagination(c),
	}

	resp := services.SendToMediator[*queries.GetBlogPostCommentsQuery, *contract.GetBlogPostCommentsResponse](c.Request.Context(), &query)
	responseSender.Send(resp)
}

func (h *GetBlogPostCommentsController) Route() string {
	return "/posts/:blogPostId/comments"
}

func (h *GetBlogPostCommentsController) Methods() []string {
	return []string{"GET"}
}
