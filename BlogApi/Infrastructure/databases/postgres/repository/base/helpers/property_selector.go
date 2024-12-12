package helpers

import (
	"fmt"
	"reflect"
)

type PropertySelector[TEntity any] struct {
	FieldName string
}

func NewPropertySelector[TEntity any](fieldName string) PropertySelector[TEntity] {
	entityType := reflect.TypeOf((*TEntity)(nil)).Elem()

	if _, found := entityType.FieldByName(fieldName); !found {
		panic(fmt.Errorf("field %s does not exist on %s", fieldName, entityType.Name()))
	}

	return PropertySelector[TEntity]{FieldName: fieldName}
}
