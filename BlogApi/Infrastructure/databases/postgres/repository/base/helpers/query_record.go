package helpers

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
	return QueryRecord[TEntity]{Selector: selector, Value: value, TableAlias: alias}

}
