using System.Collections.Immutable;
using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;

namespace BlogModule.Domain.RepositoryInterfaces.BlogPostComment;

public interface IBlogPostCommentQueryRepository
{
    public Task<BlogPostCommentModel?> GetByIdAsync(BlogPostCommentId id, CancellationToken cancellationToken);

    public Task<ImmutableArray<BlogPostCommentModel>> GetByIdsAsync(ImmutableArray<BlogPostCommentId> ids, CancellationToken cancellationToken);
    
    public Task<int> GetTotalCountAsync(CancellationToken cancellationToken);

    public Task<ImmutableArray<BlogPostCommentModel>> GetByBlogPostIdAsync(BlogPostId blogPostId,
        CancellationToken cancellationToken);
    // public Task<ImmutableArray<TruckModel>> GetCollectionAsync(
    //     int? page,
    //     int? pageSize,
    //     TruckSortColumnEnum? orderBy,
    //     SortOrderEnum? orderDirection,
    //     ImmutableArray<TruckId>? ids,
    //     ImmutableArray<string>? codes,
    //     ImmutableArray<string>? names,
    //     ImmutableArray<TruckStatusEnum>? statuses,
    //     CancellationToken cancellationToken);
}