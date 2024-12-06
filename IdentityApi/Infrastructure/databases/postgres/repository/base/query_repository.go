package cqrs

import (
	"errors"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	"reflect"
	"strings"

	selector "identity-api/Domain/property_selector"

	"gorm.io/gorm"
)

type QueryRepository[TEntity any, TId comparable, TModel any] struct {
	dbContext *gorm.DB
	mapper    mappers.PersistenceMapper[TEntity, TModel]
}

func NewQueryRepository[TEntity any, TId comparable, TModel any](dbContext *gorm.DB, mapper mappers.PersistenceMapper[TEntity, TModel]) *QueryRepository[TEntity, TId, TModel] {
	return &QueryRepository[TEntity, TId, TModel]{dbContext: dbContext, mapper: mapper}
}

func (r *QueryRepository[TEntity, TId, TModel]) GetById(id TId) (*TModel, error) {
	var entity TEntity
	if err := r.dbContext.First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}
	model := r.mapper.MapToModel(entity)
	return &model, nil
}

func (r *QueryRepository[TEntity, TId, TModel]) GetByIds(ids []TId) ([]*TModel, error) {
	var entities []TEntity
	if err := r.dbContext.Find(&entities, "id IN ?", ids).Error; err != nil {
		return nil, err
	}
	models := make([]*TModel, len(entities))
	for i, entity := range entities {
		model := r.mapper.MapToModel(entity)
		models[i] = &model
	}
	return models, nil
}

func (r *QueryRepository[TEntity, TId, TModel]) GetFirstByProp(selector selector.PropertySelector[TEntity], value interface{}) (*TModel, error) {
	var entity TEntity
	columnName := r.getColumnName(selector.FieldName)
	if columnName == "" {
		return nil, errors.New("property " + selector.FieldName + " does not exist on type " + reflect.TypeOf(entity).Name())
	}

	if err := r.dbContext.Where(columnName+" = ?", value).First(&entity).Error; err != nil {
		return nil, err
	}
	model := r.mapper.MapToModel(entity)
	return &model, nil
}

func (r *QueryRepository[TEntity, TId, TModel]) GetTotalCount() (int64, error) {
	var count int64
	if err := r.dbContext.Model(new(TEntity)).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *QueryRepository[TEntity, TId, TModel]) GetAll() ([]*TModel, error) {
	var entities []TEntity
	if err := r.dbContext.Find(&entities).Error; err != nil {
		return nil, err
	}

	models := make([]*TModel, len(entities))
	for i, entity := range entities {
		model := r.mapper.MapToModel(entity)
		models[i] = &model
	}
	return models, nil
}

func (r *QueryRepository[TEntity, TId, TModel]) GetAllByPropValues(selector selector.PropertySelector[TEntity], values ...interface{}) ([]*TModel, error) {
	var entity TEntity

	columnName := r.getColumnName(selector.FieldName)
	if columnName == "" {
		return nil, errors.New("property " + selector.FieldName + " does not exist on type " + reflect.TypeOf(entity).Name())
	}

	var entities []TEntity
	if err := r.dbContext.Where(columnName+" IN ?", values).Find(&entities).Error; err != nil {
		return nil, err
	}

	models := make([]*TModel, len(entities))
	for i, entity := range entities {
		model := r.mapper.MapToModel(entity)
		models[i] = &model
	}
	return models, nil
}

func (r *QueryRepository[TEntity, TId, TModel]) getColumnName(propName string) string {
	entityType := reflect.TypeOf(new(TEntity)).Elem()

	for i := 0; i < entityType.NumField(); i++ {
		field := entityType.Field(i)

		jsonTag := field.Tag.Get("json")
		if jsonTag == propName {
			if gormTag := field.Tag.Get("gorm"); gormTag != "" {
				for _, tag := range strings.Split(gormTag, ";") {
					if strings.HasPrefix(tag, "column:") {
						return strings.TrimPrefix(tag, "column:")
					}
				}
			}
			return jsonTag
		}
	}
	return ""
}
