using RentalModule.Domain.Models;
using RentalModule.Domain.Models.Ids;
using RentalModule.Domain.RepositoryInterfaces.CarOrder;
using RentalModule.Infrastructure.DataBase.Context;
using RentalModule.Infrastructure.DataBase.Entities;
using Shared.CQRS.Repository;
using Shared.Utilities;

namespace RentalModule.Infrastructure.DataBase.Repository.CarOrder;

public class CarOrderCommandRepository : CommandRepository<CarOrderEntity, CarOrderId, CarOrderModel, RentalDbContext>, ICarOrderCommandRepository
{
    public CarOrderCommandRepository(RentalDbContext dbContext, IPersistenceMapper<CarOrderEntity, CarOrderModel> mapper) : base(dbContext, mapper)
    {
    }
}