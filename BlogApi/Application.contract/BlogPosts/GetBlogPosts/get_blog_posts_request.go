package contract

import (
	models "identity-api/Domain/models/external"
	"identity-api/Domain/pagination"
)

type GetBlogPostsRequest struct {
	pagination.Pagination `json:",inline"`
	SortQuery             []string `json:"sort_query" validate:"validBlogPostSortable"`
	Ids                   []string `json:"ids" example:"[\"123e4567-e89b-12d3-a456-426614174000\"]" swaggertype:"array,string"`
	PublishedDates        []string `json:"publishedDates" example:"[\"2023-12-12\"]" swaggertype:"array,string"`
	Authors               []string `json:"authors" example:"[\"John Doe\"]" swaggertype:"array,string"`
	Tags                  []string `json:"tags" example:"[\"Technology\"]" swaggertype:"array,string"`
	models.JwtToken       `json:",inline"`
}
