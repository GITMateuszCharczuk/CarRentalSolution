package contract

import (
	models "identity-api/Domain/models/domestic"
	"identity-api/Domain/pagination"
	responses "identity-api/Domain/responses"
)

type GetBlogPostCommentsResponse struct {
	responses.BaseResponse
	pagination.PaginatedResult[models.BlogPostCommentModel] `json:",inline"`
}

type GetBlogPostCommentsResponse200 struct {
	Success     bool                          `json:"success" example:"true" swaggertype:"boolean"`
	Message     string                        `json:"message" example:"Comments retrieved successfully" swaggertype:"string"`
	TotalItems  int                           `json:"total_items" example:"100"`
	CurrentPage int                           `json:"current_page" example:"1"`
	PageSize    int                           `json:"page_size" example:"10"`
	TotalPages  int                           `json:"total_pages" example:"10"`
	Items       []models.BlogPostCommentModel `json:"items" swaggertype:"array,object"` //todo
}

type GetBlogPostCommentsResponse400 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid request parameters" swaggertype:"string"`
}

type GetBlogPostCommentsResponse404 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Blog post not found" swaggertype:"string"`
}

type GetBlogPostCommentsResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Internal server error while retrieving comments" swaggertype:"string"`
}
