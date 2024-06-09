using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using Shared.CQRS;

namespace BlogModule.Infrastructure.DataBase.Entities;

public record BlogPostEntity : BlogPostModel, IEntity<BlogPostId>
{
    
}