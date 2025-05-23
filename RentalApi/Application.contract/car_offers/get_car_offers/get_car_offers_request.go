package contract

import (
	"rental-api/Domain/pagination"
)

type GetCarOffersRequest struct {
	pagination.Pagination `json:",inline"`
	Ids                   []string `json:"ids" example:"[\"123e4567-e89b-12d3-a456-426614174000\"]" swaggertype:"array,string"`
	SortQuery             []string `json:"sort_query" validate:"validCarOfferSortable" swaggertype:"array,string"`
	DateTimeFrom          string   `json:"dateTimeFrom" example:"2023-12-12T00:00:00Z" validate:"datetime,futuredate" swaggertype:"string"`
	DateTimeTo            string   `json:"dateTimeTo" example:"2023-12-12T23:59:59Z" validate:"datetime,gtdate=DateTimeFrom" swaggertype:"string"`
	Tags                  []string `json:"tags" example:"[\"luxury\",\"sports\"]" swaggertype:"array,string"`
	Visible               string   `json:"visible" example:"true" swaggertype:"string" validate:"omitempty,oneof=true false"`
}
