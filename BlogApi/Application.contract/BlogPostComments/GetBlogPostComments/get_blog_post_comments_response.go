package contract

import (
	"identity-api/Domain/models"
	"identity-api/Domain/pagination"
	responses "identity-api/Domain/responses"
)

type GetBlogPostCommentsResponse struct {
	responses.BaseResponse
	pagination.PaginatedResult[models.BlogPostCommentModel] `json:",inline"`
}

type GetBlogPostCommentsResponse200 struct {
	Success        bool                       `json:"success" example:"true" swaggertype:"boolean"`
	Message        string                     `json:"message" example:"Comments retrieved successfully" swaggertype:"string"`
	Page           *int                       `json:"page" example:"1" swaggertype:"integer"`
	PageSize       *int                       `json:"pageSize" example:"10" swaggertype:"integer"`
	TotalCount     int                        `json:"totalCount" example:"100" swaggertype:"integer"`
	OrderBy        *BlogPostCommentSortColumn `json:"orderBy" example:"dateAdded" swaggertype:"string"`
	OrderDirection *SortOrder                 `json:"orderDirection" example:"desc" swaggertype:"string"`
	Items          []BlogPostCommentModel     `json:"items" swaggertype:"array,object"`
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
