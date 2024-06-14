using RentalModule.Domain.Models;
using RentalModule.Infrastructure.DataBase.Entities;
using Shared.Utilities;

namespace RentalModule.Infrastructure.Mappers;

public class CarTariffPersistenceMapper : IPersistenceMapper<CarTariffEntity, CarTariffModel>
{
    public CarTariffModel MapToModel(CarTariffEntity entity)
    {
        if (entity == null) throw new ArgumentNullException(nameof(entity));

        return new CarTariffModel
        {
            Id = entity.Id,
            CarOfferId = entity.CarOfferId,
            OneNormalDayPrice = entity.OneNormalDayPrice,
            OneWeekendDayPrice = entity.OneWeekendDayPrice,
            FullWeekendPrice = entity.FullWeekendPrice,
            OneWeekPrice = entity.OneWeekPrice,
            OneMonthPrice = entity.OneMonthPrice
        };
    }

    public CarTariffEntity MapToEntity(CarTariffModel model)
    {
        if (model == null) throw new ArgumentNullException(nameof(model));

        return new CarTariffEntity
        {
            Id = model.Id,
            CarOfferId = model.CarOfferId,
            OneNormalDayPrice = model.OneNormalDayPrice,
            OneWeekendDayPrice = model.OneWeekendDayPrice,
            FullWeekendPrice = model.FullWeekendPrice,
            OneWeekPrice = model.OneWeekPrice,
            OneMonthPrice = model.OneMonthPrice
        };
    }
}