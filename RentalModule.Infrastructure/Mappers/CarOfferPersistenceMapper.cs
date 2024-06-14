using RentalModule.Domain.Models;
using RentalModule.Infrastructure.DataBase.Entities;
using Shared.Utilities;

namespace RentalModule.Infrastructure.Mappers
{
    public class CarOfferPersistenceMapper : IPersistenceMapper<CarOfferEntity, CarOfferModel>
    {
        private readonly IPersistenceMapper<CarTagEntity, CarTagModel> _tagMapper;
        private readonly IPersistenceMapper<ImageUrlEntity, ImageUrlModel> _imageUrlMapper;
        private readonly IPersistenceMapper<TimePeriodEntity, TimePeriodModel> _timePeriodMapper;
        private readonly IPersistenceMapper<CarTariffEntity, CarTariffModel> _tariffMapper;

        public CarOfferPersistenceMapper(
            IPersistenceMapper<CarTagEntity, CarTagModel> tagMapper,
            IPersistenceMapper<ImageUrlEntity, ImageUrlModel> imageUrlMapper,
            IPersistenceMapper<TimePeriodEntity, TimePeriodModel> timePeriodMapper,
            IPersistenceMapper<CarTariffEntity, CarTariffModel> tariffMapper)
        {
            _tagMapper = tagMapper;
            _imageUrlMapper = imageUrlMapper;
            _timePeriodMapper = timePeriodMapper;
            _tariffMapper = tariffMapper;
        }

        public CarOfferModel MapToModel(CarOfferEntity entity) => new()
        {
            Id = entity.Id,
            Heading = entity.Heading,
            ShortDescription = entity.ShortDescription,
            FeaturedImageUrl = entity.FeaturedImageUrl,
            UrlHandle = entity.UrlHandle,
            Horsepower = entity.Horsepower,
            YearOfProduction = entity.YearOfProduction,
            EngineDetails = entity.EngineDetails,
            DriveDetails = entity.DriveDetails,
            GearboxDetails = entity.GearboxDetails,
            CarDeliverylocation = entity.CarDeliverylocation,
            CarReturnLocation = entity.CarReturnLocation,
            PublishedDate = entity.PublishedDate,
            Visible = entity.Visible,
            Tarrif = _tariffMapper.MapToModel(entity.Tarrif),
            Tags = entity.Tags?.Select(tag => _tagMapper.MapToModel(tag)).ToArray() ?? Array.Empty<CarTagModel>(),
            ImageUrls = entity.ImageUrls?.Select(image => _imageUrlMapper.MapToModel(image)).ToList() ?? new List<ImageUrlModel>(),
            UnavailableDates = entity.UnavailableDates?.Select(period => _timePeriodMapper.MapToModel(period)).ToList() ?? new List<TimePeriodModel>()
        };

        public CarOfferEntity MapToEntity(CarOfferModel model) => new()
        {
            Id = model.Id,
            Heading = model.Heading,
            ShortDescription = model.ShortDescription,
            FeaturedImageUrl = model.FeaturedImageUrl,
            UrlHandle = model.UrlHandle,
            Horsepower = model.Horsepower,
            YearOfProduction = model.YearOfProduction,
            EngineDetails = model.EngineDetails,
            DriveDetails = model.DriveDetails,
            GearboxDetails = model.GearboxDetails,
            CarDeliverylocation = model.CarDeliverylocation,
            CarReturnLocation = model.CarReturnLocation,
            PublishedDate = model.PublishedDate,
            Visible = model.Visible,
            Tarrif = _tariffMapper.MapToEntity(model.Tarrif),
            Tags = model.Tags?.Select(tag => _tagMapper.MapToEntity(tag)).ToArray() ?? Array.Empty<CarTagEntity>(),
            ImageUrls = model.ImageUrls?.Select(image => _imageUrlMapper.MapToEntity(image)).ToArray() ?? Array.Empty<ImageUrlEntity>(),
            UnavailableDates = model.UnavailableDates?.Select(period => _timePeriodMapper.MapToEntity(period)).ToArray() ?? Array.Empty<TimePeriodEntity>()
        };
    }
    
    
    
}
