using BlogModule.Domain.Models.Ids;
using CarRental.Web.Models.Domain.Blog;
using Shared.CQRS;

namespace BlogModule.Infrastructure.DataBase.Entities;

public record TagEntity : TagModel, IEntity<TagId>
{
}