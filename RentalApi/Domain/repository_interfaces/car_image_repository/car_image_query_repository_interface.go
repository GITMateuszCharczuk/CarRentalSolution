package repository_interfaces

import (
	models "rental-api/Domain/models/domestic"
)

type CarImageQueryRepository interface {
	GetImagesByCarOfferId(carOfferId string) (*[]models.CarOfferImageModel, error)
	GetImageById(id string) (*models.CarOfferImageModel, error)
}
