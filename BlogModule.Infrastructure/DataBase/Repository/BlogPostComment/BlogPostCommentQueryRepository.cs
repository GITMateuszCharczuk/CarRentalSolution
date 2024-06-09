using System.Collections.Immutable;
using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using BlogModule.Domain.RepositoryInterfaces.BlogPostComment;
using BlogModule.Infrastructure.DataBase.Context;
using BlogModule.Infrastructure.DataBase.Entities;
using Microsoft.EntityFrameworkCore;
using Shared.CQRS.Repository;
using Shared.Utilities;

namespace BlogModule.Infrastructure.DataBase.Repository.BlogPostComment;

public class BlogPostCommentQueryRepository : QueryRepository<BlogPostCommentEntity, BlogPostCommentId, BlogPostCommentModel, BlogDbContext>, IBlogPostCommentQueryRepository
{
    public BlogPostCommentQueryRepository(BlogDbContext dbContext, IPersistenceMapper<BlogPostCommentEntity, BlogPostCommentModel> mapper) : base(dbContext, mapper)
    {
    }
    
    public override async Task<BlogPostCommentModel?> GetByIdAsync(BlogPostCommentId id, CancellationToken cancellationToken = default) => 
        await base.GetByIdAsync(id, cancellationToken);
    
    public override async Task<ImmutableArray<BlogPostCommentModel>> GetByIdsAsync(ImmutableArray<BlogPostCommentId> ids, CancellationToken cancellationToken = default) => 
        await base.GetByIdsAsync(ids, cancellationToken);

    public override async Task<int> GetTotalCountAsync(CancellationToken cancellationToken = default) => 
        await base.GetTotalCountAsync(cancellationToken);

    public async Task<ImmutableArray<BlogPostCommentModel>> GetByBlogPostIdAsync(BlogPostId blogPostId, CancellationToken cancellationToken) =>
        await DbContext.BlogPostComment
            .AsNoTracking()
            .Where(x => x.BlogPostId == blogPostId)
            .Select(bp => Mapper.MapToModel(bp))
            .ToImmutableArrayAsync(cancellationToken);
}