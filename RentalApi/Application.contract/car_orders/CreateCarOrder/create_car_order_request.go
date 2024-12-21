package contract

import models "rental-api/Domain/models/external"

type CreateCarOrderRequest struct {
	CarOfferId       string  `json:"carOfferId" binding:"required" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string" validate:"required"`
	StartDate        string  `json:"startDate" binding:"required" example:"2024-03-15T14:30:00Z" swaggertype:"string" validate:"required,datetime,futuredate"`
	EndDate          string  `json:"endDate" binding:"required" example:"2024-03-20T12:00:00Z" swaggertype:"string" validate:"required,datetime,gtdate=StartDate"`
	DeliveryLocation string  `json:"deliveryLocation" example:"City Center" swaggertype:"string"`
	ReturnLocation   string  `json:"returnLocation" example:"City Center" swaggertype:"string"`
	NumOfDrivers     int     `json:"numOfDrivers" example:"2" swaggertype:"integer"`
	TotalCost        float64 `json:"totalCost" example:"750.00" swaggertype:"number"`
	Status           string  `json:"status" example:"pending" swaggertype:"string" validate:"required,validCarOrderStatus"`
	models.JwtToken  `json:",inline"`
}
