package repository

import (
	"fmt"
	"rental-api/Domain/constants"
	models "rental-api/Domain/models/domestic"
	"rental-api/Domain/pagination"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_offer_repository"
	"rental-api/Domain/sorting"
	postgres_db "rental-api/Infrastructure/databases/postgres/config"
	"rental-api/Infrastructure/databases/postgres/entities"
	mappers "rental-api/Infrastructure/databases/postgres/mappers/base"
	base "rental-api/Infrastructure/databases/postgres/repository/base"
	"rental-api/Infrastructure/databases/postgres/repository/base/helpers"
	unit_of_work "rental-api/Infrastructure/databases/postgres/repository/base/unit_of_work"
	"time"
)

type CarOfferQueryRepositoryImpl struct {
	*base.QueryRepository[entities.CarOfferEntity, string, models.CarOfferModel]
	mapper mappers.PersistenceMapper[entities.CarOfferEntity, models.CarOfferModel]
}

func NewCarOfferQueryRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.CarOfferEntity, models.CarOfferModel],
	uow unit_of_work.UnitOfWork,
) repository_interfaces.CarOfferQueryRepository {
	return &CarOfferQueryRepositoryImpl{
		QueryRepository: base.NewQueryRepository[entities.CarOfferEntity, string, models.CarOfferModel](postgresDatabase.DB, mapper, uow),
		mapper:          mapper,
	}
}

func (r *CarOfferQueryRepositoryImpl) GetCarOfferByID(id string) (*models.CarOfferModel, error) {
	return r.GetById(id)
}

func (r *CarOfferQueryRepositoryImpl) GetCarOfferCustodianIdByID(id string) (string, error) {
	query := r.ConstructBaseQuery()
	query = query.Where("id = ?", id)
	query = query.Select("custodian_id")

	var custodianId string
	err := query.First(&custodianId).Error
	if err != nil {
		return "", err
	}
	return custodianId, nil
}

func (r *CarOfferQueryRepositoryImpl) GetCarOffers(
	pagination *pagination.Pagination,
	sorting *sorting.Sortable,
	ids []string,
	dateTimeFrom string,
	dateTimeTo string,
	tags []string,
	visible string,
) (*pagination.PaginatedResult[models.CarOfferModel], error) {
	query := r.ConstructBaseQuery()

	if len(tags) > 0 {
		query = query.Joins("JOIN car_offer_tags ON car_offer_tags.car_offer_id = car_offer_entities.id").
			Joins("JOIN car_tag_entities ON car_tag_entities.id = car_offer_tags.tag_id").
			Where("car_tag_entities.name IN ?", tags)
	}

	queryRecords := []helpers.QueryRecord[entities.CarOfferEntity]{
		helpers.NewQueryRecord[entities.CarOfferEntity]("id", ids, "car_offer_entities"),
	}

	if visible != "" {
		queryRecords = append(queryRecords, helpers.NewQueryRecord[entities.CarOfferEntity]("visible", visible))
	}

	query = r.ApplyWhereConditions(query, queryRecords...)

	if dateTimeFrom != "" && dateTimeTo != "" {
		fromTime, err := time.Parse(time.RFC3339, dateTimeFrom)
		if err != nil {
			return nil, fmt.Errorf("invalid dateTimeFrom format: %v", err)
		}

		toTime, err := time.Parse(time.RFC3339, dateTimeTo)
		if err != nil {
			return nil, fmt.Errorf("invalid dateTimeTo format: %v", err)
		}

		fromStr := fromTime.Format("2006-01-02 15:04:05.999999-07")
		toStr := toTime.Format("2006-01-02 15:04:05.999999-07")

		query = query.Where("NOT EXISTS (SELECT 1 FROM car_order_entities co "+
			"WHERE co.car_offer_id = car_offer_entities.id "+
			"AND co.status NOT IN (?, ?) "+
			"AND ((co.start_date <= ? AND co.end_date >= ?) OR "+ // Case 1: Order spans the entire range
			"(co.start_date <= ? AND co.end_date >= ?) OR "+ // Case 2: Order starts before and ends during
			"(co.start_date >= ? AND co.end_date <= ?)))", // Case 3: Order is completely within
			constants.OrderStatusCancelled,
			constants.OrderStatusArchived,
			toStr, fromStr, // Case 1
			fromStr, fromStr, // Case 2
			fromStr, toStr, // Case 3
		)
	}

	return r.ExecutePaginatedQuery(query, pagination, sorting)
}
