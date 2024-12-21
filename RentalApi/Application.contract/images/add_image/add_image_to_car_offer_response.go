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

type AddUrlToCarOfferResponse400 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid request format" swaggertype:"string"`
	Id      string `json:"id" example:"" swaggertype:"string"`
}

type AddUrlToCarOfferResponse401 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Unauthorized" swaggertype:"string"`
	Id      string `json:"id" example:"" swaggertype:"string"`
}

type AddUrlToCarOfferResponse403 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Forbidden - Not authorized" swaggertype:"string"`
	Id      string `json:"id" example:"" swaggertype:"string"`
}

type AddUrlToCarOfferResponse404 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Car offer not found" swaggertype:"string"`
	Id      string `json:"id" example:"" swaggertype:"string"`
}

type AddUrlToCarOfferResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Server error during addition" swaggertype:"string"`
	Id      string `json:"id" example:"" swaggertype:"string"`
}
