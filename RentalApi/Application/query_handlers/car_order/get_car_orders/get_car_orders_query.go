package queries

import (
	models "rental-api/Domain/models/external"
	"rental-api/Domain/pagination"
	"rental-api/Domain/sorting"
)

type GetCarOrdersQuery struct {
	pagination.Pagination `json:",inline"`
	sorting.Sortable      `json:",inline"`
	models.JwtToken       `json:",inline"`
	StartDate             string   `json:"startDate" example:"2023-12-12T00:00:00Z" swaggertype:"string"`
	EndDate               string   `json:"endDate" example:"2023-12-19T23:59:59Z" swaggertype:"string"`
	UserId                string   `json:"userId" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	CarOfferId            string   `json:"carOfferId" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	Statuses              []string `json:"statuses" example:"PENDING,ACTIVE,OVERDUE" swaggertype:"array,string"`
	DateFilterType        string   `json:"dateFilterType" example:"created" swaggertype:"string"`
}
