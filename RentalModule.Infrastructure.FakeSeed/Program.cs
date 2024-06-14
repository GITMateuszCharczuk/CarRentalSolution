using Bogus;
using Microsoft.EntityFrameworkCore;
using RentalModule.Infrastructure.DataBase.Context;
using RentalModule.Infrastructure.DataBase.Entities;
using RentalModule.Infrastructure.FakeSeed.Fakers;

IConfiguration configuration = new ConfigurationBuilder()
    .AddJsonFile("appsettings.json", optional: false, reloadOnChange: true)
    .Build();

var connectionString = configuration.GetConnectionString("CarRentalBlogDbConnectionString");
var dbContextOptions = new DbContextOptionsBuilder<RentalDbContext>().UseSqlServer(connectionString).Options;
var blogDbContext = new RentalDbContext(dbContextOptions);

var fakeOffers = CarOfferFaker.Generate(100);
blogDbContext.CarOffers.AddRange(fakeOffers);

var count = await blogDbContext.SaveChangesAsync();

Console.WriteLine($"Number of fake Car offers added {count}");