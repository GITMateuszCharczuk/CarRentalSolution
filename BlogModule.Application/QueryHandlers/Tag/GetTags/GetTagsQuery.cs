using System.Collections.Immutable;
using BlogModule.Application.Contract.Tags.GetTags;
using BlogModule.Domain.Enums;
using BlogModule.Domain.Models.Ids;
using Results.Contract;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.QueryHandlers;

namespace BlogModule.Application.QueryHandlers.Tag.GetTags;

public record GetTagsQuery : ISortable<TagSortColumnEnum?>, IQuery<HandlerResult<GetTagsResponse, IErrorResult>>

{
    public TagSortColumnEnum? OrderBy { get; init; }
    public SortOrderEnum? OrderDirection { get; init; }
    public BlogPostId? BlogPostId { get; init; } 
}