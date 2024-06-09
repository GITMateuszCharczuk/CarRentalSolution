using Shared.UnitOfWork;

namespace BlogModule.Infrastructure.DataBase.UnitOfWork;

public interface IBlogUnitOfWork : IUnitOfWork, IDisposable, IAsyncDisposable
{
}