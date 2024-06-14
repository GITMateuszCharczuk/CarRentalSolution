using Bogus;
using RentalModule.Domain.Models.Ids;
using RentalModule.Infrastructure.DataBase.Entities;

namespace RentalModule.Infrastructure.FakeSeed.Fakers;

public static class CarOfferFaker
{
    private static Faker<CarOfferEntity>? _faker;

    public static Faker<CarOfferEntity> Faker
    {
        get { return _faker ??= Create(); }
    }

    private static Faker<CarOfferEntity> Create() =>
        new Faker<CarOfferEntity>()
            .RuleFor(entity => entity.Id, faker => new CarOfferId(Guid.NewGuid()))
            .RuleFor(entity => entity.Heading, faker => faker.Lorem.Sentence())
            .RuleFor(entity => entity.ShortDescription, faker => faker.Lorem.Sentence())
            .RuleFor(entity => entity.FeaturedImageUrl, faker => faker.Image.PicsumUrl(width: 1920, height: 1080))
            .RuleFor(entity => entity.UrlHandle, faker => faker.Lorem.Slug())
            .RuleFor(entity => entity.Horsepower, faker => faker.Random.Int(100, 500).ToString())
            .RuleFor(entity => entity.YearOfProduction, faker => faker.Date.Past(20).Year)
            .RuleFor(entity => entity.EngineDetails, faker => faker.Lorem.Sentence())
            .RuleFor(entity => entity.DriveDetails, faker => faker.Lorem.Sentence())
            .RuleFor(entity => entity.GearboxDetails, faker => faker.Lorem.Sentence())
            .RuleFor(entity => entity.CarDeliverylocation, faker => faker.Address.FullAddress())
            .RuleFor(entity => entity.CarReturnLocation, faker => faker.Address.FullAddress())
            .RuleFor(entity => entity.PublishedDate, faker => faker.Date.Past())
            .RuleFor(entity => entity.Visible, faker => faker.Random.Bool())
            .RuleFor(entity => entity.Tarrif, (faker, entity) => CarTariffFaker.Generate(1, entity.Id).First())
            .RuleFor(entity => entity.Tags, (faker, entity) => CarTagFaker.Generate(faker.Random.Int(2, 8), entity.Id).ToList())
            .RuleFor(entity => entity.ImageUrls, (faker, entity) => ImageUrlFaker.Generate(faker.Random.Int(1, 10), entity.Id).ToList())
            .RuleFor(entity => entity.UnavailableDates, (faker, entity) => TimePeriodFaker.Generate(faker.Random.Int(1, 5), entity.Id).ToList());

    public static IEnumerable<CarOfferEntity> Generate(int count) => Faker.Generate(count);
}