using System.Collections.Immutable;
using RentalModule.Application.Contract.CarOffers.GetCarOffers;
using RentalModule.Application.Contract.CarOrders.GetCarOrders;
using RentalModule.Domain.Enums;
using RentalModule.Domain.Models.Ids;
using Results.Contract;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.QueryHandlers;

namespace RentalModule.Application.QueryHandlers.CarOrder.GetCarOrders;

public record GetCarOrdersQuery : IPageableRequest, ISortable<CarOrderSortColumnEnum?>, IQuery<HandlerResult<GetCarOrdersResponse, IErrorResult>>
{
    public int? Page { get; init; }
    
    public int? PageSize { get; init; }
    
    public CarOrderSortColumnEnum? OrderBy { get; init; }
    
    public SortOrderEnum? OrderDirection { get; init; }
    
    public ImmutableArray<DateTime>? Dates { get; init; }

    public Guid UserId { get; init; }

    public CarOfferId CarOfferId { get; init; }
}