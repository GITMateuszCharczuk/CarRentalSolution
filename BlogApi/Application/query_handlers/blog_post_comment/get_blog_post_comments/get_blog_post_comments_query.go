package queries

import (
	"identity-api/Domain/pagination"
	"identity-api/Domain/sorting"
)

type GetBlogPostCommentsQuery struct {
	pagination.Pagination `json:",inline"`
	sorting.Sortable      `json:",inline"`
	BlogPostIds           []string `json:"blogPostIds" example:"[\"123e4567-e89b-12d3-a456-426614174000\"]" swaggertype:"array,string"`
	DateTimeFrom          string   `json:"dateTimeFrom" example:"2023-12-12T00:00:00Z" swaggertype:"string"`
	DateTimeTo            string   `json:"dateTimeTo" example:"2023-12-12T23:59:59Z" swaggertype:"string"`
	UserIds               []string `json:"userIds" example:"[\"456e4567-e89b-12d3-a456-426614174000\"]" swaggertype:"array,string"`
}
