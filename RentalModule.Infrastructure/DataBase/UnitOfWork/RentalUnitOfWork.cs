using RentalModule.Infrastructure.DataBase.Context;
using Shared.UnitOfWork;

namespace RentalModule.Infrastructure.DataBase.UnitOfWork;

public class RentalUnitOfWork : UnitOfWork<RentalDbContext>, IRentalUnitOfWork
{
    public RentalUnitOfWork(RentalDbContext dbContext) : base(dbContext)
    {
    }
    
    public override async Task CommitAsync(CancellationToken cancellationToken = default) =>
        await base.CommitAsync(cancellationToken);
    
    public override async Task RollbackAsync(CancellationToken cancellationToken = default) =>
        await base.RollbackAsync(cancellationToken);
    
}