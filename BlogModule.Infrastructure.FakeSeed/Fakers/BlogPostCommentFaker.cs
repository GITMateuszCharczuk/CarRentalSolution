using Bogus;
using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using BlogModule.Infrastructure.DataBase.Entities;

namespace BlogModule.Infrastructure.FakeSeed.Fakers
{
    public static class BlogPostCommentFaker
    {
        
        private static Faker<BlogPostCommentEntity>? _faker;

        public static Faker<BlogPostCommentEntity> GetFaker(Guid blogPostId)
        {
            return _faker ??= Create(blogPostId);
        }

        private static Faker<BlogPostCommentEntity> Create(Guid blogPostId) =>
            new Faker<BlogPostCommentEntity>()
                .RuleFor(entity => entity.Id, faker => new BlogPostCommentId(Guid.NewGuid()))
                .RuleFor(entity => entity.Description, faker => faker.Lorem.Sentence())
                .RuleFor(entity => entity.BlogPostId, faker => new BlogPostId(blogPostId))
                .RuleFor(entity => entity.UserId, faker => Guid.NewGuid())
                .RuleFor(entity => entity.DateAdded, faker => faker.Date.Past());
        
        public static IEnumerable<BlogPostCommentEntity> Generate(int count, Guid blogPostId) => GetFaker(blogPostId).Generate(count);
        
    }
}