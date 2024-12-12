package contract

import "identity-api/Domain/pagination"

type GetBlogPostCommentsRequest struct {
	pagination.Pagination `json:",inline"`
	SortQuery             []string `json:"sort_query" validate:"validCommentSortable"`
	BlogPostId            string   `json:"blogPostId" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	DateTimes             []string `json:"dateTimes" example:"[\"2023-12-12T10:00:00Z\"]" swaggertype:"array,string"`
	UserIds               []string `json:"userIds" example:"[\"456e4567-e89b-12d3-a456-426614174000\"]" swaggertype:"array,string"`
}
