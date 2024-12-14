package controllers

import (
	"identity-api/API/services"
	contract "identity-api/Application.contract/BlogPostLikes/CreateLikeForBlogPost"
	commands "identity-api/Application/command_handlers/create_like_for_blog_post"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateLikeForBlogPostController struct {
	validator *validator.Validate
}

func NewCreateLikeForBlogPostController(validator *validator.Validate) *CreateLikeForBlogPostController {
	return &CreateLikeForBlogPostController{validator: validator}
}

// Handle godoc
// @Summary Like a blog post
// @Description Adds a like to a specific blog post
// @Tags blog,likes
// @Accept json
// @Produce json
// @Param token header string true "JWT token" example:"your.jwt.token.here"
// @Param blogPostId path string true "Blog Post ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Success 200 {object} contract.CreateLikeForBlogPostResponse200 "Like added successfully"
// @Failure 400 {object} contract.CreateLikeForBlogPostResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.CreateLikeForBlogPostResponse401 "Unauthorized"
// @Failure 404 {object} contract.CreateLikeForBlogPostResponse404 "Blog post not found"
// @Failure 500 {object} contract.CreateLikeForBlogPostResponse500 "Server error during creation"
// @Router /blog-api/api/posts/{blogPostId}/likes [post]
func (h *CreateLikeForBlogPostController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)

	blogPostId := c.Param("blogPostId")
	if blogPostId == "" {
		response := contract.CreateLikeForBlogPostResponse{
			BaseResponse: services.NewErrorResponse(400, "Blog post ID is required"),
		}
		responseSender.Send(&response)
		return
	}

	command := commands.CreateLikeForBlogPostCommand{
		BlogPostId: blogPostId,
		JwtToken:   services.GetJwtTokenFromHeader(c),
	}

	resp := services.SendToMediator[*commands.CreateLikeForBlogPostCommand, *contract.CreateLikeForBlogPostResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *CreateLikeForBlogPostController) Route() string {
	return "/posts/:blogPostId/likes"
}

func (h *CreateLikeForBlogPostController) Methods() []string {
	return []string{"POST"}
}
