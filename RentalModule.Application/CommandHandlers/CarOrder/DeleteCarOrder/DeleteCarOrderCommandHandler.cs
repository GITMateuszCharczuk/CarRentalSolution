using RentalModule.Application.Contract.CarOrders.DeleteCarOrder;
using RentalModule.Domain.Models.Ids;
using RentalModule.Domain.RepositoryInterfaces.CarOrder;
using Results.Application;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.CommandHandlers;

namespace RentalModule.Application.CommandHandlers.CarOrder.DeleteCarOrder;

public class DeleteCarOrderCommandHandler : ICommandHandler<DeleteCarOrderCommand, HandlerResult<DeleteCarOrderResponse, IErrorResult>>
{
    private readonly ICarOrderCommandRepository _commandRepository;
    private readonly ICarOrderQueryRepository _queryRepository;

    public DeleteCarOrderCommandHandler(
        ICarOrderCommandRepository commandRepository, 
        ICarOrderQueryRepository queryRepository)
    {
        _commandRepository = commandRepository;
        _queryRepository = queryRepository;
    }

    public async Task<HandlerResult<DeleteCarOrderResponse, IErrorResult>> Handle(DeleteCarOrderCommand request, CancellationToken cancellationToken)
    {
        var isSuccess = Guid.TryParse(request.Id, out var carOrderId);
        var carOrder = isSuccess ? await _queryRepository.GetByIdAsync(new CarOrderId(carOrderId), cancellationToken) : null;
            
        if (carOrder is null) return new EntityNotFoundErrorResult() {
            Title = "Cannot delete car order",
            Message = $"Car order with ID {request.Id} was not found in the database."
        };
            
        await _commandRepository.DeleteAsync(carOrder.Id, cancellationToken);

        return new DeleteCarOrderResponse() {
            Title = "Car order deleted",
            Message = $"Car order with ID {request.Id} was deleted from the database."
        };
    }
}