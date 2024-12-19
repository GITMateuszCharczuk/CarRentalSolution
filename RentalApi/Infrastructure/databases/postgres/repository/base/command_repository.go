package cqrs

import (
	"log"
	mappers "rental-api/Infrastructure/databases/postgres/mappers/base"
	base "rental-api/Infrastructure/databases/postgres/repository/base/unit_of_work"

	"context"
	"errors"

	"gorm.io/gorm"
)

type CommandRepository[TEntity any, TId comparable, TModel any] struct {
	dbContext *gorm.DB
	mapper    mappers.PersistenceMapper[TEntity, TModel]
	uow       base.UnitOfWork
}

func NewCommandRepository[TEntity any, TId comparable, TModel any](
	dbContext *gorm.DB,
	mapper mappers.PersistenceMapper[TEntity, TModel],
	uow base.UnitOfWork,
) *CommandRepository[TEntity, TId, TModel] {
	return &CommandRepository[TEntity, TId, TModel]{
		dbContext: dbContext,
		mapper:    mapper,
		uow:       uow,
	}
}

func (r *CommandRepository[TEntity, TId, TModel]) ExecuteInTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	log.Printf("IsInTransaction: %v", r.uow.IsInTransaction())
	if r.uow.IsInTransaction() {
		log.Println("Reusing existing transaction")
		return fn(r.uow.GetTransaction())
	}
	log.Println("Creating new transaction")
	return r.uow.WithTransaction(ctx, fn)
}

func (r *CommandRepository[TEntity, TId, TModel]) Add(ctx context.Context, model TModel) (TModel, error) {
	entity := r.mapper.MapToEntity(model)
	err := r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		return tx.Create(&entity).Error
	})
	if err != nil {
		return model, err
	}
	return r.mapper.MapToModel(entity), nil
}

func (r *CommandRepository[TEntity, TId, TModel]) Update(ctx context.Context, model TModel) (TModel, error) {
	entity := r.mapper.MapToEntity(model)
	err := r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		return tx.Save(&entity).Error
	})
	if err != nil {
		return model, err
	}
	return r.mapper.MapToModel(entity), nil
}

func (r *CommandRepository[TEntity, TId, TModel]) Delete(ctx context.Context, id TId) error {
	return r.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		var entity TEntity
		if err := tx.First(&entity, "id = ?", id).Error; err != nil {
			return errors.New("cannot delete entity that does not exist")
		}
		return tx.Delete(&entity).Error
	})
}

func (r *CommandRepository[TEntity, TId, TModel]) GetUnitOfWork() base.UnitOfWork {
	return r.uow
}
