using BlogModule.Domain.Models;
using BlogModule.Infrastructure.DataBase.Entities;
using Shared.Utilities;
 
namespace BlogModule.Infrastructure.Mappers
{
    public class BlogPostPersistenceMapper : IPersistenceMapper<BlogPostEntity, BlogPostModel>
    {
        public BlogPostModel MapToModel(BlogPostEntity entity) => new()
        {
            Id = entity.Id,
            Heading = entity.Heading,
            PageTitle = entity.PageTitle,
            Content = entity.Content,
            ShortDescription = entity.ShortDescription,
            FeaturedImageUrl = entity.FeaturedImageUrl,
            UrlHandle = entity.UrlHandle,
            PublishedDate = entity.PublishedDate,
            Author = entity.Author,
            Visible = entity.Visible,
            Tags = entity.Tags.ToList(),
            Likes = entity.Likes.ToList(),
            Comments = entity.Comments.ToList()
        };

        public BlogPostEntity MapToEntity(BlogPostModel model) => new()
        {
            Id = model.Id,
            Heading = model.Heading,
            PageTitle = model.PageTitle,
            Content = model.Content,
            ShortDescription = model.ShortDescription,
            FeaturedImageUrl = model.FeaturedImageUrl,
            UrlHandle = model.UrlHandle,
            PublishedDate = model.PublishedDate,
            Author = model.Author,
            Visible = model.Visible,
            Tags = model.Tags.ToList(),
            Likes = model.Likes.ToList(),
            Comments = model.Comments.ToList()
        };
    }
}