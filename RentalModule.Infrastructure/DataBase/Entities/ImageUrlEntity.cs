using RentalModule.Domain.Models;
using RentalModule.Domain.Models.Ids;
using Shared.CQRS;

namespace RentalModule.Infrastructure.DataBase.Entities;

public record ImageUrlEntity : ImageUrlModel, IEntity<ImageUrlId>
{
}