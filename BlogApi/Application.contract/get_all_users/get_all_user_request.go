package contract

import (
	models "identity-api/Domain/models/token"
	"identity-api/Domain/pagination"
)

type GetAllUsersRequest struct {
	models.JwtToken       `json:",inline"`
	pagination.Pagination `json:",inline"`
	SortQuery             []string `json:"sort_query" validate:"validUserSortable"`
}
