using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using BlogModule.Domain.RepositoryInterfaces.BlogPostComment;
using BlogModule.Infrastructure.DataBase.Context;
using BlogModule.Infrastructure.DataBase.Entities;
using Shared.CQRS.Repository;
using Shared.Utilities;

namespace BlogModule.Infrastructure.DataBase.Repository.BlogPostComment;

public class BlogPostCommentCommandRepository : CommandRepository<BlogPostCommentEntity, BlogPostCommentId, BlogPostCommentModel, BlogDbContext>, IBlogPostCommentCommandRepository
{
    public BlogPostCommentCommandRepository(BlogDbContext dbContext, IPersistenceMapper<BlogPostCommentEntity, BlogPostCommentModel> mapper) : base(dbContext, mapper)
    {
    }

    public override async Task<BlogPostCommentModel> AddAsync(BlogPostCommentModel model, CancellationToken cancellationToken = default) =>
        await base.AddAsync(model, cancellationToken);

    public override async Task<BlogPostCommentModel> UpdateAsync(BlogPostCommentModel model, CancellationToken cancellationToken = default) =>
        await base.UpdateAsync(model, cancellationToken);

    public override async Task DeleteAsync(BlogPostCommentId id, CancellationToken cancellationToken = default) =>
        await base.DeleteAsync(id, cancellationToken);
    
}