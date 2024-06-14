using Bogus;
using RentalModule.Domain.Models.Ids;
using RentalModule.Infrastructure.DataBase.Entities;

namespace RentalModule.Infrastructure.FakeSeed.Fakers;

public static class CarOrderFaker
{
    private static Faker<CarOrderEntity>? _faker;

    public static Faker<CarOrderEntity> GetFaker(IEnumerable<CarOfferEntity> carOfferIds)
    {
        return _faker ??= Create(carOfferIds);
    }

    private static Faker<CarOrderEntity> Create(IEnumerable<CarOfferEntity> carOfferIds) =>
        new Faker<CarOrderEntity>()
            .RuleFor(entity => entity.Id, faker => new CarOrderId(Guid.NewGuid()))
            .RuleFor(entity => entity.UserId, faker => Guid.NewGuid())
            .RuleFor(entity => entity.CarOfferId, faker => new CarOfferId(faker.PickRandom(carOfferIds).Id))
            .RuleFor(entity => entity.StartDate, faker => faker.Date.Future())
            .RuleFor(entity => entity.EndDate, (faker, entity) => entity.StartDate.AddDays(faker.Random.Int(1, 30)))
            .RuleFor(entity => entity.Notes, faker => faker.Lorem.Paragraph())
            .RuleFor(entity => entity.NumOfDrivers, faker => faker.Random.Int(1, 4))
            .RuleFor(entity => entity.TotalCost, faker => faker.Random.Double(500, 10000));

    public static IEnumerable<CarOrderEntity> Generate(int count, IEnumerable<CarOfferEntity> carOfferIds) => GetFaker(carOfferIds).Generate(count);
}