package controllers

import (
	"blog-api/API/mappers"
	"blog-api/API/services"
	contract "blog-api/Application.contract/BlogPostLikes/GetLikesForBlogPost"
	queries "blog-api/Application/query_handlers/blog_post_like/get_likes_for_blog_post"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GetLikesForBlogPostController struct {
	validator *validator.Validate
}

func NewGetLikesForBlogPostController(validator *validator.Validate) *GetLikesForBlogPostController {
	return &GetLikesForBlogPostController{validator: validator}
}

// Handle godoc
// @Summary Get likes count for a blog post
// @Description Retrieves the total number of likes for a specific blog post
// @Tags likes
// @Accept json
// @Produce json
// @Param id path string true "Blog Post ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Success 200 {object} contract.GetLikesForBlogPostResponse200 "Likes count retrieved successfully"
// @Failure 404 {object} contract.GetLikesForBlogPostResponse404 "Blog post not found"
// @Failure 500 {object} contract.GetLikesForBlogPostResponse500 "Server error during retrieval"
// @Router /blog-api/api/posts/{id}/likes [get]
func (h *GetLikesForBlogPostController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	blogPostId := c.Param("id")
	req := contract.GetLikesForBlogPostRequest{
		BlogPostId: blogPostId,
	}
	if validateResponse := services.ValidateRequest[contract.GetLikesForBlogPostResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	query := mappers.MapToGetLikesForBlogPostQuery(&req)
	resp := services.SendToMediator[*queries.GetLikesForBlogPostQuery, *contract.GetLikesForBlogPostResponse](c.Request.Context(), &query)
	responseSender.Send(resp)
}

func (h *GetLikesForBlogPostController) Route() string {
	return "/posts/:id/likes"
}

func (h *GetLikesForBlogPostController) Methods() []string {
	return []string{"GET"}
}
