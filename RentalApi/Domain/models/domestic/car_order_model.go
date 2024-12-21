package models

type CarOrderModel struct {
	Id               string  `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	UserId           string  `json:"userId" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	UserEmail        string  `json:"userEmail" example:"user@example.com" swaggertype:"string"`
	CarOfferId       string  `json:"car_offer_id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	StartDate        string  `json:"start_date" example:"2023-12-12" swaggertype:"string"`
	EndDate          string  `json:"end_date" example:"2023-12-19" swaggertype:"string"`
	DeliveryLocation string  `json:"delivery_location" example:"Airport" swaggertype:"string"`
	ReturnLocation   string  `json:"return_location" example:"Airport" swaggertype:"string"`
	NumOfDrivers     int     `json:"num_of_drivers" example:"2" swaggertype:"integer"`
	TotalCost        float64 `json:"total_cost" example:"750.00" swaggertype:"number"`
	Status           string  `json:"status" example:"PENDING" swaggertype:"string"`
}
