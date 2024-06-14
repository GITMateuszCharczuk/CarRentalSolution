using Bogus;
using RentalModule.Domain.Models.Ids;
using RentalModule.Infrastructure.DataBase.Entities;

namespace RentalModule.Infrastructure.FakeSeed.Fakers;

public static class ImageUrlFaker
{
    private static Faker<ImageUrlEntity>? _faker;

    public static Faker<ImageUrlEntity> GetFaker(Guid carOfferId)
    {
        return _faker ??= Create(carOfferId);
    }

    private static Faker<ImageUrlEntity> Create(Guid carOfferId) =>
        new Faker<ImageUrlEntity>()
            .RuleFor(entity => entity.Id, faker => new ImageUrlId(Guid.NewGuid()))
            .RuleFor(entity => entity.CarOfferId, faker => new CarOfferId(carOfferId))
            .RuleFor(entity => entity.Url, faker => faker.Image.PicsumUrl(width: 1920, height:1080));

    public static IEnumerable<ImageUrlEntity> Generate(int count, Guid carOfferId) => GetFaker(carOfferId).Generate(count);
}