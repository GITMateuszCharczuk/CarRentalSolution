package controllers

import (
	"identity-api/API/mappers"
	"identity-api/API/services"
	contract "identity-api/Application.contract/BlogPostLikes/DeleteLikeForBlogPost"
	commands "identity-api/Application/command_handlers/blog_post_like/delete_like_for_blog_post"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DeleteLikeForBlogPostController struct {
	validator *validator.Validate
}

func NewDeleteLikeForBlogPostController(validator *validator.Validate) *DeleteLikeForBlogPostController {
	return &DeleteLikeForBlogPostController{validator: validator}
}

// Handle godoc
// @Summary Unlike a blog post
// @Description Removes a like from a specific blog post
// @Tags likes
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Param id path string true "Blog Post ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Success 200 {object} contract.DeleteLikeForBlogPostResponse200 "Like removed successfully"
// @Failure 400 {object} contract.DeleteLikeForBlogPostResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.DeleteLikeForBlogPostResponse401 "Unauthorized"
// @Failure 404 {object} contract.DeleteLikeForBlogPostResponse404 "Blog post not found"
// @Failure 500 {object} contract.DeleteLikeForBlogPostResponse500 "Server error during deletion"
// @Router /blog-api/api/posts/{id}/likes [delete]
func (h *DeleteLikeForBlogPostController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	blogPostId := c.Param("id")
	token := services.GetJwtTokenFromQuery(c)
	req := contract.DeleteLikeForBlogPostRequest{
		BlogPostId: blogPostId,
		JwtToken:   token,
	}
	if validateResponse := services.ValidateRequest[contract.DeleteLikeForBlogPostRequest](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	command := mappers.MapToDeleteLikeForBlogPostCommand(&req)
	resp := services.SendToMediator[*commands.DeleteLikeForBlogPostCommand, *contract.DeleteLikeForBlogPostResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *DeleteLikeForBlogPostController) Route() string {
	return "/posts/:id/likes"
}

func (h *DeleteLikeForBlogPostController) Methods() []string {
	return []string{"DELETE"}
}
