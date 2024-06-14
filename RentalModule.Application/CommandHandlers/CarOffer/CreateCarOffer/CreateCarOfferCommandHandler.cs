using System.Collections.Immutable;
using RentalModule.Application.Contract.CarOffers.CreateCarOffer;
using RentalModule.Domain.Models;
using RentalModule.Domain.Models.Ids;
using RentalModule.Domain.RepositoryInterfaces.CarOffer;
using RentalModule.Infrastructure.DataBase.Entities;
using Results.Application;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.CommandHandlers;

namespace RentalModule.Application.CommandHandlers.CarOffer.CreateCarOffer
{
    public class CreateCarOfferCommandHandler : ICommandHandler<CreateCarOfferCommand,
        HandlerResult<CreateCarOfferResponse, IErrorResult>>
    {
        private readonly ICarOfferCommandRepository _commandRepository;
        private readonly ICarOfferQueryRepository _queryRepository;

        public CreateCarOfferCommandHandler(
            ICarOfferCommandRepository commandRepository,
            ICarOfferQueryRepository queryRepository)
        {
            _commandRepository = commandRepository;
            _queryRepository = queryRepository;
        }

        public async Task<HandlerResult<CreateCarOfferResponse, IErrorResult>> Handle(CreateCarOfferCommand request,
            CancellationToken cancellationToken)
        {
            var carOffer = await _queryRepository.GetByUrlAsync(request.UrlHandle, cancellationToken);
            if (carOffer is not null)
            {
                return new EntityAlreadyExistsErrorResult
                {
                    Title = "Cannot create car offer",
                    Message = $"Car offer with URL handle {request.UrlHandle} already exists."
                };
            }

            var newCarOfferId = new CarOfferId(Guid.NewGuid());
            var carOfferToAdd = new CarOfferModel
            {
                Id = newCarOfferId,
                Heading = request.Heading,
                ShortDescription = request.ShortDescription,
                FeaturedImageUrl = request.FeaturedImageUrl,
                UrlHandle = request.UrlHandle,
                Horsepower = request.Horsepower,
                YearOfProduction = request.YearOfProduction,
                EngineDetails = request.EngineDetails,
                DriveDetails = request.DriveDetails,
                GearboxDetails = request.GearboxDetails,
                CarDeliverylocation = request.CarDeliverylocation,
                CarReturnLocation = request.CarReturnLocation,
                PublishedDate = request.PublishedDate,
                Visible = request.Visible,
                Tarrif = new CarTariffModel
                {
                    Id = new CarTariffId(new Guid()),
                    CarOfferId = newCarOfferId,
                    OneNormalDayPrice = request.OneNormalDayPrice,
                    OneWeekendDayPrice = request.OneWeekendDayPrice,
                    FullWeekendPrice = request.FullWeekendPrice,
                    OneWeekPrice = request.OneWeekPrice,
                    OneMonthPrice = request.OneMonthPrice
                },
                Tags = ConvertToTagModels(request.Tags,
                    newCarOfferId),
                ImageUrls = ConvertToImageUrlModels(request.ImageUrls,
                    newCarOfferId),
                UnavailableDates = ImmutableArray<TimePeriodModel>.Empty
            };

            var addedCarOffer = await _commandRepository.AddAsync(carOfferToAdd, cancellationToken);

            return new CreateCarOfferResponse()
            {
                Id = addedCarOffer.Id,
                Title = "Car offer created",
                Message = $"Car offer '{addedCarOffer.Heading}' was created successfully."
            };
        }

        private static ImmutableArray<CarTagModel> ConvertToTagModels(ImmutableArray<string>? tags,
            CarOfferId carOfferId)
        {
            if (tags == null || tags.Value.IsEmpty) return ImmutableArray<CarTagModel>.Empty;

            return tags.Value.Select(tag => new CarTagModel
            {
                Id = new CarTagId(Guid.NewGuid()),
                Name = tag,
                CarOfferId = carOfferId
            }).ToImmutableArray();
        }
        
        private static ImmutableArray<ImageUrlModel> ConvertToImageUrlModels(ImmutableArray<string>? urls,
            CarOfferId carOfferId)
        {
            if (urls == null || urls.Value.IsEmpty) return ImmutableArray<ImageUrlModel>.Empty;

            return urls.Value.Select(url => new ImageUrlModel
            {
                Id = new ImageUrlId(Guid.NewGuid()),
                Url = url,
                CarOfferId = carOfferId,

            }).ToImmutableArray();
        }
    }
}