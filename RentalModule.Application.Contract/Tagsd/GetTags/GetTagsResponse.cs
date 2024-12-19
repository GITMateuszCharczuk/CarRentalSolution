using System.Collections.Immutable;
using BlogModule.Domain.Enums;
using RentalModule.Domain.Models;
using Results.Contract;
using Results.Domain;

namespace RentalModule.Application.Contract.Tags.GetTags;

public class GetTagsResponse : IApiCollectionResult<CarTagModel>, ISortable<CarTagSortColumnEnum?>, ISuccessResult
{
     public CarTagSortColumnEnum? OrderBy { get; init; }
     public SortOrderEnum? OrderDirection { get; init; }
     public ImmutableArray<CarTagModel> Items { get; init; }
}