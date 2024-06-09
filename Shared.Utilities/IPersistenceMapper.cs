namespace Shared.Utilities;

public interface IPersistenceMapper<TEntity,TModel>
{
    public TModel MapToModel(TEntity entity);
    public TEntity MapToEntity(TModel model);
}