using BlogModule.Infrastructure.DataBase.Context;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Configuration;

var services = new ServiceCollection();

IConfiguration configuration = new ConfigurationBuilder()
    .AddJsonFile("appsettings.json", optional: false, reloadOnChange: true)
    .Build();

var connectionString = configuration.GetConnectionString("CarRentalBlogDbConnectionString");
services.AddDbContext<BlogDbContext>(options => options.UseSqlServer(connectionString));  

services.BuildServiceProvider();  

Console.WriteLine("This is a migration project for TrucksModule.Infrastructure.");