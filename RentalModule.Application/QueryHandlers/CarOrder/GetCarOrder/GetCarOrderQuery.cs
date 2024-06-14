using RentalModule.Application.Contract.CarOffers.GetCarOffer;
using RentalModule.Application.Contract.CarOrders.GetCarOrder;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.QueryHandlers;

namespace RentalModule.Application.QueryHandlers.CarOrder.GetCarOrder;

public record GetCarOrderQuery : IQuery<HandlerResult<GetCarOrderResponse, IErrorResult>>
{
    public string IdOrUrl { get; init; } = string.Empty;
}