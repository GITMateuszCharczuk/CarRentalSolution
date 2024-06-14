using RentalModule.Domain.Models;
using RentalModule.Domain.Models.Ids;
using Shared.CQRS;

namespace RentalModule.Infrastructure.DataBase.Entities;

public record CarTagEntity : CarTagModel, IEntity<CarTagId>
{
    
}
