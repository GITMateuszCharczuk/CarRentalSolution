using System.Reflection;
using System.Text.Json.Serialization;
using Blog.API.Mappers;
using BlogModule.API.Mappers;
using BlogModule.Application.QueryHandlers.BlogPost.GetBlogPosts;
using BlogModule.Domain.Models;
using BlogModule.Domain.RepositoryInterfaces.BlogPost;
using BlogModule.Domain.RepositoryInterfaces.BlogPostComment;
using BlogModule.Domain.RepositoryInterfaces.BlogPostLike;
using BlogModule.Domain.RepositoryInterfaces.Tag;
using BlogModule.Infrastructure.Binders.BlogPostCommentId;
using BlogModule.Infrastructure.Binders.BlogPostId;
using BlogModule.Infrastructure.Binders.BlogPostLikeId;
using BlogModule.Infrastructure.Binders.TagId;
using BlogModule.Infrastructure.DataBase.Context;
using BlogModule.Infrastructure.DataBase.Entities;
using BlogModule.Infrastructure.DataBase.Repository.BlogPost;
using BlogModule.Infrastructure.DataBase.Repository.BlogPostComment;
using BlogModule.Infrastructure.DataBase.Repository.BlogPostLike;
using BlogModule.Infrastructure.DataBase.Repository.Tag;
using BlogModule.Infrastructure.DataBase.UnitOfWork;
using BlogModule.Infrastructure.Decorators;
using BlogModule.Infrastructure.Mappers;
using MediatR;
using Microsoft.EntityFrameworkCore;
using Shared.Behaviors;
using Shared.Utilities;

var builder = WebApplication.CreateBuilder(args);

var connectionString = builder.Configuration.GetConnectionString("CarRentalBlogDbConnectionString");
ArgumentNullException.ThrowIfNull(connectionString, nameof(connectionString));
builder.Services.AddDbContext<BlogDbContext>(options => options.UseSqlServer(connectionString));

builder.Services.AddScoped<IBlogUnitOfWork, BlogUnitOfWork>();
builder.Services.AddScoped<IBlogPostLikeCommandRepository, BlogPostLikeCommandRepository>();
builder.Services.AddScoped<IBlogPostLikeQueryRepository, BlogPostLikeQueryRepository>();
builder.Services.AddScoped<IBlogPostCommentCommandRepository,BlogPostCommentCommandRepository>();
builder.Services.AddScoped<IBlogPostCommentQueryRepository, BlogPostCommentQueryRepository>();
builder.Services.AddScoped<IBlogPostCommandRepository,BlogPostCommandRepository>();
builder.Services.AddScoped<IBlogPostQueryRepository,BlogPostQueryRepository>();
builder.Services.AddScoped<ITagCommandRepository,TagCommandRepository>();
builder.Services.AddScoped<ITagQueryRepository,TagQueryRepository>();

builder.Services.AddScoped<IBlogPostApiMapper, BlogPostApiMapper>();
builder.Services.AddScoped<IPersistenceMapper<BlogPostLikeEntity, BlogPostLikeModel>, BlogPostLikePersistenceMapper>();
builder.Services.AddScoped<IPersistenceMapper<BlogPostCommentEntity, BlogPostCommentModel>, BlogPostCommentPersistenceMapper>();
builder.Services.AddScoped<IPersistenceMapper<TagEntity, TagModel>, TagPersistenceMapper>();
builder.Services.AddScoped<IPersistenceMapper<BlogPostEntity, BlogPostModel>, BlogPostPersistenceMapper>();


builder.Services.AddMediatR(cfg =>
{
    cfg.RegisterServicesFromAssembly(typeof(GetBlogPostsQuery).Assembly);
});
builder.Services.AddScoped(typeof(IPipelineBehavior<,>), typeof(ValidationHandlerBehaviour<,>));
builder.Services.AddScoped(typeof(IPipelineBehavior<,>), typeof(BlogCommandHandlerBehavior<,>));

builder.Services.AddControllers(options =>
    {
        options.ModelBinderProviders.Insert(0, new BlogPostIdModelBinderProvider());
        options.ModelBinderProviders.Insert(0, new TagIdModelBinderProvider());
        options.ModelBinderProviders.Insert(0, new BlogPostCommentIdModelBinderProvider());
        options.ModelBinderProviders.Insert(0, new BlogPostLikeIdModelBinderProvider());
    })
    .AddJsonOptions(options => options.JsonSerializerOptions.Converters.Add(new JsonStringEnumConverter()));


builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

app.UseSwagger();
app.UseSwaggerUI();

app.MapControllers();
app.MapGet("/", () => "BLOG API");

app.Run();