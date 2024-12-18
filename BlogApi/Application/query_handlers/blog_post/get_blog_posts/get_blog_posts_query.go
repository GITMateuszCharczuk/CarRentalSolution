package queries

import (
	"blog-api/Domain/pagination"
	"blog-api/Domain/sorting"
)

type GetBlogPostsQuery struct {
	pagination.Pagination `json:",inline"`
	sorting.Sortable      `json:",inline"`
	Ids                   []string `json:"ids"`
	DateTimeFrom          string   `json:"dateTimeFrom"`
	DateTimeTo            string   `json:"dateTimeTo"`
	AuthorIds             []string `json:"authorIds"`
	Tags                  []string `json:"tags"`
	Visible               string   `json:"visible"`
}
