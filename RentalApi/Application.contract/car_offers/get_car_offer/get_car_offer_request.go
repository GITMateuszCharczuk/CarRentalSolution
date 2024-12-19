package contract

type GetCarOfferRequest struct {
	CarOfferId string `json:"carOfferId" validate:"required,uuid" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
}
