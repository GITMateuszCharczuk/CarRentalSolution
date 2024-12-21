package contract

import (
	models "rental-api/Domain/models/external"
	"rental-api/Domain/pagination"
)

type GetCarOrdersRequest struct {
	pagination.Pagination `json:",inline"`
	SortQuery             []string `json:"sort_query" validate:"validCarOrderSortable"`
	Dates                 []string `json:"dates" example:"[\"2023-12-12T00:00:00Z\"]" swaggertype:"array,string"`
	UserId                string   `json:"userId" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	StartDate             string   `json:"startDate" example:"2023-12-12T00:00:00Z" swaggertype:"string"`
	EndDate               string   `json:"endDate" example:"2023-12-12T00:00:00Z" swaggertype:"string"`
	CarOfferId            string   `json:"carOfferId" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	Statuses              []string `json:"statuses" example:"PENDING,ACTIVE,OVERDUE" swaggertype:"array,string" validate:"validCarOrderStatusArray"`
	DateFilterType        string   `json:"dateFilterType" example:"created" swaggertype:"string"`
	models.JwtToken       `json:",inline"`
}
