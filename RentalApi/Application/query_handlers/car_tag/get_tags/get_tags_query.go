package queries

import (
	"rental-api/Domain/sorting"
)

type GetTagsQuery struct {
	sorting.Sortable `json:",inline"`
	CarOfferId       string `json:"carOfferId" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string" validate:"omitempty,len=36"`
}
