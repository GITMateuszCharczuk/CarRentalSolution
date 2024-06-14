using RentalModule.Domain.Models;
using RentalModule.Infrastructure.DataBase.Entities;
using Shared.Utilities;

namespace RentalModule.Infrastructure.Mappers
{
    public class CarOrderPersistenceMapper : IPersistenceMapper<CarOrderEntity, CarOrderModel>
    {
        public CarOrderModel MapToModel(CarOrderEntity entity) => new()
        {
            Id = entity.Id,
            UserId = entity.UserId,
            CarOfferId = entity.CarOfferId,
            StartDate = entity.StartDate,
            EndDate = entity.EndDate,
            Notes = entity.Notes,
            NumOfDrivers = entity.NumOfDrivers,
            TotalCost = entity.TotalCost
        };

        public CarOrderEntity MapToEntity(CarOrderModel model) => new()
        {
            Id = model.Id,
            UserId = model.UserId,
            CarOfferId = model.CarOfferId,
            StartDate = model.StartDate,
            EndDate = model.EndDate,
            Notes = model.Notes,
            NumOfDrivers = model.NumOfDrivers,
            TotalCost = model.TotalCost
        };
    }
}