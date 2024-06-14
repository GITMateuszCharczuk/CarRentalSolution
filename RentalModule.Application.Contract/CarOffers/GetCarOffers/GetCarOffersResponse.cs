using System.Collections.Immutable;
using RentalModule.Domain.Enums;
using RentalModule.Domain.Models;
using Results.Contract;
using Results.Domain;

namespace RentalModule.Application.Contract.CarOffers.GetCarOffers;

public record GetCarOffersResponse : IApiCollectionResult<CarOfferModel>, IPageableResponse, ISortable<CarOfferSortColumnEnum?>, ISuccessResult
{
    public int? Page { get; init; }
    public int? PageSize { get; init; }
    public int TotalCount { get; init; }
    public CarOfferSortColumnEnum? OrderBy { get; init; }
    public SortOrderEnum? OrderDirection { get; init; }
    public ImmutableArray<CarOfferModel> Items { get; init; }
}