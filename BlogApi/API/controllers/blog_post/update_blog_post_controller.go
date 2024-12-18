package controllers

import (
	"blog-api/API/mappers"
	"blog-api/API/services"
	contract "blog-api/Application.contract/BlogPosts/UpdateBlogPost"
	commands "blog-api/Application/command_handlers/blog_post/update_blog_post"
	"blog-api/Domain/responses"

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
// @Param token query string true "JWT token" example:"your.jwt.token.here"
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
	token := services.GetJwtTokenFromQuery(c)
	id := services.ExtractFromPath(c, "id")
	var req contract.UpdateBlogPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response := responses.NewResponse[contract.UpdateBlogPostResponse](400, "Invalid request format")
		responseSender.Send(response)
		return
	}
	req.Id = id
	req.JwtToken = token
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
