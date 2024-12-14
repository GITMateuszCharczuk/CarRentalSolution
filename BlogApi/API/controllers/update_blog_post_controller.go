package controllers

import (
	"identity-api/API/mappers"
	"identity-api/API/services"
	contract "identity-api/Application.contract/BlogPosts/UpdateBlogPost"
	commands "identity-api/Application/command_handlers/update_blog_post"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UpdateBlogPostController struct {
	validator *validator.Validate
}

func NewUpdateBlogPostController(validator *validator.Validate) *UpdateBlogPostController {
	return &UpdateBlogPostController{validator: validator}
}

// Handle godoc
// @Summary Update a blog post
// @Description Updates an existing blog post with the provided details
// @Tags blog
// @Accept json
// @Produce json
// @Param token header string true "JWT token" example:"your.jwt.token.here"
// @Param id path string true "Blog Post ID" example:"123e4567-e89b-12d3-a456-426614174000"
// @Param request body contract.UpdateBlogPostRequest true "Updated blog post details"
// @Success 200 {object} contract.UpdateBlogPostResponse200 "Blog post updated successfully"
// @Failure 400 {object} contract.UpdateBlogPostResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.UpdateBlogPostResponse401 "Unauthorized"
// @Failure 403 {object} contract.UpdateBlogPostResponse403 "Forbidden - Not the post owner"
// @Failure 404 {object} contract.UpdateBlogPostResponse404 "Blog post not found"
// @Failure 500 {object} contract.UpdateBlogPostResponse500 "Server error during update"
// @Router /blog-api/api/posts/{id} [put]
func (h *UpdateBlogPostController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)

	id := c.Param("id")
	if id == "" {
		response := contract.UpdateBlogPostResponse{
			BaseResponse: services.NewErrorResponse(400, "Blog post ID is required"),
		}
		responseSender.Send(&response)
		return
	}

	var req contract.UpdateBlogPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response := contract.UpdateBlogPostResponse{
			BaseResponse: services.NewErrorResponse(400, "Invalid request format"),
		}
		responseSender.Send(&response)
		return
	}

	req.ID = id
	req.JwtToken = services.GetJwtTokenFromHeader(c)

	if validateResponse := services.ValidateRequest[contract.UpdateBlogPostResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}

	command := mappers.MapToUpdateBlogPostCommand(&req)
	resp := services.SendToMediator[*commands.UpdateBlogPostCommand, *contract.UpdateBlogPostResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *UpdateBlogPostController) Route() string {
	return "/posts/:id"
}

func (h *UpdateBlogPostController) Methods() []string {
	return []string{"PUT"}
}
