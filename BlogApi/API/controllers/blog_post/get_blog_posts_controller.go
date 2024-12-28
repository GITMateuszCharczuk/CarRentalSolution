package controllers

import (
	"blog-api/API/mappers"
	"blog-api/API/services"
	contract "blog-api/Application.contract/BlogPosts/GetBlogPosts"
	queries "blog-api/Application/query_handlers/blog_post/get_blog_posts"

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
// @Param page_size query int false "Page size" example:"10"
// @Param current_page query int false "Current page" example:"1"
// @Param sort_fields query []string false "Sort fields (field:asc|desc)" example:"created_at:desc"
// @Param ids query []string false "Blog post IDs" example:"id1,id2"
// @Param date-time-from query string false "Start date" example:"2023-12-12T00:00:00Z"
// @Param date-time-to query string false "End date" example:"2023-12-12T23:59:59Z"
// @Param author-ids query []string false "Author IDs" example:"author1,author2"
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
		Ids:          services.ExtractQueryArray(c, "ids"), //TODO
		DateTimeFrom: c.Query("date-time-from"),
		DateTimeTo:   c.Query("date-time-to"),
		AuthorIds:    services.ExtractQueryArray(c, "author-ids"),
		Tags:         services.ExtractQueryArray(c, "tags"),
		Visible:      c.Query("visible"),
		SortQuery:    services.ExtractSortQuery(c),
		Pagination:   services.ExtractPagination(c),
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
