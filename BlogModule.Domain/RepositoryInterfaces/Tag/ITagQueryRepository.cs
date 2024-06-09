using System.Collections.Immutable;
using BlogModule.Domain.Models.Ids;
using CarRental.Web.Models.Domain.Blog;

namespace BlogModule.Domain.RepositoryInterfaces.BlogPost;

public interface ITagQueryRepository
{
    public Task<TagModel?> GetByIdAsync(TagId id, CancellationToken cancellationToken);

    public Task<ImmutableArray<TagModel>> GetByIdsAsync(ImmutableArray<TagId> ids, CancellationToken cancellationToken);
    
    public Task<int> GetTotalCountAsync(CancellationToken cancellationToken);
    
    public Task<ImmutableArray<TagModel>> GetAllDistinctAsync(CancellationToken cancellationToken);
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