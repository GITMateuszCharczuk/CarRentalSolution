package cqrs

import (
	"log"
	"reflect"
	p "rental-api/Domain/pagination"
	s "rental-api/Domain/sorting"
	mappers "rental-api/Infrastructure/databases/postgres/mappers/base"
	helpers "rental-api/Infrastructure/databases/postgres/repository/base/helpers"
	unit_of_work "rental-api/Infrastructure/databases/postgres/repository/base/unit_of_work"
	"strings"
	"time"

	"gorm.io/gorm"
)

type QueryRepository[TEntity any, TId comparable, TModel any] struct {
	DbContext *gorm.DB
	mapper    mappers.PersistenceMapper[TEntity, TModel]
	helper    *helpers.QueryHelper[TEntity, TModel]
	uow       unit_of_work.UnitOfWork
}

func NewQueryRepository[TEntity any, TId comparable, TModel any](
	dbContext *gorm.DB,
	mapper mappers.PersistenceMapper[TEntity, TModel],
	uow unit_of_work.UnitOfWork,
) *QueryRepository[TEntity, TId, TModel] {
	return &QueryRepository[TEntity, TId, TModel]{
		DbContext: dbContext,
		mapper:    mapper,
		helper:    helpers.NewQueryHelper[TEntity, TModel](),
		uow:       uow,
	}
}

func (r *QueryRepository[TEntity, TId, TModel]) ConstructBaseQuery() *gorm.DB {
	db := r.uow.GetTransaction()
	query := db.Model(new(TEntity))
	return query
}

func (r *QueryRepository[TEntity, TId, TModel]) GetById(id TId) (*TModel, error) {
	var entity TEntity
	query := r.ConstructBaseQuery()
	query = query.Where("id = ?", id)
	if err := query.First(&entity).Error; err != nil {
		return nil, err
	}
	model := r.mapper.MapToModel(entity)
	return &model, nil
}

func (r *QueryRepository[TEntity, TId, TModel]) GetFirstByQueryRecord(
	queryRecord helpers.QueryRecord[TEntity],
) (*TModel, error) {
	var entity TEntity
	query := r.ConstructBaseQuery()
	query = query.Where(queryRecord.Selector.FieldName+" = ?", queryRecord.Value)
	if err := query.First(&entity).Error; err != nil {
		return nil, err
	}
	model := r.mapper.MapToModel(entity)
	return &model, nil
}

func (r *QueryRepository[TEntity, TId, TModel]) GetTotalCount(query ...*gorm.DB) (int64, error) {
	var count int64
	db := r.uow.GetTransaction()
	if len(query) == 0 || query[0] == nil {
		query = append(query, db.Model(new(TEntity)))
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
	db := r.uow.GetTransaction()
	query := db.Model(new(TEntity))
	return r.ExecutePaginatedQuery(query, pagination, sorting)
}

func (r *QueryRepository[TEntity, TId, TModel]) ApplyWhereConditions(query *gorm.DB, queryRecords ...helpers.QueryRecord[TEntity]) *gorm.DB {
	for _, record := range queryRecords {
		var columnField string
		if record.TableAlias != "" {
			columnField = record.TableAlias + "." + record.Selector.FieldName
		} else {
			columnField = record.Selector.FieldName
		}
		valueType := reflect.TypeOf(record.Value).Kind()

		if valueType == reflect.Slice {
			sliceValue := reflect.ValueOf(record.Value)
			if sliceValue.Len() == 0 {
				continue
			}
			values := make([]interface{}, 0, sliceValue.Len())
			for i := 0; i < sliceValue.Len(); i++ {
				value := strings.TrimSpace(sliceValue.Index(i).Interface().(string))
				if value != "" {
					values = append(values, value)
				}
			}
			query = query.Where(columnField+" IN ?", values)
		} else {
			if record.Value.(string) == "" {
				continue
			}
			query = query.Where(columnField+" = ?", record.Value)
		}
	}
	return query
}

func (r *QueryRepository[TEntity, TId, TModel]) ApplyDateRangeFilter(query *gorm.DB, startDate, endDate, filterType string) *gorm.DB {
	var start, end time.Time
	var err error

	if startDate != "" {
		start, err = time.Parse(time.RFC3339, startDate)
		if err != nil {
			log.Printf("Invalid start date format: %v", err)
			return query
		}
	}

	if endDate != "" {
		end, err = time.Parse(time.RFC3339, endDate)
		if err != nil {
			log.Printf("Invalid end date format: %v", err)
			return query
		}
	}

	switch filterType {
	case "BETWEEN":
		if !start.IsZero() && !end.IsZero() {
			return query.Where("(start_date BETWEEN ? AND ? OR end_date BETWEEN ? AND ?)",
				start, end, start, end)
		}
	case "START_BETWEEN":
		if !start.IsZero() && !end.IsZero() {
			return query.Where("start_date >= ? AND start_date < ?", start, end)
		} else if !end.IsZero() {
			return query.Where("start_date >= ?", start)
		}
	case "END_BETWEEN":
		if !start.IsZero() && !end.IsZero() {
			return query.Where("end_date >= ? AND end_date < ?", start, end)
		} else if !start.IsZero() {
			return query.Where("end_date >= ?", start)
		}
	default:
		if !start.IsZero() && !end.IsZero() {
			return query.Where("(start_date BETWEEN ? AND ? OR end_date BETWEEN ? AND ?)",
				start, end, start, end)
		} else if !start.IsZero() {
			return query.Where("start_date >= ?", start)
		} else if !end.IsZero() {
			return query.Where("end_date <= ?", end)
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

func (r *QueryRepository[TEntity, TId, TModel]) GetUnitOfWork() unit_of_work.UnitOfWork {
	return r.uow
}
