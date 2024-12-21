package contract

import (
	responses "rental-api/Domain/responses"
)

type DeleteImageFromCarOfferResponse struct {
	responses.BaseResponse
}

type DeleteImageFromCarOfferResponse200 struct {
	Success bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message string `json:"message" example:"Image added successfully" swaggertype:"string"`
}

type DeleteImageFromCarOfferResponse400 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Image not found" swaggertype:"string"`
}

type DeleteImageFromCarOfferResponse401 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Unauthorized" swaggertype:"string"`
}

type DeleteImageFromCarOfferResponse403 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Forbidden - Not authorized" swaggertype:"string"`
}

type DeleteImageFromCarOfferResponse404 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Car offer not found" swaggertype:"string"`
}

type DeleteImageFromCarOfferResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Internal server error" swaggertype:"string"`
}
