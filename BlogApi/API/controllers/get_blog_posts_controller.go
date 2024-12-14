package controllers

import (
	"identity-api/API/mappers"
	"identity-api/API/services"
	contract "identity-api/Application.contract/BlogPosts/GetBlogPosts"
	queries "identity-api/Application/query_handlers/get_blog_posts"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GetBlogPostsController struct {
	validator *validator.Validate
}

func NewGetBlogPostsController(validator *validator.Validate) *GetBlogPostsController {
	return &GetBlogPostsController{validator: validator}
}

// Handle godoc
// @Summary Get blog posts
// @Description Retrieves a list of blog posts with optional filtering, pagination and sorting.
// @Tags blog
// @Accept json
// @Produce json
// @Param token query string true "JWT token" example:"your.jwt.token.here"
// @Param page_size query int false "Page size" example:"10"
// @Param current_page query int false "Current page" example:"1"
// @Param sort_query query []string false "Sort fields" example:"createdAt:desc"
// @Param ids query []string false "Blog post IDs" example:"id1,id2"
// @Param dateTimeFrom query string false "Start date" example:"2023-12-12T00:00:00Z"
// @Param dateTimeTo query string false "End date" example:"2023-12-12T23:59:59Z"
// @Param authorIds query []string false "Author IDs" example:"author1,author2"
// @Param tags query []string false "Tags" example:"tag1,tag2"
// @Param visible query bool false "Visibility status" example:"true"
// @Success 200 {object} contract.GetBlogPostsResponse200 "Blog posts retrieved successfully"
// @Failure 400 {object} contract.GetBlogPostsResponse400 "Invalid request parameters"
// @Failure 401 {object} contract.GetBlogPostsResponse401 "Unauthorized"
// @Failure 500 {object} contract.GetBlogPostsResponse500 "Server error during retrieval"
// @Router /blog-api/api/posts [get]
func (h *GetBlogPostsController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)

	req := contract.GetBlogPostsRequest{
		JwtToken:     services.GetJwtTokenFromQuery(c),
		Pagination:   services.ExtractPagination(c),
		SortQuery:    services.ExtractArrayFromQuery(c, "sort_query"),
		Ids:          services.ExtractArrayFromQuery(c, "ids"),
		DateTimeFrom: c.Query("dateTimeFrom"),
		DateTimeTo:   c.Query("dateTimeTo"),
		AuthorIds:    services.ExtractArrayFromQuery(c, "authorIds"),
		Tags:         services.ExtractArrayFromQuery(c, "tags"),
		Visible:      services.ExtractBoolFromQuery(c, "visible", false),
	}

	if validateResponse := services.ValidateRequest[contract.GetBlogPostsResponse](&req, h.validator); validateResponse != nil {
		responseSender.Send(validateResponse)
		return
	}

	query := mappers.MapToGetBlogPostsQuery(&req)
	resp := services.SendToMediator[*queries.GetBlogPostsQuery, *contract.GetBlogPostsResponse](c.Request.Context(), &query)
	responseSender.Send(resp)
}

func (h *GetBlogPostsController) Route() string {
	return "/posts"
}

func (h *GetBlogPostsController) Methods() []string {
	return []string{"GET"}
}
