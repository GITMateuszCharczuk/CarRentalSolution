using System.Collections.Immutable;
using RentalModule.Domain.Enums;
using RentalModule.Domain.Models;
using Results.Contract;
using Results.Domain;

namespace RentalModule.Application.Contract.CarOrders.GetCarOrders;

public record GetCarOrdersResponse : IApiCollectionResult<CarOrderModel>, IPageableResponse, ISortable<CarOrderSortColumnEnum?>, ISuccessResult
{
    public int? Page { get; init; }
    public int? PageSize { get; init; }
    public int TotalCount { get; init; }
    public CarOrderSortColumnEnum? OrderBy { get; init; }
    public SortOrderEnum? OrderDirection { get; init; }
    public ImmutableArray<CarOrderModel> Items { get; init; }
}