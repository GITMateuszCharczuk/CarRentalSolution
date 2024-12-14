package cqrs

import (
	p "identity-api/Domain/pagination"
	s "identity-api/Domain/sorting"
	mappers "identity-api/Infrastructure/databases/postgres/mappers/base"
	helpers "identity-api/Infrastructure/databases/postgres/repository/base/helpers"
	"reflect"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type QueryRepository[TEntity any, TId comparable, TModel any] struct {
	DbContext *gorm.DB
	mapper    mappers.PersistenceMapper[TEntity, TModel]
	helper    *helpers.QueryHelper[TEntity, TModel]
}

func NewQueryRepository[TEntity any, TId comparable, TModel any](
	dbContext *gorm.DB,
	mapper mappers.PersistenceMapper[TEntity, TModel],
) *QueryRepository[TEntity, TId, TModel] {
	return &QueryRepository[TEntity, TId, TModel]{
		DbContext: dbContext,
		mapper:    mapper,
		helper:    helpers.NewQueryHelper[TEntity, TModel](),
	}
}

func (r *QueryRepository[TEntity, TId, TModel]) ConstructBaseQuery() *gorm.DB {
	query := r.DbContext.Model(new(TEntity))
	return query
}

func (r *QueryRepository[TEntity, TId, TModel]) GetFirstByQueryRecord(
	queryRecord helpers.QueryRecord[TEntity],
) (*TModel, error) {
	var entity TEntity

	if err := r.DbContext.Where(queryRecord.Selector.FieldName+" = ?", queryRecord.Value).First(&entity).Error; err != nil {
		return nil, err
	}
	model := r.mapper.MapToModel(entity)
	return &model, nil
}

func (r *QueryRepository[TEntity, TId, TModel]) GetTotalCount(query ...*gorm.DB) (int64, error) {
	var count int64
	if len(query) == 0 || query[0] == nil {
		query = append(query, r.DbContext.Model(new(TEntity)))
	}
	if err := query[0].Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *QueryRepository[TEntity, TId, TModel]) GetAllByQueryRecords(
	pagination *p.Pagination,
	sorting *s.Sortable,
	queryRecords ...helpers.QueryRecord[TEntity],
) (*p.PaginatedResult[TModel], error) {
	query := r.ConstructBaseQuery()
	query = r.ApplyWhereConditions(query, queryRecords...)
	return r.ExecutePaginatedQuery(query, pagination, sorting)
}

func (r *QueryRepository[TEntity, TId, TModel]) GetAllSortedAndPaginated(
	pagination *p.Pagination,
	sorting *s.Sortable,
) (*p.PaginatedResult[TModel], error) {
	query := r.DbContext.Model(new(TEntity))
	return r.ExecutePaginatedQuery(query, pagination, sorting)
}

// TODO change name of method
func (r *QueryRepository[TEntity, TId, TModel]) ApplyWhereConditions(query *gorm.DB, queryRecords ...helpers.QueryRecord[TEntity]) *gorm.DB {
	tableName := query.Statement.Table

	for _, record := range queryRecords {
		prefix := tableName
		if record.TableAlias != "" {
			prefix = record.TableAlias
		}

		fieldName := prefix + "." + record.Selector.FieldName

		if reflect.TypeOf(record.Value).Kind() == reflect.Slice {
			if len(record.Value.([]string)) == 0 {
				continue
			}
			query = query.Where(fieldName+" && ?", pq.Array(record.Value.([]string)))
		} else {
			if record.Value.(string) == "" {
				continue
			}
			query = query.Where(fieldName+" = ?", record.Value)
		}
	}
	return query
}

func (r *QueryRepository[TEntity, TId, TModel]) ExecutePaginatedQuery(
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

func (r *QueryRepository[TEntity, TId, TModel]) ExecuteSortedQuery(
	query *gorm.DB,
	sorting *s.Sortable,
) (*[]TModel, error) {
	query = r.helper.ApplySorting(query, sorting)

	var entities []TEntity
	if err := query.Find(&entities).Error; err != nil {
		return nil, err
	}

	models := r.mapEntitiesToModels(entities)

	return &models, nil
}

func (r *QueryRepository[TEntity, TId, TModel]) mapEntitiesToModels(entities []TEntity) []TModel {
	models := make([]TModel, len(entities))
	for i, entity := range entities {
		models[i] = r.mapper.MapToModel(entity)
	}
	return models
}
