using System.Collections.Immutable;
using BlogModule.Domain.Enums;
using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using Results.Contract;

namespace BlogModule.Domain.RepositoryInterfaces.BlogPost;

public interface IBlogPostQueryRepository
{
    public Task<BlogPostModel?> GetByIdAsync(BlogPostId id, CancellationToken cancellationToken);

    public Task<ImmutableArray<BlogPostModel>> GetByIdsAsync(ImmutableArray<BlogPostId> ids, CancellationToken cancellationToken);
    
    public Task<int> GetTotalCountAsync(CancellationToken cancellationToken);
    public Task<BlogPostModel?> GetByUrlAsync(string urlHandle, CancellationToken cancellationToken);
    public Task<ImmutableArray<BlogPostModel>> GetByTagAsync(string tagName, CancellationToken cancellationToken);
    public Task<ImmutableArray<BlogPostModel>> GetAllAsync(CancellationToken cancellationToken);
    
    public Task<ImmutableArray<BlogPostModel>> GetCollectionAsync(
        int? page,
        int? pageSize,
        BlogPostSortColumnEnum? orderBy,
        SortOrderEnum? orderDirection,
        ImmutableArray<BlogPostId>? ids,
        ImmutableArray<DateTime>? publishedDates,
        ImmutableArray<string>? authors,
        CancellationToken cancellationToken);
}