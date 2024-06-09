using BlogModule.Infrastructure.DataBase.Context;
using Shared.UnitOfWork;

namespace BlogModule.Infrastructure.DataBase.UnitOfWork;

public class BlogUnitOfWork : UnitOfWork<BlogDbContext>, IBlogUnitOfWork
{
    public BlogUnitOfWork(BlogDbContext dbContext) : base(dbContext)
    {
    }
    
    public override async Task CommitAsync(CancellationToken cancellationToken = default) =>
        await base.CommitAsync(cancellationToken);
    
    public override async Task RollbackAsync(CancellationToken cancellationToken = default) =>
        await base.RollbackAsync(cancellationToken);
    
}