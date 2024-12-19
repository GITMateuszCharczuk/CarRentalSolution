package contract

import (
	responses "rental-api/Domain/responses"
)

type CreateCarOfferResponse struct {
	responses.BaseResponse
	Id string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
}

type CreateCarOfferResponse200 struct {
	Success bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message string `json:"message" example:"Car offer created successfully" swaggertype:"string"`
	Id      string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
}

type CreateCarOfferResponse400 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid request parameters" swaggertype:"string"`
	Id      string `json:"id" example:"" swaggertype:"string"`
}

type CreateCarOfferResponse401 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Unauthorized" swaggertype:"string"`
	Id      string `json:"id" example:"" swaggertype:"string"`
}

type CreateCarOfferResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Internal server error while creating car offer" swaggertype:"string"`
	Id      string `json:"id" example:"" swaggertype:"string"`
}
