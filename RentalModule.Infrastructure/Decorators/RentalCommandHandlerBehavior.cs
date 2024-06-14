using MediatR;
using RentalModule.Infrastructure.DataBase.UnitOfWork;
using Shared.CQRS.CommandHandlers;

namespace RentalModule.Infrastructure.Decorators;

public class RentalCommandHandlerBehavior<TRequest, TResponse> : IPipelineBehavior<TRequest, TResponse>
    where TRequest : ICommand<TResponse>
{
    private readonly IRentalUnitOfWork _unitOfWork;

    public RentalCommandHandlerBehavior(IRentalUnitOfWork unitOfWork)
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