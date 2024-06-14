using RentalModule.Domain.Models;
using RentalModule.Infrastructure.DataBase.Entities;
using Shared.Utilities;

namespace RentalModule.Infrastructure.Mappers
{
    public class CarTagPersistenceMapper : IPersistenceMapper<CarTagEntity, CarTagModel>
    {
        public CarTagModel MapToModel(CarTagEntity entity) => new()
        {
            Id = entity.Id,
            Name = entity.Name,
            CarOfferId = entity.CarOfferId
        };

        public CarTagEntity MapToEntity(CarTagModel model) => new()
        {
            Id = model.Id,
            Name = model.Name,
            CarOfferId = model.CarOfferId
        };
    }
}