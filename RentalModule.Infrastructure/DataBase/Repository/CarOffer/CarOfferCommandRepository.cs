using RentalModule.Domain.Models;
using RentalModule.Domain.Models.Ids;
using RentalModule.Domain.RepositoryInterfaces.CarOffer;
using RentalModule.Infrastructure.DataBase.Context;
using RentalModule.Infrastructure.DataBase.Entities;
using Shared.CQRS.Repository;
using Shared.Utilities;

namespace RentalModule.Infrastructure.DataBase.Repository.CarOffer;

public class CarOfferCommandRepository : CommandRepository<CarOfferEntity, CarOfferId, CarOfferModel, RentalDbContext>, ICarOfferCommandRepository
{
    public CarOfferCommandRepository(RentalDbContext dbContext, IPersistenceMapper<CarOfferEntity, CarOfferModel> mapper) : base(dbContext, mapper)
    {
    }
}