package cqrs

import (
	"errors"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	"log"

	"gorm.io/gorm"
)

type CommandRepository[TEntity any, TId comparable, TModel any] struct {
	DbContext *gorm.DB
	mapper    mappers.PersistenceMapper[TEntity, TModel]
}

func NewCommandRepository[TEntity any, TId comparable, TModel any](dbContext *gorm.DB, mapper mappers.PersistenceMapper[TEntity, TModel]) *CommandRepository[TEntity, TId, TModel] {
	return &CommandRepository[TEntity, TId, TModel]{DbContext: dbContext, mapper: mapper}
}

func (r *CommandRepository[TEntity, TId, TModel]) Add(model TModel) (TModel, error) {
	entity := r.mapper.MapToEntity(model)
	result := r.DbContext.Create(&entity)
	if result.Error != nil {
		log.Println(entity)
		return model, result.Error
	}
	return r.mapper.MapToModel(entity), nil
}

func (r *CommandRepository[TEntity, TId, TModel]) Update(model TModel) (TModel, error) {
	entity := r.mapper.MapToEntity(model)
	result := r.DbContext.Save(&entity)
	if result.Error != nil {
		return model, result.Error
	}
	return r.mapper.MapToModel(entity), nil
}

func (r *CommandRepository[TEntity, TId, TModel]) Delete(id TId) error {
	var entity TEntity
	if err := r.DbContext.First(&entity, "id = ?", id).Error; err != nil {
		return errors.New("cannot delete entity that does not exist")
	}
	return r.DbContext.Delete(&entity).Error
}
