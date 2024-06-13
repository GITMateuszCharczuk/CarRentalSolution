using BlogModule.Domain.Models;
using BlogModule.Infrastructure.DataBase.Entities;
using Shared.Utilities;

namespace BlogModule.Infrastructure.Mappers
{
    public class BlogPostPersistenceMapper : IPersistenceMapper<BlogPostEntity, BlogPostModel>
    {
        private readonly IPersistenceMapper<BlogPostCommentEntity, BlogPostCommentModel> _commentMapper;
        private readonly IPersistenceMapper<BlogPostLikeEntity, BlogPostLikeModel> _likeMapper;
        private readonly IPersistenceMapper<TagEntity, TagModel> _tagMapper;

        public BlogPostPersistenceMapper(
            IPersistenceMapper<BlogPostCommentEntity, BlogPostCommentModel> commentMapper,
            IPersistenceMapper<BlogPostLikeEntity, BlogPostLikeModel> likeMapper,
            IPersistenceMapper<TagEntity, TagModel> tagMapper)
        {
            _tagMapper = tagMapper;
            _likeMapper = likeMapper;
            _commentMapper = commentMapper;
        }

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
            Tags = entity.Tags?.Select(tag => _tagMapper.MapToModel(tag)).ToArray() ?? Array.Empty<TagModel>(),
            Likes = entity.Likes?.Select(like => _likeMapper.MapToModel(like)).ToArray() ?? Array.Empty<BlogPostLikeModel>(),
            Comments = entity.Comments?.Select(comment => _commentMapper.MapToModel(comment)).ToArray() ?? Array.Empty<BlogPostCommentModel>()
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
            Tags = model.Tags?.Select(tag => _tagMapper.MapToEntity(tag)).ToArray() ?? Array.Empty<TagEntity>(),
            Likes = model.Likes?.Select(like => _likeMapper.MapToEntity(like)).ToArray() ?? Array.Empty<BlogPostLikeEntity>(),
            Comments = model.Comments?.Select(comment => _commentMapper.MapToEntity(comment)).ToArray() ?? Array.Empty<BlogPostCommentEntity>()
        };
    }
}