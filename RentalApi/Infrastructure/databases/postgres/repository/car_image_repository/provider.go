package repository

import (
	models "rental-api/Domain/models/domestic"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_image_repository"
	postgres_db "rental-api/Infrastructure/databases/postgres/config"
	"rental-api/Infrastructure/databases/postgres/entities"
	mappers "rental-api/Infrastructure/databases/postgres/mappers/base"
	unit_of_work "rental-api/Infrastructure/databases/postgres/repository/base/unit_of_work"

	"github.com/google/wire"
)

func ProvideCarImageQueryRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.CarImageEntity, models.CarOfferImageModel],
	uow unit_of_work.UnitOfWork,
) repository_interfaces.CarImageQueryRepository {
	return NewCarImageQueryRepositoryImpl(postgresDatabase, mapper, uow)
}

func ProvideCarImageCommandRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.CarImageEntity, models.CarOfferImageModel],
	uow unit_of_work.UnitOfWork,
) repository_interfaces.CarImageCommandRepository {
	return NewCarImageCommandRepositoryImpl(postgresDatabase, mapper, uow)
}

var WireSet = wire.NewSet(
	ProvideCarImageQueryRepository,
	ProvideCarImageCommandRepository,
)
