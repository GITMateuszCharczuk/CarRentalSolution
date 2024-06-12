using Bogus;
using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using BlogModule.Infrastructure.DataBase.Entities;

namespace BlogModule.Infrastructure.FakeSeed.Fakers
{
    public static class BlogPostLikeFaker
    {
        private static Faker<BlogPostLikeEntity>? _faker;

        public static Faker<BlogPostLikeEntity> GetFaker(Guid blogPostId)
        {
            return _faker ??= Create(blogPostId);
        }

        private static Faker<BlogPostLikeEntity> Create(Guid blogPostId) =>
            new Faker<BlogPostLikeEntity>()
                .RuleFor(entity => entity.Id, faker => new BlogPostLikeId(Guid.NewGuid()))
                .RuleFor(entity => entity.BlogPostId, faker => new BlogPostId(blogPostId))
                .RuleFor(entity => entity.UserId, faker => Guid.NewGuid());

        public static IEnumerable<BlogPostLikeEntity> Generate(int count, Guid blogPostId) => GetFaker(blogPostId).Generate(count);
    }
}