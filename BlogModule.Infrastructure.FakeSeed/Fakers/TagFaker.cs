using Bogus;
using BlogModule.Domain.Models.Ids;
using BlogModule.Infrastructure.DataBase.Entities;

namespace BlogModule.Infrastructure.FakeSeed.Fakers
{
    public static class TagFaker
    {
        private static Faker<TagEntity>? _faker;

        public static Faker<TagEntity> GetFaker(Guid blogPostId)
        {
            return _faker ??= Create(blogPostId);
        }

        private static Faker<TagEntity> Create(Guid blogPostId) =>
            new Faker<TagEntity>()
                .RuleFor(entity => entity.Id, faker => new TagId(Guid.NewGuid()))
                .RuleFor(entity => entity.Name, faker => faker.Lorem.Word())
                .RuleFor(entity => entity.BlogPostId, faker => new BlogPostId(blogPostId));


        public static IEnumerable<TagEntity> Generate(int count, Guid blogPostId) => GetFaker(blogPostId).Generate(count);
    }
}