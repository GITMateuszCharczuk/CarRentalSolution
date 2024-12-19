package contract

import (
	responses "rental-api/Domain/responses"
)

type UpdateCarOfferResponse struct {
	responses.BaseResponse
}

type UpdateCarOfferResponse200 struct {
	Success bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message string `json:"message" example:"Car offer updated successfully" swaggertype:"string"`
}

type UpdateCarOfferResponse400 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid request parameters" swaggertype:"string"`
}

type UpdateCarOfferResponse401 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Unauthorized" swaggertype:"string"`
}

type UpdateCarOfferResponse404 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Car offer not found" swaggertype:"string"`
}

type UpdateCarOfferResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Internal server error while updating car offer" swaggertype:"string"`
}
