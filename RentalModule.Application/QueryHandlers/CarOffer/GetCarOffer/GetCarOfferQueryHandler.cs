using System.Collections.Immutable;
using RentalModule.Application.Contract.CarOffers.GetCarOffer;
using RentalModule.Domain.Models.Ids;
using RentalModule.Domain.RepositoryInterfaces.CarOffer;
using Results.Application;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.QueryHandlers;

namespace RentalModule.Application.QueryHandlers.CarOffer.GetCarOffer;

public class GetCarOfferQueryHandler : IQueryHandler<GetCarOfferQuery, HandlerResult<GetCarOfferResponse, IErrorResult>>
{
    private readonly ICarOfferQueryRepository _repository;

    public GetCarOfferQueryHandler(ICarOfferQueryRepository repository)
    {
        _repository = repository;
    }

    public async Task<HandlerResult<GetCarOfferResponse, IErrorResult>> Handle(GetCarOfferQuery request,
        CancellationToken cancellationToken)
    {
        var isSuccess = Guid.TryParse(request.IdOrUrl, out var blogPostId);

        var carOffer = isSuccess
            ? await _repository.GetByIdAsync(new CarOfferId(blogPostId), cancellationToken)
            : await _repository.GetByUrlAsync(request.IdOrUrl, cancellationToken);


        return carOffer is null ? new EntityNotFoundErrorResult()
            {
                Title = "Car offer was not found",
                Message = $"Car offer with ID/Url {request.IdOrUrl} was not found in the database."
            } : new GetCarOfferResponse
        {
            Id = carOffer.Id,
            Heading = carOffer.Heading,
            ShortDescription = carOffer.ShortDescription,
            FeaturedImageUrl = carOffer.FeaturedImageUrl,
            UrlHandle = carOffer.UrlHandle,
            Horsepower = carOffer.Horsepower,
            YearOfProduction = carOffer.YearOfProduction,
            EngineDetails = carOffer.EngineDetails,
            DriveDetails = carOffer.DriveDetails,
            GearboxDetails = carOffer.GearboxDetails,
            CarDeliverylocation = carOffer.CarDeliverylocation,
            CarReturnLocation = carOffer.CarReturnLocation,
            PublishedDate = carOffer.PublishedDate,
            Visible = carOffer.Visible,
            Tarrif = carOffer.Tarrif,
            Tags = carOffer.Tags.ToImmutableArray(),
            ImageUrls = carOffer.ImageUrls.ToImmutableArray(),
            UnavailableDates = carOffer.UnavailableDates.ToImmutableArray()
        };
    }
}