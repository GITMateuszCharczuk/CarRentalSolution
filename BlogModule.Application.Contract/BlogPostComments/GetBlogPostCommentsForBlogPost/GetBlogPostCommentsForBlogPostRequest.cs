using BlogModule.Domain.Enums;
using Results.Contract;

namespace BlogModule.Application.Contract.BlogPostComments.GetBlogPostCommentsForBlogPost;

public class GetBlogPostCommentsForBlogPostRequest : IPageableRequest, ISortable<BlogPostCommentSortColumnEnum?>
{
    public int? Page { get; init; }
    public int? PageSize { get; init; }
    public BlogPostCommentSortColumnEnum? OrderBy { get; init; }
    public SortOrderEnum? OrderDirection { get; init; }
}