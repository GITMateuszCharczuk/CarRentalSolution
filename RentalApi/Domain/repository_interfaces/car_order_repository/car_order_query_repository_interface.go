package repository_interfaces

import (
	models "rental-api/Domain/models/domestic"
	"rental-api/Domain/pagination"
	"rental-api/Domain/sorting"
)

type CarOrderQueryRepository interface {
	GetCarOrderByID(id string) (*models.CarOrderModel, error)
	GetCarOrders(
		pagination *pagination.Pagination,
		sorting *sorting.Sortable,
		startDate string,
		endDate string,
		userId string,
		carOfferId string,
		status string,
		dateFilterType string,
	) (*pagination.PaginatedResult[models.CarOrderModel], error)
}
