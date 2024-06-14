using RentalModule.Application.Contract.CarOrders.DeleteCarOrder;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.CommandHandlers;

namespace RentalModule.Application.CommandHandlers.CarOrder.DeleteCarOrder;

public class DeleteCarOrderCommand : ICommand<HandlerResult<DeleteCarOrderResponse, IErrorResult>>
{
    public string Id { get; init; } = string.Empty;
}