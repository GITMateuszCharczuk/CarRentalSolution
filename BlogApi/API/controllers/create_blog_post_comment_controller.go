package controllers

import (
	"identity-api/API/mappers"
	"identity-api/API/services"
	contract "identity-api/Application.contract/BlogPostComments/CreateBlogPostComment"
	commands "identity-api/Application/command_handlers/create_blog_post_comment"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateBlogPostCommentController struct {
	validator *validator.Validate
}

func NewCreateBlogPostCommentController(validator *validator.Validate) *CreateBlogPostCommentController {
	return &CreateBlogPostCommentController{validator: validator}
}

// Handle godoc
// @Summary Create a comment on a blog post
// @Description Creates a new comment on a specific blog post
// @Tags blog,comments
// @Accept json
// @Produce json
// @Param token header string true "JWT token" example:"your.jwt.token.here"
// @Param blogPostId path string true "Blog Post ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Param request body contract.CreateBlogPostCommentRequest true "Comment details"
// @Success 200 {object} contract.CreateBlogPostCommentResponse200 "Comment created successfully"
// @Failure 400 {object} contract.CreateBlogPostCommentResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.CreateBlogPostCommentResponse401 "Unauthorized"
// @Failure 404 {object} contract.CreateBlogPostCommentResponse404 "Blog post not found"
// @Failure 500 {object} contract.CreateBlogPostCommentResponse500 "Server error during creation"
// @Router /blog-api/api/posts/{blogPostId}/comments [post]
func (h *CreateBlogPostCommentController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)

	blogPostId := c.Param("blogPostId")
	if blogPostId == "" {
		response := contract.CreateBlogPostCommentResponse{
			BaseResponse: services.NewErrorResponse(400, "Blog post ID is required"),
		}
		responseSender.Send(&response)
		return
	}

	var req contract.CreateBlogPostCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response := contract.CreateBlogPostCommentResponse{
			BaseResponse: services.NewErrorResponse(400, "Invalid request format"),
		}
		responseSender.Send(&response)
		return
	}

	req.BlogPostId = blogPostId
	req.JwtToken = services.GetJwtTokenFromHeader(c)

	if validateResponse := services.ValidateRequest[contract.CreateBlogPostCommentResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}

	command := mappers.MapToCreateBlogPostCommentCommand(&req)
	resp := services.SendToMediator[*commands.CreateBlogPostCommentCommand, *contract.CreateBlogPostCommentResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *CreateBlogPostCommentController) Route() string {
	return "/posts/:blogPostId/comments"
}

func (h *CreateBlogPostCommentController) Methods() []string {
	return []string{"POST"}
}
