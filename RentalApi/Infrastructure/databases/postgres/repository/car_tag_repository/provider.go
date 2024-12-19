package repository

import (
	models "rental-api/Domain/models/domestic"
	car_offer_repository "rental-api/Domain/repository_interfaces/car_offer_repository"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_tag_repository"
	postgres_db "rental-api/Infrastructure/databases/postgres/config"
	"rental-api/Infrastructure/databases/postgres/entities"
	mappers "rental-api/Infrastructure/databases/postgres/mappers/base"
	unit_of_work "rental-api/Infrastructure/databases/postgres/repository/base/unit_of_work"

	"github.com/google/wire"
)

func ProvideCarTagQueryRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.CarTagEntity, models.CarOfferTagModel],
	uow unit_of_work.UnitOfWork,
) repository_interfaces.CarTagQueryRepository {
	return NewCarTagQueryRepositoryImpl(postgresDatabase, mapper, uow)
}

func ProvideCarTagCommandRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.CarTagEntity, models.CarOfferTagModel],
	uow unit_of_work.UnitOfWork,
	carOfferQueryRepository car_offer_repository.CarOfferQueryRepository,
) repository_interfaces.CarTagCommandRepository {
	return NewCarTagCommandRepositoryImpl(postgresDatabase, mapper, uow, carOfferQueryRepository)
}

var WireSet = wire.NewSet(
	ProvideCarTagQueryRepository,
	ProvideCarTagCommandRepository,
)
