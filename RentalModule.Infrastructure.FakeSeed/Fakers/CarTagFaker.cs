using Bogus;
using RentalModule.Domain.Models.Ids;
using RentalModule.Infrastructure.DataBase.Entities;

namespace RentalModule.Infrastructure.FakeSeed.Fakers;

public static class CarTagFaker
{
    private static Faker<CarTagEntity>? _faker;

    public static Faker<CarTagEntity> GetFaker(Guid carOfferId)
    {
        return _faker ??= Create(carOfferId);
    }

    private static Faker<CarTagEntity> Create(Guid carOfferId) =>
        new Faker<CarTagEntity>()
            .RuleFor(entity => entity.Id, faker => new CarTagId(Guid.NewGuid()))
            .RuleFor(entity => entity.Name, faker => faker.Vehicle.Type())
            .RuleFor(entity => entity.CarOfferId, faker => new CarOfferId(carOfferId));

    public static IEnumerable<CarTagEntity> Generate(int count, Guid carOfferId) => GetFaker(carOfferId).Generate(count);
}