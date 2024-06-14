using System.Collections.Immutable;
using RentalModule.Application.Contract.CarOffers.GetCarOffers;
using RentalModule.Domain.Enums;
using Results.Contract;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.QueryHandlers;

namespace RentalModule.Application.QueryHandlers.CarOffer.GetCarOffers;

public record GetCarOffersQuery : IPageableRequest, ISortable<CarOfferSortColumnEnum?>, IQuery<HandlerResult<GetCarOffersResponse, IErrorResult>>
{
    public int? Page { get; init; }
    
    public int? PageSize { get; init; }
    
    public CarOfferSortColumnEnum? OrderBy { get; init; }
    
    public SortOrderEnum? OrderDirection { get; init; }
    
    public ImmutableArray<DateTime>? PossibleDates { get; init; }
    
    public ImmutableArray<string>? Tags { get; init; }
}