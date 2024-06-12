using System.Collections.Immutable;
using BlogModule.Domain.Enums;
using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using BlogModule.Domain.RepositoryInterfaces.BlogPost;
using BlogModule.Infrastructure.DataBase.Context;
using BlogModule.Infrastructure.DataBase.Entities;
using Microsoft.EntityFrameworkCore;
using Results.Contract;
using Shared.CQRS.Repository;
using Shared.Utilities;

namespace BlogModule.Infrastructure.DataBase.Repository.BlogPost;

public class BlogPostQueryRepository : QueryRepository<BlogPostEntity, BlogPostId, BlogPostModel, BlogDbContext>,
    IBlogPostQueryRepository
{
    public BlogPostQueryRepository(BlogDbContext dbContext, IPersistenceMapper<BlogPostEntity, BlogPostModel> mapper) :
        base(dbContext, mapper)
    {
    }

    public override async Task<BlogPostModel?> GetByIdAsync(BlogPostId id,
        CancellationToken cancellationToken = default) =>
        await base.GetByIdAsync(id, cancellationToken);

    public override async Task<ImmutableArray<BlogPostModel>> GetByIdsAsync(ImmutableArray<BlogPostId> ids,
        CancellationToken cancellationToken = default) =>
        await base.GetByIdsAsync(ids, cancellationToken);

    public override async Task<int> GetTotalCountAsync(CancellationToken cancellationToken = default) =>
        await base.GetTotalCountAsync(cancellationToken);

    public async Task<BlogPostModel?> GetByUrlAsync(string urlHandle, CancellationToken cancellationToken) =>
        await DbContext.BlogPosts.AsNoTracking()
            .FirstOrDefaultAsync(x => x.UrlHandle == urlHandle, cancellationToken)
            .ContinueWith(x => x.Result is null ? null : Mapper.MapToModel(x.Result), cancellationToken);

    public async Task<ImmutableArray<BlogPostModel>> GetByTagAsync(string tagName, CancellationToken cancellationToken)
    {
        return await DbContext.BlogPosts
            .Include(bp => bp.Tags)
            .Where(bp => bp.Tags.Any(tag => tag.Name == tagName))
            .Select(bp => Mapper.MapToModel(bp))
            .ToImmutableArrayAsync(cancellationToken);
    }

    public async Task<ImmutableArray<BlogPostModel>> GetAllAsync(CancellationToken cancellationToken) =>
        await DbContext.BlogPosts
            .AsNoTracking()
            .Select(bp => Mapper.MapToModel(bp))
            .ToImmutableArrayAsync(cancellationToken);

    public async Task<ImmutableArray<BlogPostModel>> GetCollectionAsync(int? page, int? pageSize,
        BlogPostSortColumnEnum? orderBy, SortOrderEnum? orderDirection,
        ImmutableArray<BlogPostId>? ids, ImmutableArray<DateTime>? publishedDates, ImmutableArray<string>? authors,
        CancellationToken cancellationToken)
    {
        var queryablePosts = DbContext.BlogPosts
            .AsNoTracking()
            .AsQueryable();

        if (ids is not null && ids.Any<BlogPostId>())
        {
            queryablePosts = queryablePosts.Where(x => ids.Contains(x.Id));
        }

        if (publishedDates is not null && publishedDates.Any<DateTime>())
        {
            queryablePosts = queryablePosts.Where(x => publishedDates.Contains(x.PublishedDate));
        }

        if (authors is not null && authors.Any<string>())
        {
            queryablePosts = queryablePosts.Where(x => authors.Contains(x.Author));
        }

        if (orderBy is not null && orderDirection is not null)
        {
            var isOrderDirectionAscending = orderDirection == SortOrderEnum.Ascending;
            queryablePosts = orderBy switch
            {
                BlogPostSortColumnEnum.PublishedDate => isOrderDirectionAscending
                    ? queryablePosts.OrderBy(x => x.PublishedDate)
                    : queryablePosts.OrderByDescending(x => x.PublishedDate),
                BlogPostSortColumnEnum.Author => isOrderDirectionAscending
                    ? queryablePosts.OrderBy(x => x.Author)
                    : queryablePosts.OrderByDescending(x => x.Author),
                BlogPostSortColumnEnum.Heading => isOrderDirectionAscending
                    ? queryablePosts.OrderBy(x => x.Heading)
                    : queryablePosts.OrderByDescending(x => x.Heading),
                _ => queryablePosts
            };
        }

        if (page is not null && pageSize is not null)
        {
            queryablePosts = queryablePosts
                .Skip((page.Value - 1) * pageSize.Value)
                .Take(pageSize.Value);
        }

        return await queryablePosts
            .Select(x => Mapper.MapToModel(x))
            .ToImmutableArrayAsync(cancellationToken);
    }
}