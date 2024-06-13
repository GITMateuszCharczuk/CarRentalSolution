using System.Collections.Immutable;
using BlogModule.Domain.Enums;
using BlogModule.Domain.Models;
using Results.Contract;
using Results.Domain;

namespace BlogModule.Application.Contract.Tags.GetTags;

public class GetTagsResponse : IApiCollectionResult<TagModel>, ISortable<TagSortColumnEnum?>, ISuccessResult
{
     public TagSortColumnEnum? OrderBy { get; init; }
     public SortOrderEnum? OrderDirection { get; init; }
     public ImmutableArray<TagModel> Items { get; init; }
}