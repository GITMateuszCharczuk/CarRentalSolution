package repository

import (
	models "rental-api/Domain/models/domestic"
	"rental-api/Domain/pagination"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_order_repository"
	"rental-api/Domain/sorting"
	postgres_db "rental-api/Infrastructure/databases/postgres/config"
	"rental-api/Infrastructure/databases/postgres/entities"
	mappers "rental-api/Infrastructure/databases/postgres/mappers/base"
	base "rental-api/Infrastructure/databases/postgres/repository/base"
	"rental-api/Infrastructure/databases/postgres/repository/base/helpers"
	unit_of_work "rental-api/Infrastructure/databases/postgres/repository/base/unit_of_work"
)

type CarOrderQueryRepositoryImpl struct {
	*base.QueryRepository[entities.CarOrderEntity, string, models.CarOrderModel]
	mapper mappers.PersistenceMapper[entities.CarOrderEntity, models.CarOrderModel]
}

func NewCarOrderQueryRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.CarOrderEntity, models.CarOrderModel],
	uow unit_of_work.UnitOfWork,
) repository_interfaces.CarOrderQueryRepository {
	return &CarOrderQueryRepositoryImpl{
		QueryRepository: base.NewQueryRepository[entities.CarOrderEntity, string, models.CarOrderModel](postgresDatabase.DB, mapper, uow),
		mapper:          mapper,
	}
}

func (r *CarOrderQueryRepositoryImpl) GetCarOrderByID(id string) (*models.CarOrderModel, error) {
	return r.GetById(id)
}

func (r *CarOrderQueryRepositoryImpl) GetCarOrders(
	pagination *pagination.Pagination,
	sorting *sorting.Sortable,
	startDate string,
	endDate string,
	userId string,
	carOfferId string,
	status string,
	dateFilterType string,
) (*pagination.PaginatedResult[models.CarOrderModel], error) {
	query := r.ConstructBaseQuery()

	queryRecords := []helpers.QueryRecord[entities.CarOrderEntity]{}

	queryRecords = append(queryRecords, helpers.NewQueryRecord[entities.CarOrderEntity]("user_id", userId))

	queryRecords = append(queryRecords, helpers.NewQueryRecord[entities.CarOrderEntity]("car_offer_id", carOfferId))

	queryRecords = append(queryRecords, helpers.NewQueryRecord[entities.CarOrderEntity]("status", status))

	query = r.ApplyWhereConditions(query, queryRecords...)

	if startDate != "" || endDate != "" {
		query = r.ApplyDateRangeFilter(query, startDate, endDate, dateFilterType)
	}

	return r.ExecutePaginatedQuery(query, pagination, sorting)
}
