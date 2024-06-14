using Shared.UnitOfWork;

namespace RentalModule.Infrastructure.DataBase.UnitOfWork;

public interface IRentalUnitOfWork : IUnitOfWork, IDisposable, IAsyncDisposable
{
}