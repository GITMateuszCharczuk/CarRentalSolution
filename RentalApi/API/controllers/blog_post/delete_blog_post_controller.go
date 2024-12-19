package controllers

import (
	"rental-api/API/mappers"
	"rental-api/API/services"
	contract "rental-api/Application.contract/BlogPosts/DeleteBlogPost"
	commands "rental-api/Application/command_handlers/blog_post/delete_blog_post"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DeleteBlogPostController struct {
	validator *validator.Validate
}

func NewDeleteBlogPostController(validator *validator.Validate) *DeleteBlogPostController {
	return &DeleteBlogPostController{validator: validator}
}

// Handle godoc
// @Summary Delete a blog post
// @Description Deletes an existing blog post
// @Tags blog
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Param id path string true "Blog Post ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Success 200 {object} contract.DeleteBlogPostResponse200 "Blog post deleted successfully"
// @Failure 400 {object} contract.DeleteBlogPostResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.DeleteBlogPostResponse401 "Unauthorized"
// @Failure 403 {object} contract.DeleteBlogPostResponse403 "Forbidden - Not authorized to delete"
// @Failure 404 {object} contract.DeleteBlogPostResponse404 "Blog post not found"
// @Failure 500 {object} contract.DeleteBlogPostResponse500 "Server error during deletion"
// @Router /rental-api/api/posts/{id} [delete]
func (h *DeleteBlogPostController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	token := services.GetJwtTokenFromQuery(c)
	id := c.Param("id")
	req := contract.DeleteBlogPostRequest{
		BlogPostId: id,
		JwtToken:   token,
	}
	if validateResponse := services.ValidateRequest[contract.DeleteBlogPostResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	command := mappers.MapToDeleteBlogPostCommand(&req)
	resp := services.SendToMediator[*commands.DeleteBlogPostCommand, *contract.DeleteBlogPostResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *DeleteBlogPostController) Route() string {
	return "/posts/:id"
}

func (h *DeleteBlogPostController) Methods() []string {
	return []string{"DELETE"}
}
