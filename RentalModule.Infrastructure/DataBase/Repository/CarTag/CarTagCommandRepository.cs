using RentalModule.Domain.Models;
using RentalModule.Domain.Models.Ids;
using RentalModule.Domain.RepositoryInterfaces.CarTag;
using RentalModule.Infrastructure.DataBase.Context;
using RentalModule.Infrastructure.DataBase.Entities;
using Shared.CQRS.Repository;
using Shared.Utilities;

namespace RentalModule.Infrastructure.DataBase.Repository.CarTag;

public class CarTagCommandRepository :  CommandRepository<CarTagEntity, CarTagId, CarTagModel, RentalDbContext>, ICarTagCommandRepository
{
    public CarTagCommandRepository(RentalDbContext dbContext, IPersistenceMapper<CarTagEntity, CarTagModel> mapper) : base(dbContext, mapper)
    {
    }
}