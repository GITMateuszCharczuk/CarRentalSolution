using System.Collections.Immutable;
using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;

namespace BlogModule.Domain.RepositoryInterfaces.BlogPostLike;

public interface IBlogPostLikeQueryRepository
{
    public Task<BlogPostLikeModel?> GetByIdAsync(BlogPostLikeId id, CancellationToken cancellationToken);

    public Task<ImmutableArray<BlogPostLikeModel>> GetByIdsAsync(ImmutableArray<BlogPostLikeId> ids, CancellationToken cancellationToken);
    
    public Task<int> GetTotalCountAsync(CancellationToken cancellationToken);

    public Task<int> GetTotalLikesForBlogPostAsync(BlogPostId blogPostId, CancellationToken cancellationToken);

    public Task<ImmutableArray<BlogPostLikeModel>> GetLikesForBlogPostAsync(BlogPostId blogPostId,
        CancellationToken cancellationToken);
    // public Task<ImmutableArray<BlogPostLikeModel>> GetCollectionAsync(
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