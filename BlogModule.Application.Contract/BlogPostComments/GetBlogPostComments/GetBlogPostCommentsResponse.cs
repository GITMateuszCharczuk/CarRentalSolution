using System.Collections.Immutable;
using BlogModule.Domain.Enums;
using BlogModule.Domain.Models;
using Results.Contract;
using Results.Domain;

namespace BlogModule.Application.Contract.BlogPostComments.GetBlogPostComments;

public class GetBlogPostCommentsResponse: IApiCollectionResult<BlogPostCommentModel>, IPageableResponse, ISortable<BlogPostCommentSortColumnEnum?>, ISuccessResult
{
    public int? Page { get; init; }
    public int? PageSize { get; init; }
    public int TotalCount { get; init; }
    public BlogPostCommentSortColumnEnum? OrderBy { get; init; }
    public SortOrderEnum? OrderDirection { get; init; }
    public ImmutableArray<BlogPostCommentModel> Items { get; init; }
}