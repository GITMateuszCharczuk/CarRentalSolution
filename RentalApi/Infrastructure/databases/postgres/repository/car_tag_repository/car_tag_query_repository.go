package repository

import (
	models "rental-api/Domain/models/domestic"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_tag_repository"
	"rental-api/Domain/sorting"
	postgres_db "rental-api/Infrastructure/databases/postgres/config"
	"rental-api/Infrastructure/databases/postgres/entities"
	mappers "rental-api/Infrastructure/databases/postgres/mappers/base"
	base "rental-api/Infrastructure/databases/postgres/repository/base"
	"rental-api/Infrastructure/databases/postgres/repository/base/helpers"
	unit_of_work "rental-api/Infrastructure/databases/postgres/repository/base/unit_of_work"
)

type CarTagQueryRepositoryImpl struct {
	*base.QueryRepository[entities.CarTagEntity, string, models.CarOfferTagModel]
	mapper mappers.PersistenceMapper[entities.CarTagEntity, models.CarOfferTagModel]
}

func NewCarTagQueryRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.CarTagEntity, models.CarOfferTagModel],
	uow unit_of_work.UnitOfWork,
) repository_interfaces.CarTagQueryRepository {
	return &CarTagQueryRepositoryImpl{
		QueryRepository: base.NewQueryRepository[entities.CarTagEntity, string, models.CarOfferTagModel](postgresDatabase.DB, mapper, uow),
		mapper:          mapper,
	}
}

func (r *CarTagQueryRepositoryImpl) GetTagByID(id string) (*models.CarOfferTagModel, error) {
	queryRecord := helpers.NewQueryRecord[entities.CarTagEntity]("id", id)
	return r.GetFirstByQueryRecord(queryRecord)
}

func (r *CarTagQueryRepositoryImpl) GetTagByName(name string) (*models.CarOfferTagModel, error) {
	queryRecord := helpers.NewQueryRecord[entities.CarTagEntity]("name", name)
	return r.GetFirstByQueryRecord(queryRecord)
}

func (r *CarTagQueryRepositoryImpl) GetTagsByCarOfferId(
	carOfferId string,
	sorting sorting.Sortable,
) (*[]models.CarOfferTagModel, error) {
	db := r.GetUnitOfWork().GetTransaction().Model(&entities.CarTagEntity{})
	if carOfferId != "" {
		db = db.Joins("JOIN car_offer_tags ON car_offer_tags.tag_id = car_tag_entities.id").
			Where("car_offer_tags.car_offer_id = ?", carOfferId)
	}
	return r.ExecuteSortedQuery(db, &sorting)
}
