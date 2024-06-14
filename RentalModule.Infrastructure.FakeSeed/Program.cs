using Bogus;
using Microsoft.EntityFrameworkCore;
using RentalModule.Infrastructure.DataBase.Context;
using RentalModule.Infrastructure.DataBase.Entities;
using RentalModule.Infrastructure.FakeSeed.Fakers;

IConfiguration configuration = new ConfigurationBuilder()
    .AddJsonFile("appsettings.json", optional: false, reloadOnChange: true)
    .Build();

var connectionString = configuration.GetConnectionString("CarRentalCarDbConnectionString");
var dbContextOptions = new DbContextOptionsBuilder<RentalDbContext>().UseSqlServer(connectionString).Options;
var carDbContext = new RentalDbContext(dbContextOptions);

var fakeOffers = CarOfferFaker.Generate(100);
carDbContext.CarOffers.AddRange(fakeOffers);
var fakeOrders = CarOrderFaker.Generate(100,fakeOffers);
carDbContext.CarOrders.AddRange(fakeOrders);

var count = await carDbContext.SaveChangesAsync();

Console.WriteLine($"Number of fake Entities offers added {count}");