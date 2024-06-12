using Bogus;
using BlogModule.Domain.Models.Ids;
using BlogModule.Infrastructure.DataBase.Entities;

namespace BlogModule.Infrastructure.FakeSeed.Fakers
{
    public static class BlogPostFaker
    {
        private static Faker<BlogPostEntity>? _faker;

        public static Faker<BlogPostEntity> Faker
        {
            get { return _faker ??= Create(); }
        }

        private static Faker<BlogPostEntity> Create() =>
            new Faker<BlogPostEntity>()
                .RuleFor(entity => entity.Id, faker => new BlogPostId(Guid.NewGuid()))
                .RuleFor(entity => entity.Heading, faker => faker.Lorem.Sentence())
                .RuleFor(entity => entity.PageTitle, faker => faker.Lorem.Sentence())
                .RuleFor(entity => entity.Content, faker => faker.Lorem.Paragraphs())
                .RuleFor(entity => entity.ShortDescription, faker => faker.Lorem.Sentence())
                .RuleFor(entity => entity.FeaturedImageUrl, faker => faker.Image.PicsumUrl(width: 1920, height:1080))
                .RuleFor(entity => entity.UrlHandle, faker => faker.Lorem.Slug())
                .RuleFor(entity => entity.PublishedDate, faker => faker.Date.Past())
                .RuleFor(entity => entity.Author, faker => faker.Person.FullName)
                .RuleFor(entity => entity.Visible, faker => faker.Random.Bool())
                .RuleFor(entity => entity.Tags, (faker, entity) => TagFaker.Generate(faker.Random.Int(2, 8), entity.Id.Value).ToList())
                .RuleFor(entity => entity.Likes, (faker, entity) => BlogPostLikeFaker.Generate(faker.Random.Int(5, 60), entity.Id.Value).ToList())
                .RuleFor(entity => entity.Comments, (faker, entity) => BlogPostCommentFaker.Generate(faker.Random.Int(1, 8), entity.Id.Value).ToList());

        public static IEnumerable<BlogPostEntity> Generate(int count) => Faker.Generate(count);
    }
}