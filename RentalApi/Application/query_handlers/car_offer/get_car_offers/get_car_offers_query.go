package queries

import (
	"rental-api/Domain/pagination"
	"rental-api/Domain/sorting"
)

type GetCarOffersQuery struct {
	pagination.Pagination `json:",inline"`
	sorting.Sortable      `json:",inline"`
	Ids                   []string `json:"ids" example:"[\"123e4567-e89b-12d3-a456-426614174000\"]" swaggertype:"array,string"`
	DateTimeFrom          string   `json:"dateTimeFrom" example:"2023-12-12T00:00:00Z" swaggertype:"string"`
	DateTimeTo            string   `json:"dateTimeTo" example:"2023-12-12T23:59:59Z" swaggertype:"string"`
	Tags                  []string `json:"tags" example:"[\"luxury\",\"sports\"]" swaggertype:"array,string"`
	Visible               string   `json:"visible" example:"true" swaggertype:"string"`
	Status                string   `json:"status" example:"available" swaggertype:"string"`
}
