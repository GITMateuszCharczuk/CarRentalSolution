package repository

import (
	"context"
	models "rental-api/Domain/models/domestic"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_order_repository"
	postgres_db "rental-api/Infrastructure/databases/postgres/config"
	"rental-api/Infrastructure/databases/postgres/entities"
	mappers "rental-api/Infrastructure/databases/postgres/mappers/base"
	base "rental-api/Infrastructure/databases/postgres/repository/base"
	unit_of_work "rental-api/Infrastructure/databases/postgres/repository/base/unit_of_work"

	"gorm.io/gorm"
)

type CarOrderCommandRepositoryImpl struct {
	*base.CommandRepository[entities.CarOrderEntity, string, models.CarOrderModel]
}

func NewCarOrderCommandRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.CarOrderEntity, models.CarOrderModel],
	uow unit_of_work.UnitOfWork,
) repository_interfaces.CarOrderCommandRepository {
	return &CarOrderCommandRepositoryImpl{
		CommandRepository: base.NewCommandRepository[entities.CarOrderEntity, string, models.CarOrderModel](postgresDatabase.DB, mapper, uow),
	}
}

func (r *CarOrderCommandRepositoryImpl) CreateCarOrder(ctx context.Context, carOrder *models.CarOrderModel) (*string, error) {
	var result string
	err := r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		createdOrder, err := r.Add(ctx, *carOrder)
		if err != nil {
			return err
		}
		result = createdOrder.Id
		return nil
	})
	return &result, err
}

func (r *CarOrderCommandRepositoryImpl) UpdateCarOrder(ctx context.Context, carOrder *models.CarOrderModel) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		_, err := r.Update(ctx, *carOrder)
		return err
	})
}

func (r *CarOrderCommandRepositoryImpl) DeleteCarOrder(ctx context.Context, id string) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		return r.Delete(ctx, id) //TODO
	})
}
