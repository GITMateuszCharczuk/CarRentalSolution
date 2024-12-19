package repository_interfaces

import (
	models "rental-api/Domain/models/domestic"
	"rental-api/Domain/pagination"
	"rental-api/Domain/sorting"
)

type CarOfferQueryRepository interface {
	GetCarOfferByID(id string) (*models.CarOfferModel, error)
	GetCarOffers(
		pagination *pagination.Pagination,
		sorting *sorting.Sortable,
		ids []string,
		dateTimeFrom string,
		dateTimeTo string,
		tags []string,
		visible string,
	) (*pagination.PaginatedResult[models.CarOfferModel], error)
}
