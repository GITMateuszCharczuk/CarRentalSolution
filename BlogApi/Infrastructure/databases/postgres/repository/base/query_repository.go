package cqrs

import (
	p "identity-api/Domain/pagination"
	selector "identity-api/Domain/property_selector"
	s "identity-api/Domain/sorting"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	helpers "identity-api/Infrastructure/databases/postgres/repository/base/helpers"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type QueryRepository[TEntity any, TId comparable, TModel any] struct {
	dbContext *gorm.DB
	mapper    mappers.PersistenceMapper[TEntity, TModel]
	helper    *helpers.QueryHelper[TEntity, TModel]
}

func NewQueryRepository[TEntity any, TId comparable, TModel any](
	dbContext *gorm.DB,
	mapper mappers.PersistenceMapper[TEntity, TModel],
) *QueryRepository[TEntity, TId, TModel] {
	return &QueryRepository[TEntity, TId, TModel]{
		dbContext: dbContext,
		mapper:    mapper,
		helper:    helpers.NewQueryHelper[TEntity, TModel](),
	}
}

func (r *QueryRepository[TEntity, TId, TModel]) GetById(id TId) (*TModel, error) {
	var entity TEntity
	if err := r.dbContext.First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}
	model := r.mapper.MapToModel(entity)
	return &model, nil
}

func (r *QueryRepository[TEntity, TId, TModel]) GetByIds(ids []TId) ([]TModel, error) {
	var entities []TEntity
	if err := r.dbContext.Find(&entities, "id IN ?", ids).Error; err != nil {
		return nil, err
	}
	return r.mapEntitiesToModels(entities), nil
}

func (r *QueryRepository[TEntity, TId, TModel]) GetFirstByProp(
	selector selector.PropertySelector[TEntity],
	value interface{},
) (*TModel, error) {
	if err := r.helper.ValidateProperty(selector.FieldName); err != nil {
		return nil, err
	}

	var entity TEntity
	columnName := r.helper.GetColumnName(selector.FieldName)
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

func (r *QueryRepository[TEntity, TId, TModel]) GetAllByPropValues(
	selector selector.PropertySelector[TEntity],
	values interface{},
	pagination *p.Pagination,
	sorting *s.Sortable,
) (*p.PaginatedResult[TModel], error) {
	if err := r.helper.ValidateProperty(selector.FieldName); err != nil {
		return nil, err
	}

	columnName := r.helper.GetColumnName(selector.FieldName)
	query := r.dbContext.Model(new(TEntity))

	if columnName == "roles" {
		query = query.Where(columnName+" && ?", pq.Array(values))
	} else {
		query = query.Where(columnName+" IN ?", values)
	}

	return r.executePaginatedQuery(query, pagination, sorting)
}

func (r *QueryRepository[TEntity, TId, TModel]) GetAll(
	pagination *p.Pagination,
	sorting *s.Sortable,
) (*p.PaginatedResult[TModel], error) {
	query := r.dbContext.Model(new(TEntity))
	return r.executePaginatedQuery(query, pagination, sorting)
}

func (r *QueryRepository[TEntity, TId, TModel]) executePaginatedQuery(
	query *gorm.DB,
	pagination *p.Pagination,
	sorting *s.Sortable,
) (*p.PaginatedResult[TModel], error) {
	query = r.helper.ApplySorting(query, sorting)

	var totalItems int64
	if err := query.Count(&totalItems).Error; err != nil {
		return nil, err
	}

	query = r.helper.ApplyPagination(query, pagination)

	var entities []TEntity
	if err := query.Find(&entities).Error; err != nil {
		return nil, err
	}

	models := r.mapEntitiesToModels(entities)
	totalPages := r.helper.CalculateTotalPages(totalItems, pagination)

	return &p.PaginatedResult[TModel]{
		Items:       models,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: pagination.CurrentPage,
		PageSize:    pagination.PageSize,
	}, nil
}

func (r *QueryRepository[TEntity, TId, TModel]) mapEntitiesToModels(entities []TEntity) []TModel {
	models := make([]TModel, len(entities))
	for i, entity := range entities {
		models[i] = r.mapper.MapToModel(entity)
	}
	return models
}
