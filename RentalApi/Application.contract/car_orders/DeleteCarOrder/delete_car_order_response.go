package contract

import (
	responses "rental-api/Domain/responses"
)

type DeleteCarOrderResponse struct {
	responses.BaseResponse
}

type DeleteCarOrderResponse200 struct {
	Success bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message string `json:"message" example:"Car order deleted successfully" swaggertype:"string"`
}

type DeleteCarOrderResponse400 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid request parameters" swaggertype:"string"`
}

type DeleteCarOrderResponse401 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Unauthorized" swaggertype:"string"`
}

type DeleteCarOrderResponse404 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Car order not found" swaggertype:"string"`
}

type DeleteCarOrderResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Internal server error while deleting car order" swaggertype:"string"`
}
