using BlogModule.Domain.Models;
using BlogModule.Infrastructure.DataBase.Entities;
using Shared.Utilities;

namespace BlogModule.Infrastructure.Mappers
{
    public class TagPersistenceMapper : IPersistenceMapper<TagEntity, TagModel>
    {
        public TagModel MapToModel(TagEntity entity) => new()
        {
            Id = entity.Id,
            Name = entity.Name,
            BlogPostId = entity.BlogPostId
        };

        public TagEntity MapToEntity(TagModel model) => new()
        {
            Id = model.Id,
            Name = model.Name,
            BlogPostId = model.BlogPostId
        };
    }
}