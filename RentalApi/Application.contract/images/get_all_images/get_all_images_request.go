package contract

type GetAllImagesRequest struct {
	CarOfferId string `json:"carOfferId" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string" validate:"required"`
}
