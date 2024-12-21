package contract

import (
	models "rental-api/Domain/models/domestic"
	"rental-api/Domain/pagination"
	responses "rental-api/Domain/responses"
)

type GetCarOrdersResponse struct {
	responses.BaseResponse
	pagination.PaginatedResult[models.CarOrderModel] `json:",inline"`
}

type GetCarOrdersResponse200 struct {
	Success     bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message     string `json:"message" example:"Car orders retrieved successfully" swaggertype:"string"`
	TotalItems  int    `json:"total_items" example:"100"`
	CurrentPage int    `json:"current_page" example:"1"`
	PageSize    int    `json:"page_size" example:"10"`
	TotalPages  int    `json:"total_pages" example:"10"`
	Items       []struct {
		Id               string  `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
		UserId           string  `json:"userId" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
		CarOfferId       string  `json:"carOfferId" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
		StartDate        string  `json:"startDate" example:"2023-12-12" swaggertype:"string"`
		EndDate          string  `json:"endDate" example:"2023-12-19" swaggertype:"string"`
		DeliveryLocation string  `json:"deliveryLocation" example:"City Center" swaggertype:"string"`
		ReturnLocation   string  `json:"returnLocation" example:"City Center" swaggertype:"string"`
		NumOfDrivers     int     `json:"numOfDrivers" example:"2" swaggertype:"integer"`
		TotalCost        float64 `json:"totalCost" example:"750.00" swaggertype:"number"`
	} `json:"items"`
}

type GetCarOrdersResponse400 struct {
	Success bool                   `json:"success" example:"false" swaggertype:"boolean"`
	Message string                 `json:"message" example:"Invalid request parameters" swaggertype:"string"`
	Items   []models.CarOrderModel `json:"items" swaggertype:"array,object"`
}

type GetCarOrdersResponse401 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Unauthorized" swaggertype:"string"`
}

type GetCarOrdersResponse403 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Forbidden - Not authorized" swaggertype:"string"`
}

type GetCarOrdersResponse500 struct {
	Success bool                   `json:"success" example:"false" swaggertype:"boolean"`
	Message string                 `json:"message" example:"Internal server error while retrieving car orders" swaggertype:"string"`
	Items   []models.CarOrderModel `json:"items" swaggertype:"array,object"`
}
