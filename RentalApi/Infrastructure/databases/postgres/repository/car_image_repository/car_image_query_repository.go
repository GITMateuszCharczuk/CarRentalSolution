package repository

import (
	models "rental-api/Domain/models/domestic"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_image_repository"
	postgres_db "rental-api/Infrastructure/databases/postgres/config"
	"rental-api/Infrastructure/databases/postgres/entities"
	mappers "rental-api/Infrastructure/databases/postgres/mappers/base"
	base "rental-api/Infrastructure/databases/postgres/repository/base"
	"rental-api/Infrastructure/databases/postgres/repository/base/helpers"
	unit_of_work "rental-api/Infrastructure/databases/postgres/repository/base/unit_of_work"
)

type CarImageQueryRepositoryImpl struct {
	*base.QueryRepository[entities.CarImageEntity, string, models.CarOfferImageModel]
	mapper mappers.PersistenceMapper[entities.CarImageEntity, models.CarOfferImageModel]
}

func NewCarImageQueryRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.CarImageEntity, models.CarOfferImageModel],
	uow unit_of_work.UnitOfWork,
) repository_interfaces.CarImageQueryRepository {
	return &CarImageQueryRepositoryImpl{
		QueryRepository: base.NewQueryRepository[entities.CarImageEntity, string, models.CarOfferImageModel](postgresDatabase.DB, mapper, uow),
		mapper:          mapper,
	}
}

func (r *CarImageQueryRepositoryImpl) GetImageById(id string) (*models.CarOfferImageModel, error) {
	queryRecord := helpers.NewQueryRecord[entities.CarImageEntity]("id", id)
	return r.GetFirstByQueryRecord(queryRecord)
}

func (r *CarImageQueryRepositoryImpl) GetImagesByCarOfferId(carOfferId string) (*[]models.CarOfferImageModel, error) {
	queryRecord := helpers.NewQueryRecord[entities.CarImageEntity]("car_offer_id", carOfferId)
	query := r.ConstructBaseQuery()
	query = r.ApplyWhereConditions(query, queryRecord)
	result, err := r.ExecutePaginatedQuery(query, nil, nil)
	if err != nil {
		return nil, err
	}
	return &result.Items, nil
}
