package repository_interfaces

import (
	models "rental-api/Domain/models/domestic"
	"rental-api/Domain/sorting"
)

type CarTagQueryRepository interface {
	GetTagByID(id string) (*models.CarOfferTagModel, error)
	GetTagByName(name string) (*models.CarOfferTagModel, error)
	GetTagsByCarOfferId(
		carOfferId string,
		sorting sorting.Sortable,
	) (*[]models.CarOfferTagModel, error)
}
