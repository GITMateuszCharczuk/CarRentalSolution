package contract

import (
	responses "rental-api/Domain/responses"
)

type AddUrlToCarOfferResponse struct {
	responses.BaseResponse
	Id string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
}

type AddUrlToCarOfferResponse200 struct {
	Success bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message string `json:"message" example:"Image added successfully" swaggertype:"string"`
	Id      string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
}

// ... other response types (400, 401, 404, 500)
