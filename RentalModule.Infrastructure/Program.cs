using Microsoft.EntityFrameworkCore;
using RentalModule.Infrastructure.DataBase.Context;

var services = new ServiceCollection();

IConfiguration configuration = new ConfigurationBuilder()
    .AddJsonFile("appsettings.json", optional: false, reloadOnChange: true)
    .Build();

var connectionString = configuration.GetConnectionString("CarRentalBlogDbConnectionString");
services.AddDbContext<RentalDbContext>(options => options.UseSqlServer(connectionString));  

services.BuildServiceProvider();  

Console.WriteLine("This is a migration project for TrucksModule.Infrastructure.");