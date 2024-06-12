using BlogModule.Domain.Models;
using BlogModule.Infrastructure.DataBase.Entities;
using CarRental.Web.Models.Domain.Blog;
using Shared.Utilities;
 
namespace BlogModule.Infrastructure.Mappers
{
    public class BlogPostPersistenceMapper : IPersistenceMapper<BlogPostEntity, BlogPostModel>
    {
        private readonly IPersistenceMapper<BlogPostCommentModel, BlogPostCommentEntity> _commentMapperToEntity;
        private readonly IPersistenceMapper<BlogPostLikeModel, BlogPostLikeEntity> _likeMapperToEntity;
        private readonly IPersistenceMapper<TagModel, TagEntity> _tagMapperToEntity;
        private readonly IPersistenceMapper<BlogPostCommentEntity, BlogPostCommentModel> _commentMapperToModel;
        private readonly IPersistenceMapper<BlogPostLikeEntity, BlogPostLikeModel> _likeMapperToModel;
        private readonly IPersistenceMapper<TagEntity, TagModel> _tagMapperToModel;

        public BlogPostPersistenceMapper(IPersistenceMapper<BlogPostCommentModel, BlogPostCommentEntity> commentMapperToEntity,
            IPersistenceMapper<BlogPostLikeModel, BlogPostLikeEntity> likeMapperToEntity,
            IPersistenceMapper<TagModel, TagEntity> tagMapperToEntity,
            IPersistenceMapper<BlogPostCommentEntity, BlogPostCommentModel> commentMapperToModel,
            IPersistenceMapper<BlogPostLikeEntity, BlogPostLikeModel> likeMapperToModel,
            IPersistenceMapper<TagEntity, TagModel> tagMapperToModel)
        {
            _tagMapperToModel = tagMapperToModel;
            _likeMapperToModel = likeMapperToModel;
            _commentMapperToModel = commentMapperToModel;
            _tagMapperToEntity = tagMapperToEntity;
            _likeMapperToEntity = likeMapperToEntity;
            _commentMapperToEntity = commentMapperToEntity;
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
            Tags = entity.Tags.Select(_tagMapperToEntity.MapToModel).ToArray(),
            Likes = entity.Likes.Select(_likeMapperToEntity.MapToModel).ToArray(),
            Comments = entity.Comments.Select(_commentMapperToEntity.MapToModel).ToArray()
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
            Tags = model.Tags.Select(_tagMapperToModel.MapToEntity).ToArray(), 
            Likes = model.Likes.Select(_likeMapperToModel.MapToEntity).ToArray(),
            Comments = model.Comments.Select(_commentMapperToModel.MapToEntity).ToArray()
        };
    }
}