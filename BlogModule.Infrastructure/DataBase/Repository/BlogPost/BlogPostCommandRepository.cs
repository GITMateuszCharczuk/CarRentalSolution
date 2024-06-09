using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using BlogModule.Domain.RepositoryInterfaces.BlogPost;
using BlogModule.Infrastructure.DataBase.Context;
using BlogModule.Infrastructure.DataBase.Entities;
using Shared.CQRS.Repository;
using Shared.Utilities;

namespace BlogModule.Infrastructure.DataBase.Repository.BlogPost;

public class BlogPostCommandRepository : CommandRepository<BlogPostEntity, BlogPostId, BlogPostModel, BlogDbContext>, IBlogPostCommandRepository
{
    public BlogPostCommandRepository(BlogDbContext dbContext, IPersistenceMapper<BlogPostEntity, BlogPostModel> mapper) : base(dbContext, mapper)
    {
    }

    public override async Task<BlogPostModel> AddAsync(BlogPostModel model, CancellationToken cancellationToken = default) =>
        await base.AddAsync(model, cancellationToken);

    public override async Task<BlogPostModel> UpdateAsync(BlogPostModel model, CancellationToken cancellationToken = default) =>
        await base.UpdateAsync(model, cancellationToken);

    public override async Task DeleteAsync(BlogPostId id, CancellationToken cancellationToken = default) =>
        await base.DeleteAsync(id, cancellationToken);
}