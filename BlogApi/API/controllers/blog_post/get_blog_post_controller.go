package controllers

import (
	"identity-api/API/mappers"
	"identity-api/API/services"
	contract "identity-api/Application.contract/BlogPosts/GetBlogPost"
	queries "identity-api/Application/query_handlers/blog_post/get_blog_post"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GetBlogPostController struct {
	validator *validator.Validate
}

func NewGetBlogPostController(validator *validator.Validate) *GetBlogPostController {
	return &GetBlogPostController{validator: validator}
}

// Handle godoc
// @Summary Get a blog post by ID
// @Description Retrieves a specific blog post by its ID
// @Tags blog
// @Accept json
// @Produce json
// @Param id path string true "Blog Post ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Success 200 {object} contract.GetBlogPostResponse200 "Blog post retrieved successfully"
// @Failure 404 {object} contract.GetBlogPostResponse404 "Blog post not found"
// @Failure 500 {object} contract.GetBlogPostResponse500 "Server error during retrieval"
// @Router /blog-api/api/posts/{id} [get]
func (h *GetBlogPostController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	id := c.Param("id")
	req := contract.GetBlogPostRequest{
		BlogPostId: id,
	}
	if validateResponse := services.ValidateRequest[contract.GetBlogPostResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	query := mappers.MapToGetBlogPostQuery(&req)
	resp := services.SendToMediator[*queries.GetBlogPostQuery, *contract.GetBlogPostResponse](c.Request.Context(), &query)
	responseSender.Send(resp)
}

func (h *GetBlogPostController) Route() string {
	return "/posts/:id"
}

func (h *GetBlogPostController) Methods() []string {
	return []string{"GET"}
}
