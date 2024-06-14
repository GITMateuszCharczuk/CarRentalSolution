using Bogus;
using RentalModule.Domain.Models.Ids;
using RentalModule.Infrastructure.DataBase.Entities;

namespace RentalModule.Infrastructure.FakeSeed.Fakers;

public static class TimePeriodFaker
{
    private static Faker<TimePeriodEntity>? _faker;

    public static Faker<TimePeriodEntity> GetFaker(Guid carOfferId)
    {
        return _faker ??= Create(carOfferId);
    }

    private static Faker<TimePeriodEntity> Create(Guid carOfferId) =>
        new Faker<TimePeriodEntity>()
            .RuleFor(entity => entity.Id, faker => new TimePeriodId(Guid.NewGuid()))
            .RuleFor(entity => entity.CarOfferId, faker => new CarOfferId(carOfferId))
            .RuleFor(entity => entity.StartDate, faker => faker.Date.Future())
            .RuleFor(entity => entity.EndDate, (faker, entity) => entity.StartDate.AddDays(faker.Random.Int(1, 5)));

    public static IEnumerable<TimePeriodEntity> Generate(int count, Guid carOfferId) => GetFaker(carOfferId).Generate(count);
}