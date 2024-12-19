package contract

import models "rental-api/Domain/models/external"

type CreateCarOrderRequest struct {
	UserId           string  `json:"userId" binding:"required" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string" validate:"required"`
	CarOfferId       string  `json:"carOfferId" binding:"required" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string" validate:"required"`
	StartDate        string  `json:"startDate" binding:"required" example:"2023-12-12" swaggertype:"string" validate:"required"`
	EndDate          string  `json:"endDate" binding:"required" example:"2023-12-19" swaggertype:"string" validate:"required"`
	DeliveryLocation string  `json:"deliveryLocation" example:"City Center" swaggertype:"string"`
	ReturnLocation   string  `json:"returnLocation" example:"City Center" swaggertype:"string"`
	NumOfDrivers     int     `json:"numOfDrivers" example:"2" swaggertype:"integer"`
	TotalCost        float64 `json:"totalCost" example:"750.00" swaggertype:"number"`
	Status           string  `json:"status" example:"pending" swaggertype:"string"`
	models.JwtToken  `json:",inline"`
}
