package property_selector

type PropertySelector[TEntity any] struct {
	FieldName string
}

func NewPropertySelector[TEntity any](fieldName string) PropertySelector[TEntity] {
	return PropertySelector[TEntity]{FieldName: fieldName}
}
