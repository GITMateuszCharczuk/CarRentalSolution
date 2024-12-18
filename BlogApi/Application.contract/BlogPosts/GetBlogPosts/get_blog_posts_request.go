package contract

import (
	"blog-api/Domain/pagination"
)

type GetBlogPostsRequest struct {
	pagination.Pagination `json:",inline"`
	SortQuery             []string `json:"sort_query" validate:"validBlogPostSortable"`
	Ids                   []string `json:"ids" example:"[\"123e4567-e89b-12d3-a456-426614174000\"]" swaggertype:"array,string"`
	DateTimeFrom          string   `json:"dateTimeFrom" example:"2023-12-12T00:00:00Z" validate:"omitempty,datetime=2006-01-02T15:04:05Z" swaggertype:"string"`
	DateTimeTo            string   `json:"dateTimeTo" example:"2023-12-12T23:59:59Z" validate:"omitempty,datetime=2006-01-02T15:04:05Z,gtefield=DateTimeFrom" swaggertype:"string"`
	AuthorIds             []string `json:"authorIds" example:"[\"John Doe\"]" swaggertype:"array,string"`
	Tags                  []string `json:"tags" example:"[\"Technology\"]" swaggertype:"array,string"`
	Visible               string   `json:"visible" example:"true" swaggertype:"string" validate:"omitempty,oneof=true false"`
}
