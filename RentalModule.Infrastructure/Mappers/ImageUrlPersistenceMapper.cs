using RentalModule.Domain.Models;
using RentalModule.Infrastructure.DataBase.Entities;
using Shared.Utilities;

namespace RentalModule.Infrastructure.Mappers
{
    public class ImageUrlPersistenceMapper : IPersistenceMapper<ImageUrlEntity, ImageUrlModel>
    {
        public ImageUrlModel MapToModel(ImageUrlEntity entity) => new()
        {
            Id = entity.Id,
            CarOfferId = entity.CarOfferId,
            Url = entity.Url
        };

        public ImageUrlEntity MapToEntity(ImageUrlModel model) => new()
        {
            Id = model.Id,
            CarOfferId = model.CarOfferId,
            Url = model.Url
        };
    }
}