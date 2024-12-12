package contract

import (
	models "identity-api/Domain/models/token"
	"identity-api/Domain/pagination"
)

type BlogPostSortColumn string
type SortOrder string

const (
	SortByHeading   BlogPostSortColumn = "heading"
	SortByAuthor    BlogPostSortColumn = "author"
	SortByPublished BlogPostSortColumn = "publishedDate"
	SortByCreated   BlogPostSortColumn = "createdAt"

	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
)

type GetBlogPostsRequest struct {
	pagination.Pagination `json:",inline"`
	SortQuery             []string `json:"sort_query" validate:"validBlogPostSortable"`
	models.JwtToken       `json:",inline"`
	Ids                   []string `json:"ids" example:"[\"123e4567-e89b-12d3-a456-426614174000\"]" swaggertype:"array,string"`
	PublishedDates        []string `json:"publishedDates" example:"[\"2023-12-12\"]" swaggertype:"array,string"`
	Authors               []string `json:"authors" example:"[\"John Doe\"]" swaggertype:"array,string"`
}
