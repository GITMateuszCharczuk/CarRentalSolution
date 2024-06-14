using Bogus;
using RentalModule.Domain.Models.Ids;
using RentalModule.Infrastructure.DataBase.Entities;

namespace RentalModule.Infrastructure.FakeSeed.Fakers;

public static class CarTariffFaker
{
    private static Faker<CarTariffEntity>? _faker;

    public static Faker<CarTariffEntity> GetFaker(Guid carOfferId)
    {
        return _faker ??= Create(carOfferId);
    }

    private static Faker<CarTariffEntity> Create(Guid carOfferId) =>
        new Faker<CarTariffEntity>()
            .RuleFor(entity => entity.Id, faker => new CarTariffId(Guid.NewGuid()))
            .RuleFor(entity => entity.CarOfferId, faker => new CarOfferId(carOfferId))
            .RuleFor(entity => entity.OneNormalDayPrice, faker => faker.Random.Double(50, 200))
            .RuleFor(entity => entity.OneWeekendDayPrice, faker => faker.Random.Double(60, 250))
            .RuleFor(entity => entity.FullWeekendPrice, faker => faker.Random.Double(100, 400))
            .RuleFor(entity => entity.OneWeekPrice, faker => faker.Random.Double(300, 1500))
            .RuleFor(entity => entity.OneMonthPrice, faker => faker.Random.Double(1000, 4000));
    

    public static IEnumerable<CarTariffEntity> Generate(int count, Guid carOfferId) => GetFaker(carOfferId).Generate(count);
}