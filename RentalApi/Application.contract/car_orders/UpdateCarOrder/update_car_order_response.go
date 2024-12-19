package contract

import (
	responses "rental-api/Domain/responses"
)

type UpdateCarOrderResponse struct {
	responses.BaseResponse
}

type UpdateCarOrderResponse200 struct {
	Success bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message string `json:"message" example:"Car order updated successfully" swaggertype:"string"`
}

type UpdateCarOrderResponse400 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid request parameters" swaggertype:"string"`
}

type UpdateCarOrderResponse401 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Unauthorized" swaggertype:"string"`
}

type UpdateCarOrderResponse404 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Car order not found" swaggertype:"string"`
}

type UpdateCarOrderResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Internal server error while updating car order" swaggertype:"string"`
}
