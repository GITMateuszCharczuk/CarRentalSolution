using Microsoft.EntityFrameworkCore;

namespace Shared.UnitOfWork;

public class UnitOfWork<TDbContext> : IUnitOfWork, IDisposable, IAsyncDisposable
    where TDbContext : DbContext
{
    private readonly DbContext _dbContext;
    
    public UnitOfWork(TDbContext dbContext)
    {
        _dbContext = dbContext;
    }
    
    public virtual async Task CommitAsync(CancellationToken cancellationToken = default)
    {
        await _dbContext.SaveChangesAsync(cancellationToken);
    }
    
    public virtual async Task RollbackAsync(CancellationToken cancellationToken = default)
    {
        await _dbContext.DisposeAsync();
    }
    
    public virtual void Dispose()
    {
        _dbContext.Dispose();
    }

    public virtual async ValueTask DisposeAsync()
    {
        await _dbContext.DisposeAsync();
    }
}