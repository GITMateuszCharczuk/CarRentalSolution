using System.Collections.Immutable;
using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using BlogModule.Domain.RepositoryInterfaces.BlogPost;
using BlogModule.Domain.RepositoryInterfaces.BlogPostLike;
using BlogModule.Infrastructure.DataBase.Context;
using BlogModule.Infrastructure.DataBase.Entities;
using Microsoft.EntityFrameworkCore;
using Shared.CQRS.Repository;
using Shared.Utilities;

namespace BlogModule.Infrastructure.DataBase.Repository.BlogPostLike;

public class BlogPostLikeQueryRepository : QueryRepository<BlogPostLikeEntity, BlogPostLikeId, BlogPostLikeModel, BlogDbContext>, IBlogPostLikeQueryRepository
{
    public BlogPostLikeQueryRepository(BlogDbContext dbContext, IPersistenceMapper<BlogPostLikeEntity, BlogPostLikeModel> mapper) : base(dbContext, mapper)
    {
    }
    
    public override async Task<BlogPostLikeModel?> GetByIdAsync(BlogPostLikeId id, CancellationToken cancellationToken = default) => 
        await base.GetByIdAsync(id, cancellationToken);
    
    public override async Task<ImmutableArray<BlogPostLikeModel>> GetByIdsAsync(ImmutableArray<BlogPostLikeId> ids, CancellationToken cancellationToken = default) => 
        await base.GetByIdsAsync(ids, cancellationToken);

    public override async Task<int> GetTotalCountAsync(CancellationToken cancellationToken = default) => 
        await base.GetTotalCountAsync(cancellationToken);
    
    public async Task<int> GetTotalLikesForBlogPostAsync(BlogPostId blogPostId, CancellationToken cancellationToken) =>
        await DbContext.BlogPostLikes
            .CountAsync(x => x.BlogPostId == blogPostId, cancellationToken);

    public async Task<ImmutableArray<BlogPostLikeModel>> GetLikesForBlogPostAsync(BlogPostId blogPostId, CancellationToken cancellationToken) =>
        await DbContext.BlogPostLikes
            .AsNoTracking()
            .Where(x => x.BlogPostId == blogPostId)
            .Select(bp => Mapper.MapToModel(bp))
            .ToImmutableArrayAsync(cancellationToken);
}