using BlogModule.Domain.Models;
using BlogModule.Infrastructure.DataBase.Entities;
using Shared.Utilities;

namespace BlogModule.Infrastructure.Mappers
{
    public class BlogPostLikePersistenceMapper : IPersistenceMapper<BlogPostLikeEntity, BlogPostLikeModel>
    {
        public BlogPostLikeModel MapToModel(BlogPostLikeEntity entity) => new()
        {
            Id = entity.Id,
            BlogPostId = entity.BlogPostId,
            UserId = entity.UserId
        };

        public BlogPostLikeEntity MapToEntity(BlogPostLikeModel model) => new()
        {
            Id = model.Id,
            BlogPostId = model.BlogPostId,
            UserId = model.UserId
        };
    }
}