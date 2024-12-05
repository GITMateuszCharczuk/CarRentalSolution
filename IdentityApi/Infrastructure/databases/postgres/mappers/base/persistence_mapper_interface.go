package mappers

type PersistenceMapper[TEntity any, TModel any] interface {
	MapToModel(entity TEntity) TModel
	MapToEntity(model TModel) TEntity
}
