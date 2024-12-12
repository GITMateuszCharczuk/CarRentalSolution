using System.Collections.Immutable;
using BlogModule.Domain.Enums;
using BlogModule.Domain.Models;
using Results.Contract;
using Results.Domain;

namespace BlogModule.Application.Contract.BlogPosts.GetBlogPosts;

public record GetBlogPostsResponse : IApiCollectionResult<BlogPostModel>, IPageableResponse, ISortable<BlogPostSortColumnEnum?>, ISuccessResult
{
    public int? Page { get; init; }
    public int? PageSize { get; init; }
    public int TotalCount { get; init; }
    public BlogPostSortColumnEnum? OrderBy { get; init; }
    public SortOrderEnum? OrderDirection { get; init; }
    public ImmutableArray<BlogPostModel> Items { get; init; }
}