package contract

import (
	"identity-api/Domain/models/external"
	"identity-api/Domain/pagination"
)

type GetAllUsersRequest struct {
	models.JwtToken       `json:",inline"`
	pagination.Pagination `json:",inline"`
	SortQuery             []string `json:"sort_query" validate:"validUserSortable"`
}
