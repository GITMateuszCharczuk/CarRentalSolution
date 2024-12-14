package helpers

import (
	"fmt"
	"reflect"
)

type QueryRecord[TEntity any] struct {
	Selector   PropertySelector[TEntity]
	Value      any
	TableAlias string
}

func NewQueryRecord[TEntity any](fieldName string, value any, tableAlias ...string) QueryRecord[TEntity] {
	var alias string
	if len(tableAlias) > 0 {
		alias = tableAlias[0]
	}

	selector := NewPropertySelector[TEntity](fieldName)

	entityType := reflect.TypeOf((*TEntity)(nil)).Elem()

	if reflect.TypeOf(value) == entityType {
		return QueryRecord[TEntity]{Selector: selector, Value: value, TableAlias: alias}
	}

	if reflect.TypeOf(value).Kind() == reflect.Slice && reflect.TypeOf(value).Elem() == entityType {
		return QueryRecord[TEntity]{Selector: selector, Value: value, TableAlias: alias}
	}

	panic(fmt.Errorf("value must be of type %s or []%s", entityType, entityType))
}
