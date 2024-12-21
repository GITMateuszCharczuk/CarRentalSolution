package contract

import (
	responses "rental-api/Domain/responses"
)

type DeleteCarOfferResponse struct {
	responses.BaseResponse
}

type DeleteCarOfferResponse200 struct {
	Success bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message string `json:"message" example:"Car offer deleted successfully" swaggertype:"string"`
}

type DeleteCarOfferResponse400 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid request parameters" swaggertype:"string"`
}

type DeleteCarOfferResponse401 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Unauthorized" swaggertype:"string"`
}

type DeleteCarOfferResponse403 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Forbidden - Not authorized" swaggertype:"string"`
}

type DeleteCarOfferResponse404 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Car offer not found" swaggertype:"string"`
}

type DeleteCarOfferResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Internal server error while deleting car offer" swaggertype:"string"`
}
