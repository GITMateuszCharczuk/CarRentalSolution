using System.Collections.Immutable;
using BlogModule.Domain.Enums;
using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using Results.Contract;

namespace BlogModule.Domain.RepositoryInterfaces.BlogPostComment;

public interface IBlogPostCommentQueryRepository
{
    public Task<BlogPostCommentModel?> GetByIdAsync(BlogPostCommentId id, CancellationToken cancellationToken);

    public Task<ImmutableArray<BlogPostCommentModel>> GetByIdsAsync(ImmutableArray<BlogPostCommentId> ids,
        CancellationToken cancellationToken);

    public Task<int> GetTotalCommentsCountAsync(
        BlogPostId? blogPostId,
        ImmutableArray<BlogPostCommentId>? ids,
        ImmutableArray<DateTime>? dateTimes,
        ImmutableArray<Guid>? userIds,
        CancellationToken cancellationToken);

    public Task<ImmutableArray<BlogPostCommentModel>> GetByBlogPostIdAsync(BlogPostId blogPostId,
        CancellationToken cancellationToken);

    public Task<ImmutableArray<BlogPostCommentModel>> GetCollectionAsync(int? page, int? pageSize,
        BlogPostId? blogPostId, BlogPostCommentSortColumnEnum? orderBy,
        SortOrderEnum? orderDirection, ImmutableArray<BlogPostCommentId>? ids,
        ImmutableArray<DateTime>? dateTimes,
        ImmutableArray<Guid>? userIds,
        CancellationToken cancellationToken);
}