using System.Collections.Immutable;
using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using BlogModule.Domain.RepositoryInterfaces.BlogPost;
using BlogModule.Infrastructure.DataBase.Context;
using BlogModule.Infrastructure.DataBase.Entities;
using Microsoft.EntityFrameworkCore;
using Shared.CQRS.Repository;
using Shared.Utilities;

namespace BlogModule.Infrastructure.DataBase.Repository.BlogPost;

public class BlogPostQueryRepository : QueryRepository<BlogPostEntity, BlogPostId, BlogPostModel, BlogDbContext>, IBlogPostQueryRepository
{
    public BlogPostQueryRepository(BlogDbContext dbContext, IPersistenceMapper<BlogPostEntity, BlogPostModel> mapper) : base(dbContext, mapper)
    {
    }
    
    public override async Task<BlogPostModel?> GetByIdAsync(BlogPostId id, CancellationToken cancellationToken = default) => 
        await base.GetByIdAsync(id, cancellationToken);
    
    public override async Task<ImmutableArray<BlogPostModel>> GetByIdsAsync(ImmutableArray<BlogPostId> ids, CancellationToken cancellationToken = default) => 
        await base.GetByIdsAsync(ids, cancellationToken);

    public override async Task<int> GetTotalCountAsync(CancellationToken cancellationToken = default) => 
        await base.GetTotalCountAsync(cancellationToken);
    
    public async Task<BlogPostModel?> GetByUrlHandleAsync(string urlHandle, CancellationToken cancellationToken) =>
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
}
