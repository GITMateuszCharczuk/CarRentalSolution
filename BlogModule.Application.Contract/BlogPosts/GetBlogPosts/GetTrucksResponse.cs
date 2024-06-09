using System.Collections.Immutable;
using TrucksModule.Application.Contract.Base;
using TrucksModule.Domain.Enums;
using TrucksModule.Domain.Models;
using TrucksModule.Domain.Results.Base;

namespace TrucksModule.Application.Contract.Trucks.GetTrucks;

public record GetTrucksResponse : IApiCollectionResult<TruckModel>, IPageableResponse, ISortable<TruckSortColumnEnum?>, ISuccessResult
{
    public int? Page { get; init; }
    public int? PageSize { get; init; }
    public int TotalCount { get; init; }
    public TruckSortColumnEnum? OrderBy { get; init; }
    public SortOrderEnum? OrderDirection { get; init; }
    public ImmutableArray<TruckModel> Items { get; init; }
}