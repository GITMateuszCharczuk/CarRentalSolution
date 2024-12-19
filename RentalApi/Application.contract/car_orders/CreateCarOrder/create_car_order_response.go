package contract

import (
	responses "rental-api/Domain/responses"
)

type CreateCarOrderResponse struct {
	responses.BaseResponse
	Id string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
}

type CreateCarOrderResponse200 struct {
	Success bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message string `json:"message" example:"Car order created successfully" swaggertype:"string"`
	Id      string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
}

type CreateCarOrderResponse400 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid request parameters" swaggertype:"string"`
	Id      string `json:"id" example:"" swaggertype:"string"`
}

type CreateCarOrderResponse401 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Unauthorized" swaggertype:"string"`
	Id      string `json:"id" example:"" swaggertype:"string"`
}

type CreateCarOrderResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Internal server error while creating car order" swaggertype:"string"`
	Id      string `json:"id" example:"" swaggertype:"string"`
}
