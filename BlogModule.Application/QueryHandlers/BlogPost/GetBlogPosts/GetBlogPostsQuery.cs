using System.Collections.Immutable;
using BlogModule.Application.Contract.BlogPosts.GetBlogPosts;
using BlogModule.Domain.Enums;
using BlogModule.Domain.Models.Ids;
using Results.Contract;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.QueryHandlers;

namespace BlogModule.Application.QueryHandlers.BlogPost.GetBlogPosts;

public record GetBlogPostsQuery : IPageableRequest, ISortable<BlogPostSortColumnEnum?>, IQuery<HandlerResult<GetBlogPostsResponse, IErrorResult>>
{
    public int? Page { get; init; }
    public int? PageSize { get; init; }
    public BlogPostSortColumnEnum? OrderBy { get; init; }
    public SortOrderEnum? OrderDirection { get; init; }
    public ImmutableArray<BlogPostId>? Ids { get; init; }
    public ImmutableArray<DateTime>? PublishedDates { get; init; }
    public ImmutableArray<string>? Authors { get; init; }
}