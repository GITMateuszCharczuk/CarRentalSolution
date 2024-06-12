using BlogModule.Infrastructure.DataBase.Context;
using BlogModule.Infrastructure.FakeSeed.Fakers;
using Microsoft.EntityFrameworkCore;

IConfiguration configuration = new ConfigurationBuilder()
    .AddJsonFile("appsettings.json", optional: false, reloadOnChange: true)
    .Build();

var connectionString = configuration.GetConnectionString("CarRentalBlogDbConnectionString");
var dbContextOptions = new DbContextOptionsBuilder<BlogDbContext>().UseSqlServer(connectionString).Options;
var blogDbContext = new BlogDbContext(dbContextOptions);

var fakeBlogs = BlogPostFaker.Generate(100);
blogDbContext.BlogPosts.AddRange(fakeBlogs);

var count = await blogDbContext.SaveChangesAsync();

Console.WriteLine($"Number of fake Blog posts added {count}");