using System.Collections.Immutable;
using RentalModule.Application.Contract.CarOffers.UpdateCarOffer;
using RentalModule.Domain.Models;
using RentalModule.Domain.Models.Ids;
using RentalModule.Domain.RepositoryInterfaces.CarOffer;
using Results.Application;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.CommandHandlers;

namespace RentalModule.Application.CommandHandlers.CarOffer.UpdateCarOffer;

public class UpdateCarOfferCommandHandler : ICommandHandler<UpdateCarOfferCommand, HandlerResult<UpdateCarOfferResponse, IErrorResult>>
{
    private readonly ICarOfferCommandRepository _commandRepository;
    private readonly ICarOfferQueryRepository _queryRepository;
    
    public UpdateCarOfferCommandHandler(
        ICarOfferCommandRepository repository, 
        ICarOfferQueryRepository queryRepository)
    {
        _commandRepository = repository;
        _queryRepository = queryRepository;
    }
    
    public async Task<HandlerResult<UpdateCarOfferResponse, IErrorResult>> Handle(UpdateCarOfferCommand request, CancellationToken cancellationToken)
    {
        var carOffer = request.Id is null
            ? await _queryRepository.GetByUrlAsync(request.UrlHandle, cancellationToken) 
            : await _queryRepository.GetByIdAsync(request.Id.Value, cancellationToken);
        
        if (carOffer is null) return new EntityNotFoundErrorResult() {
            Title = "Cannot update blog post",
            Message = $"Blog post with url {request.UrlHandle} does not exist in the database."
        };

        var carOfferId = request.Id ?? carOffer.Id;
        var carOfferToUpdate = new CarOfferModel
        {
            Id = carOfferId,
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
                CarOfferId = carOfferId,
                OneNormalDayPrice = request.OneNormalDayPrice,
                OneWeekendDayPrice = request.OneWeekendDayPrice,
                FullWeekendPrice = request.FullWeekendPrice,
                OneWeekPrice = request.OneWeekPrice,
                OneMonthPrice = request.OneMonthPrice
            },
            Tags = ConvertToTagModels(request.Tags, carOfferId),
            ImageUrls = ConvertToImageUrlModels(request.ImageUrls, carOfferId)
        };
        
        var updatedCarOffer = await _commandRepository.UpdateAsync(carOfferToUpdate, cancellationToken);
        
        return new UpdateCarOfferResponse() {
            Title = "Car offer updated",
            Message = $"Car offer with url {updatedCarOffer.UrlHandle} was updated in the database."
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