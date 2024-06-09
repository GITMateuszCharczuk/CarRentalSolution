using BlogModule.Domain.Models;
using BlogModule.Infrastructure.DataBase.Entities;
using Shared.Utilities;

namespace BlogModule.Infrastructure.Mappers
{
    public class BlogPostCommentPersistenceMapper : IPersistenceMapper<BlogPostCommentEntity, BlogPostCommentModel>
    {
        public BlogPostCommentModel MapToModel(BlogPostCommentEntity entity) => new()
        {
            Id = entity.Id,
            Description = entity.Description,
            BlogPostId = entity.BlogPostId,
            UserId = entity.UserId,
            DateAdded = entity.DateAdded
        };

        public BlogPostCommentEntity MapToEntity(BlogPostCommentModel model) => new()
        {
            Id = model.Id,
            Description = model.Description,
            BlogPostId = model.BlogPostId,
            UserId = model.UserId,
            DateAdded = model.DateAdded
        };
    }
}