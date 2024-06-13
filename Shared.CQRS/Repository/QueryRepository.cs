using System.Collections.Immutable;
using Microsoft.EntityFrameworkCore;
using Shared.Utilities;

namespace Shared.CQRS.Repository;

public class QueryRepository<TEntity, TId, TModel, TDbContext>
    where TEntity : class, IEntity<TId>
    where TId : struct//BaseId<TId>
    where TModel : class
    where TDbContext : DbContext
{
    protected readonly TDbContext DbContext;
    protected readonly IPersistenceMapper<TEntity, TModel> Mapper;

    public QueryRepository(
        TDbContext dbContext, 
        IPersistenceMapper<TEntity, TModel> mapper)
    {
        DbContext = dbContext;
        Mapper = mapper;
    }

    public virtual async Task<TModel?> GetByIdAsync(TId id, CancellationToken cancellationToken = default) =>
        await DbContext.Set<TEntity>()
            .AsNoTracking()
            .FirstOrDefaultAsync(x => x.Id.Equals(id), cancellationToken)
            .ContinueWith(x => x.Result is null ? null : Mapper.MapToModel(x.Result), cancellationToken);

    public virtual async Task<ImmutableArray<TModel>> GetByIdsAsync(ImmutableArray<TId> ids, CancellationToken cancellationToken = default) =>
        await DbContext.Set<TEntity>()
            .AsNoTracking()
            .Where(x => ids.Contains(x.Id))
            .Select(x => Mapper.MapToModel(x))
            .ToImmutableArrayAsync(cancellationToken);
    
    public virtual async Task<int> GetTotalCountAsync(CancellationToken cancellationToken = default) =>
        await DbContext.Set<TEntity>()
            .AsNoTracking()
            .CountAsync(cancellationToken);
}