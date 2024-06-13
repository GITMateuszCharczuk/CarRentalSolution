using BlogModule.Infrastructure.DataBase.UnitOfWork;
using MediatR;
using Shared.CQRS.CommandHandlers;

namespace BlogModule.Infrastructure.Decorators;

public class BlogCommandHandlerBehavior<TRequest, TResponse> : IPipelineBehavior<TRequest, TResponse>
    where TRequest : ICommand<TResponse>
{
    private readonly IBlogUnitOfWork _unitOfWork;

    public BlogCommandHandlerBehavior(IBlogUnitOfWork unitOfWork)
    {
        _unitOfWork = unitOfWork;
    }

    public async Task<TResponse> Handle(TRequest request, RequestHandlerDelegate<TResponse> next, CancellationToken cancellationToken)
    {
        try
        {
            var response = await next();
            await _unitOfWork.CommitAsync(cancellationToken);
            _unitOfWork.Dispose();
            return response;
        }
        catch (Exception)
        {
            await _unitOfWork.RollbackAsync(cancellationToken);
            throw;
        }
    }
}