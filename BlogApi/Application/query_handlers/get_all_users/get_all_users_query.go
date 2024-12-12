package queries

import (
	models "identity-api/Domain/models/token"
	"identity-api/Domain/pagination"
	"identity-api/Domain/sorting"
)

type GetAllUsersQuery struct {
	models.JwtToken       `json:",inline"`
	pagination.Pagination `json:",inline"`
	sorting.Sortable      `json:",inline"`
}
