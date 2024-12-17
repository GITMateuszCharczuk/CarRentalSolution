package controllers

import (
	"identity-api/API/mappers"
	"identity-api/API/services"
	contract "identity-api/Application.contract/BlogPosts/CreateBlogPost"
	commands "identity-api/Application/command_handlers/blog_post/create_blog_post"
	"identity-api/Domain/responses"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateBlogPostController struct {
	validator *validator.Validate
}

func NewCreateBlogPostController(validator *validator.Validate) *CreateBlogPostController {
	return &CreateBlogPostController{validator: validator}
}

// Handle godoc
// @Summary Create a new blog post
// @Description Creates a new blog post with the provided details
// @Tags blog
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Param request body contract.CreateBlogPostRequest true "Blog post details"
// @Success 200 {object} contract.CreateBlogPostResponse200 "Blog post created successfully"
// @Failure 400 {object} contract.CreateBlogPostResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.CreateBlogPostResponse401 "Unauthorized"
// @Failure 500 {object} contract.CreateBlogPostResponse500 "Server error during creation"
// @Router /blog-api/api/posts [post]
func (h *CreateBlogPostController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	token := services.GetJwtTokenFromQuery(c)
	var req contract.CreateBlogPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		response := responses.NewResponse[contract.CreateBlogPostResponse](400, "Invalid request format")
		responseSender.Send(response)
		return
	}

	if validateResponse := services.ValidateRequest[contract.CreateBlogPostResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}

	command := mappers.MapToCreateBlogPostCommand(&req)
	command.JwtToken = token
	resp := services.SendToMediator[*commands.CreateBlogPostCommand, *contract.CreateBlogPostResponse](c.Request.Context(), &command)
	responseSender.Send(resp)
}

func (h *CreateBlogPostController) Route() string {
	return "/posts"
}

func (h *CreateBlogPostController) Methods() []string {
	return []string{"POST"}
}
