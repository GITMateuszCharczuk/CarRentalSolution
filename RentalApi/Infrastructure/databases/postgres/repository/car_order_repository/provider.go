package repository

import (
	models "rental-api/Domain/models/domestic"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_order_repository"
	postgres_db "rental-api/Infrastructure/databases/postgres/config"
	"rental-api/Infrastructure/databases/postgres/entities"
	mappers "rental-api/Infrastructure/databases/postgres/mappers/base"
	unit_of_work "rental-api/Infrastructure/databases/postgres/repository/base/unit_of_work"

	"github.com/google/wire"
)

func ProvideCarOrderQueryRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.CarOrderEntity, models.CarOrderModel],
	uow unit_of_work.UnitOfWork,
) repository_interfaces.CarOrderQueryRepository {
	return NewCarOrderQueryRepositoryImpl(postgresDatabase, mapper, uow)
}

func ProvideCarOrderCommandRepository(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.CarOrderEntity, models.CarOrderModel],
	uow unit_of_work.UnitOfWork,
) repository_interfaces.CarOrderCommandRepository {
	return NewCarOrderCommandRepositoryImpl(postgresDatabase, mapper, uow)
}

var WireSet = wire.NewSet(
	ProvideCarOrderQueryRepository,
	ProvideCarOrderCommandRepository,
)
