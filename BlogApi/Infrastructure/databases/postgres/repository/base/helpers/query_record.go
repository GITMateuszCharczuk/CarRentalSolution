package helpers

import (
	"fmt"
	"reflect"
)

type QueryRecord[TEntity any] struct {
	Selector PropertySelector[TEntity]
	Value    any
}

func NewQueryRecord[TEntity any](fieldName string, value any) QueryRecord[TEntity] {
	selector := NewPropertySelector[TEntity](fieldName)

	entityType := reflect.TypeOf((*TEntity)(nil)).Elem()

	if reflect.TypeOf(value) == entityType {
		return QueryRecord[TEntity]{Selector: selector, Value: value}
	}

	if reflect.TypeOf(value).Kind() == reflect.Slice && reflect.TypeOf(value).Elem() == entityType {
		return QueryRecord[TEntity]{Selector: selector, Value: value}
	}

	panic(fmt.Errorf("value must be of type %s or []%s", entityType, entityType))
}
