using Microsoft.EntityFrameworkCore;
using Shared.Utilities;

namespace Shared.CQRS.Repository;

public abstract class CommandRepository<TEntity, TId, TModel, TDbContext>
    where TEntity : class, IEntity<TId>
    where TId : struct //BaseId<TId>
    where TModel : class
    where TDbContext : DbContext
{
    protected readonly TDbContext DbContext;
    protected readonly IPersistenceMapper<TEntity, TModel> Mapper;

    public CommandRepository(
        TDbContext dbContext, 
        IPersistenceMapper<TEntity, TModel> mapper)
    {
        DbContext = dbContext;
        Mapper = mapper;
    }
    
    public virtual async Task<TModel> AddAsync(TModel model, CancellationToken cancellationToken = default)
    {
        var entity = Mapper.MapToEntity(model);
        return Mapper.MapToModel((await DbContext.Set<TEntity>().AddAsync(entity, cancellationToken)).Entity);
    }

    public virtual async Task<TModel> UpdateAsync(TModel model, CancellationToken cancellationToken = default)
    {
        var entity = Mapper.MapToEntity(model);
        return await Task.FromResult(Mapper.MapToModel(DbContext.Set<TEntity>().Update(entity).Entity));
    }

    public virtual async Task DeleteAsync(TId id, CancellationToken cancellationToken = default)
    {
        var entity = await DbContext.Set<TEntity>().FirstOrDefaultAsync(entity => entity.Id.Equals(id), cancellationToken);
        if (entity is null) throw new InvalidOperationException("Cannot delete entity that does not exist");
        await Task.FromResult(DbContext.Set<TEntity>().Remove(entity));
    }
}