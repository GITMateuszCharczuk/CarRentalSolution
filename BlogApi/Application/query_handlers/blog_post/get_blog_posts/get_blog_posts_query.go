package queries

import (
	models "identity-api/Domain/models/external"
	"identity-api/Domain/pagination"
	"identity-api/Domain/sorting"
)

type GetBlogPostsQuery struct {
	pagination.Pagination `json:",inline"`
	sorting.Sortable      `json:",inline"`
	Ids                   []string `json:"ids"`
	DateTimeFrom          string   `json:"dateTimeFrom"`
	DateTimeTo            string   `json:"dateTimeTo"`
	AuthorIds             []string `json:"authorIds"`
	Tags                  []string `json:"tags"`
	Visible               bool     `json:"visible"`
	models.JwtToken       `json:",inline"`
}
