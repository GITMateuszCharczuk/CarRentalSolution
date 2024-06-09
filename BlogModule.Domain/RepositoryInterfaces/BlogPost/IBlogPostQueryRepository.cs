using System.Collections.Immutable;
using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using TrucksModule.Domain.Enums;

namespace BlogModule.Domain.RepositoryInterfaces.BlogPost;

public interface IBlogPostQueryRepository
{
    public Task<BlogPostModel?> GetByIdAsync(BlogPostId id, CancellationToken cancellationToken);

    public Task<ImmutableArray<BlogPostModel>> GetByIdsAsync(ImmutableArray<BlogPostId> ids, CancellationToken cancellationToken);
    
    public Task<int> GetTotalCountAsync(CancellationToken cancellationToken);
    public Task<BlogPostModel?> GetByUrlHandleAsync(string urlHandle, CancellationToken cancellationToken);
    public Task<ImmutableArray<BlogPostModel>> GetByTagAsync(string tagName, CancellationToken cancellationToken);
    public Task<ImmutableArray<BlogPostModel>> GetAllAsync(CancellationToken cancellationToken);
    // public Task<ImmutableArray<BlogPostModel>> GetCollectionAsync(
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