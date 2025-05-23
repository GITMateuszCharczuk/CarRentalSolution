package helpers

import (
	p "blog-api/Domain/pagination"
	s "blog-api/Domain/sorting"
	"log"
	"math"
	"reflect"

	"gorm.io/gorm"
)

type QueryHelper[TEntity any, TModel any] struct {
	entityType reflect.Type
}

func NewQueryHelper[TEntity any, TModel any]() *QueryHelper[TEntity, TModel] {
	return &QueryHelper[TEntity, TModel]{
		entityType: reflect.TypeOf(new(TEntity)).Elem(),
	}
}

func (h *QueryHelper[TEntity, TModel]) ApplySorting(query *gorm.DB, sorting *s.Sortable) *gorm.DB {
	if !sorting.Enabled || len(sorting.SortFields) == 0 {
		return query
	}

	for _, sort := range sorting.SortFields {
		columnName := sort.Field
		if columnName == "" {
			continue
		}

		if sort.Direction == s.DESC {
			query = query.Order(columnName + " DESC")
			log.Println("query", query)
		} else {
			query = query.Order(columnName + " ASC")
			log.Println("query", query)
		}
	}

	return query
}

func (h *QueryHelper[TEntity, TModel]) ApplyPagination(query *gorm.DB, pagination *p.Pagination) *gorm.DB {
	if pagination.Enabled {
		offset := (pagination.CurrentPage - 1) * pagination.PageSize
		query = query.Offset(offset).Limit(pagination.PageSize)
	}
	return query
}

func (h *QueryHelper[TEntity, TModel]) CalculateTotalPages(totalItems int64, pagination *p.Pagination) int {
	if !pagination.Enabled || pagination.PageSize <= 0 {
		return 1
	}
	return int(math.Ceil(float64(totalItems) / float64(pagination.PageSize)))
}
