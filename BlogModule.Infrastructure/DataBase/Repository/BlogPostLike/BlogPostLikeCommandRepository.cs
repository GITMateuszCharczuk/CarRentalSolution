using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using BlogModule.Domain.RepositoryInterfaces.BlogPostLike;
using BlogModule.Infrastructure.DataBase.Context;
using BlogModule.Infrastructure.DataBase.Entities;
using Shared.CQRS.Repository;
using Shared.Utilities;

namespace BlogModule.Infrastructure.DataBase.Repository.BlogPostLike;

public class BlogPostLikeCommandRepository : CommandRepository<BlogPostLikeEntity, BlogPostLikeId, BlogPostLikeModel, BlogDbContext>, IBlogPostLikeCommandRepository
{
    public BlogPostLikeCommandRepository(BlogDbContext dbContext, IPersistenceMapper<BlogPostLikeEntity, BlogPostLikeModel> mapper) : base(dbContext, mapper)
    {

    }

    public override async Task<BlogPostLikeModel> AddAsync(BlogPostLikeModel model, CancellationToken cancellationToken = default) =>
        await base.AddAsync(model, cancellationToken);

    public override async Task<BlogPostLikeModel> UpdateAsync(BlogPostLikeModel model, CancellationToken cancellationToken = default) =>
        await base.UpdateAsync(model, cancellationToken);

    public override async Task DeleteAsync(BlogPostLikeId id, CancellationToken cancellationToken = default) =>
        await base.DeleteAsync(id, cancellationToken);
    public async Task AddLikeForBlogAsync(BlogPostId blogPostId, Guid userId)
    {
        var likeModel = new BlogPostLikeModel()
        {
            Id = new BlogPostLikeId(Guid.NewGuid()),
            BlogPostId = blogPostId,
            UserId = userId
        };
        var likeEntity = Mapper.MapToEntity(likeModel);
        await DbContext.BlogPostLikes.AddAsync(likeEntity);
        await DbContext.SaveChangesAsync();
    }
}