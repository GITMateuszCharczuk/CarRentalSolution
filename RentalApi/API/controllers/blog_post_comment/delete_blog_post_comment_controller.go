package controllers

import (
	"rental-api/API/mappers"
	"rental-api/API/services"
	contract "rental-api/Application.contract/BlogPostComments/DeketeBlogPostComment"
	commands "rental-api/Application/command_handlers/blog_post_comment/delete_blog_post_comment"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DeleteBlogPostCommentController struct {
	validator *validator.Validate
}

func NewDeleteBlogPostCommentController(validator *validator.Validate) *DeleteBlogPostCommentController {
	return &DeleteBlogPostCommentController{validator: validator}
}

// Handle godoc
// @Summary Delete a comment from a blog post
// @Description Deletes a specific comment from a blog post
// @Tags comments
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Param id path string true "Comment ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Success 200 {object} contract.DeleteBlogPostCommentResponse200 "Comment deleted successfully"
// @Failure 400 {object} contract.DeleteBlogPostCommentResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.DeleteBlogPostCommentResponse401 "Unauthorized"
// @Failure 403 {object} contract.DeleteBlogPostCommentResponse403 "Forbidden - Not the comment owner"
// @Failure 404 {object} contract.DeleteBlogPostCommentResponse404 "Comment not found"
// @Failure 500 {object} contract.DeleteBlogPostCommentResponse500 "Server error during deletion"
// @Router /rental-api/api/comments/{id} [delete]
func (h *DeleteBlogPostCommentController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	token := services.GetJwtTokenFromQuery(c)
	id := c.Param("id")
	req := contract.DeleteBlogPostCommentRequest{
		BlogPostCommentId: id,
		JwtToken:          token,
	}
	if validateResponse := services.ValidateRequest[contract.DeleteBlogPostCommentResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}
	command := mappers.MapToDeleteBlogPostCommentCommand(&req)
	resp := services.SendToMediator[*commands.DeleteBlogPostCommentCommand, *contract.DeleteBlogPostCommentResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *DeleteBlogPostCommentController) Route() string {
	return "/comments/:id"
}

func (h *DeleteBlogPostCommentController) Methods() []string {
	return []string{"DELETE"}
}
