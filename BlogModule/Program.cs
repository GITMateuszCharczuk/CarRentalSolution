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
using RentalModule.API.Mappers;
using RentalModule.Application.QueryHandlers.CarOffer.GetCarOffers;
using RentalModule.Domain.Models;
using RentalModule.Domain.RepositoryInterfaces.CarOffer;
using RentalModule.Domain.RepositoryInterfaces.CarOrder;
using RentalModule.Domain.RepositoryInterfaces.CarTag;
using RentalModule.Infrastructure.Binders.CarOfferId;
using RentalModule.Infrastructure.DataBase.Context;
using RentalModule.Infrastructure.DataBase.Entities;
using RentalModule.Infrastructure.DataBase.Repository.CarOffer;
using RentalModule.Infrastructure.DataBase.Repository.CarOrder;
using RentalModule.Infrastructure.DataBase.Repository.CarTag;
using RentalModule.Infrastructure.DataBase.UnitOfWork;
using RentalModule.Infrastructure.Decorators;
using RentalModule.Infrastructure.Mappers;
using Shared.Behaviors;
using Shared.Utilities;

var builder = WebApplication.CreateBuilder(args);

var connectionStringBLog = builder.Configuration.GetConnectionString("CarRentalBlogDbConnectionString");
ArgumentNullException.ThrowIfNull(connectionStringBLog, nameof(connectionStringBLog));
builder.Services.AddDbContext<BlogDbContext>(options => options.UseSqlServer(connectionStringBLog));

var connectionStringCar = builder.Configuration.GetConnectionString("CarRentalCarDbConnectionString");
ArgumentNullException.ThrowIfNull(connectionStringCar, nameof(connectionStringCar));
builder.Services.AddDbContext<RentalDbContext>(options => options.UseSqlServer(connectionStringCar));

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


builder.Services.AddScoped<IRentalUnitOfWork, RentalUnitOfWork>();

// CarOffer Repositories
builder.Services.AddScoped<ICarOfferCommandRepository, CarOfferCommandRepository>();
builder.Services.AddScoped<ICarOfferQueryRepository, CarOfferQueryRepository>();
builder.Services.AddScoped<ICarTagCommandRepository, CarTagCommandRepository>();
builder.Services.AddScoped<ICarTagQueryRepository, CarTagQueryRepository>();
builder.Services.AddScoped<ICarOrderCommandRepository, CarOrderCommandRepository>();
builder.Services.AddScoped<ICarOrderQueryRepository, CarOrderQueryRepository>();

// API Mappers
builder.Services.AddScoped<ICarApiMapper, CarApiMapper>();

// Persistence Mappers
builder.Services.AddScoped<IPersistenceMapper<CarOfferEntity, CarOfferModel>, CarOfferPersistenceMapper>();
builder.Services.AddScoped<IPersistenceMapper<CarTagEntity, CarTagModel>, CarTagPersistenceMapper>();
builder.Services.AddScoped<IPersistenceMapper<CarOrderEntity, CarOrderModel>, CarOrderPersistenceMapper>();
builder.Services.AddScoped<IPersistenceMapper<ImageUrlEntity, ImageUrlModel>, ImageUrlPersistenceMapper>();
builder.Services.AddScoped<IPersistenceMapper<CarTariffEntity, CarTariffModel>, CarTariffPersistenceMapper>();
builder.Services.AddScoped<IPersistenceMapper<TimePeriodEntity, TimePeriodModel>, TimePeriodPersistenceMapper>();

builder.Services.AddMediatR(cfg =>
{
    cfg.RegisterServicesFromAssembly(typeof(GetBlogPostsQuery).Assembly);
    cfg.RegisterServicesFromAssembly(typeof(GetCarOffersQuery).Assembly);
});
builder.Services.AddScoped(typeof(IPipelineBehavior<,>), typeof(ValidationHandlerBehaviour<,>));
builder.Services.AddScoped(typeof(IPipelineBehavior<,>), typeof(BlogCommandHandlerBehavior<,>));
builder.Services.AddScoped(typeof(IPipelineBehavior<,>), typeof(RentalCommandHandlerBehavior<,>));

builder.Services.AddControllers(options =>
    {
        options.ModelBinderProviders.Insert(0, new BlogPostIdModelBinderProvider());
        options.ModelBinderProviders.Insert(0, new TagIdModelBinderProvider());
        options.ModelBinderProviders.Insert(0, new BlogPostCommentIdModelBinderProvider());
        options.ModelBinderProviders.Insert(0, new BlogPostLikeIdModelBinderProvider());
        
        options.ModelBinderProviders.Insert(0, new CarOfferIdModelBinderProvider());
    })
    .AddJsonOptions(options => options.JsonSerializerOptions.Converters.Add(new JsonStringEnumConverter()));


builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

app.UseSwagger();
app.UseSwaggerUI();

app.MapControllers();
app.MapGet("/", () => "API");

app.Run();