namespace BlogModule.Infrastructure.FakeSeed.Fakers;

public static class BlogFaker
{
    private static Faker<TruckEntity>? _faker;

    public static Faker<TruckEntity> Faker
    {
        get { return _faker ??= Create(); }
    }

    private static Faker<TruckEntity> Create() =>
        new Faker<TruckEntity>()
            .RuleFor(entity => entity.Code, faker => faker.Vehicle.Vin())
            .RuleFor(entity => entity.Name, faker => $"{faker.Vehicle.Manufacturer()} {faker.Vehicle.Model()}")
            .RuleFor(entity => entity.Description, faker => faker.Lorem.Sentence().OrNull(faker, 0.7f))
            .RuleFor(entity => entity.Status, faker => faker.PickRandom<TruckStatusEnum>());
    
    public static IEnumerable<TruckEntity> Generate(int count) => Faker.Generate(count);
            

}