package contract

import (
	models "rental-api/Domain/models/domestic"
	responses "rental-api/Domain/responses"
)

type GetCarOrderResponse struct {
	responses.BaseResponse
	CarOrder models.CarOrderModel `json:"car_order" swaggertype:"object"`
}

type GetCarOrderResponse200 struct {
	Success  bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message  string `json:"message" example:"Car order retrieved successfully" swaggertype:"string"`
	CarOrder struct {
		Id               string  `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
		UserId           string  `json:"userId" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
		CarOfferId       string  `json:"carOfferId" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
		StartDate        string  `json:"startDate" example:"2023-12-12" swaggertype:"string"`
		DeliveryLocation string  `json:"deliveryLocation" example:"City Center" swaggertype:"string"`
		ReturnLocation   string  `json:"returnLocation" example:"City Center" swaggertype:"string"`
		EndDate          string  `json:"endDate" example:"2023-12-19" swaggertype:"string"`
		NumOfDrivers     int     `json:"numOfDrivers" example:"2" swaggertype:"integer"`
		TotalCost        float64 `json:"totalCost" example:"750.00" swaggertype:"number"`
	} `json:"car_order"`
}

type GetCarOrderResponse400 struct {
	Success  bool                 `json:"success" example:"false" swaggertype:"boolean"`
	Message  string               `json:"message" example:"Invalid request parameters" swaggertype:"string"`
	CarOrder models.CarOrderModel `json:"car_order" swaggertype:"object"`
}

type GetCarOrderResponse404 struct {
	Success  bool                 `json:"success" example:"false" swaggertype:"boolean"`
	Message  string               `json:"message" example:"Car order not found" swaggertype:"string"`
	CarOrder models.CarOrderModel `json:"car_order" swaggertype:"object"`
}

type GetCarOrderResponse500 struct {
	Success  bool                 `json:"success" example:"false" swaggertype:"boolean"`
	Message  string               `json:"message" example:"Internal server error while retrieving car order" swaggertype:"string"`
	CarOrder models.CarOrderModel `json:"car_order" swaggertype:"object"`
}
