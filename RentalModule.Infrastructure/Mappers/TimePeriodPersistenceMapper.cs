using RentalModule.Domain.Models;
using RentalModule.Infrastructure.DataBase.Entities;
using Shared.Utilities;

namespace RentalModule.Infrastructure.Mappers
{
    public class TimePeriodPersistenceMapper : IPersistenceMapper<TimePeriodEntity, TimePeriodModel>
    {
        public TimePeriodModel MapToModel(TimePeriodEntity entity) => new()
        {
            Id = entity.Id,
            CarOfferId = entity.CarOfferId,
            StartDate = entity.StartDate,
            EndDate = entity.EndDate
        };

        public TimePeriodEntity MapToEntity(TimePeriodModel model) => new()
        {
            Id = model.Id,
            CarOfferId = model.CarOfferId,
            StartDate = model.StartDate,
            EndDate = model.EndDate
        };
    }
}