package repository

import (
	models "rental-api/Domain/models/domestic"
	car_image_repository "rental-api/Domain/repository_interfaces/car_image_repository"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_offer_repository"
	car_tag_repository "rental-api/Domain/repository_interfaces/car_tag_repository"
	postgres_db "rental-api/Infrastructure/databases/postgres/config"
	"rental-api/Infrastructure/databases/postgres/entities"
	mappers "rental-api/Infrastructure/databases/postgres/mappers/base"
	unit_of_work "rental-api/Infrastructure/databases/postgres/repository/base/unit_of_work"

	"github.com/google/wire"
)

func ProvideCarOfferQueryRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.CarOfferEntity, models.CarOfferModel],
	uow unit_of_work.UnitOfWork,
) repository_interfaces.CarOfferQueryRepository {
	return NewCarOfferQueryRepositoryImpl(postgresDatabase, mapper, uow)
}

func ProvideCarOfferCommandRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.CarOfferEntity, models.CarOfferModel],
	uow unit_of_work.UnitOfWork,
	carTagCommandRepository car_tag_repository.CarTagCommandRepository,
	carImageCommandRepository car_image_repository.CarImageCommandRepository,
) repository_interfaces.CarOfferCommandRepository {
	return NewCarOfferCommandRepositoryImpl(postgresDatabase, mapper, uow, carTagCommandRepository, carImageCommandRepository)
}

var WireSet = wire.NewSet(
	ProvideCarOfferQueryRepository,
	ProvideCarOfferCommandRepository,
)
