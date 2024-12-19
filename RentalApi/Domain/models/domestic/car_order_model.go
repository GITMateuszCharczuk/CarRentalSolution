package models

type CarOrderModel struct {
	Id               string  `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	UserId           string  `json:"userId" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	UserEmail        string  `json:"userEmail" example:"user@example.com" swaggertype:"string"`
	CarOfferId       string  `json:"carOfferId" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	StartDate        string  `json:"startDate" example:"2023-12-12" swaggertype:"string"`
	EndDate          string  `json:"endDate" example:"2023-12-19" swaggertype:"string"`
	DeliveryLocation string  `json:"deliveryLocation" example:"Airport" swaggertype:"string"`
	ReturnLocation   string  `json:"returnLocation" example:"Airport" swaggertype:"string"`
	NumOfDrivers     int     `json:"numOfDrivers" example:"2" swaggertype:"integer"`
	TotalCost        float64 `json:"totalCost" example:"750.00" swaggertype:"number"`
	Status           string  `json:"status" example:"PENDING" swaggertype:"string"`
}
