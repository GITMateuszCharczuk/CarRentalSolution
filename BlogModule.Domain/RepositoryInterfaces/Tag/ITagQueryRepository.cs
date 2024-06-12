using System.Collections.Immutable;
using BlogModule.Domain.Enums;
using BlogModule.Domain.Models.Ids;
using CarRental.Web.Models.Domain.Blog;
using Results.Contract;

namespace BlogModule.Domain.RepositoryInterfaces.Tag;

public interface ITagQueryRepository
{
    public Task<TagModel?> GetByIdAsync(TagId id, CancellationToken cancellationToken);

    public Task<ImmutableArray<TagModel>> GetByIdsAsync(ImmutableArray<TagId> ids, CancellationToken cancellationToken);

    public Task<int> GetTotalCountAsync(CancellationToken cancellationToken);

    public Task<ImmutableArray<TagModel>> GetAllDistinctAsync(TagSortColumnEnum? orderBy,
        SortOrderEnum? orderDirection,BlogPostId? blogPostId, CancellationToken cancellationToken);
}