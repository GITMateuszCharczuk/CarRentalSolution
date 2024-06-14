using RentalModule.Application.Contract.CarOffers.GetCarOffer;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.QueryHandlers;

namespace RentalModule.Application.QueryHandlers.CarOffer.GetCarOffer;

public record GetCarOfferQuery : IQuery<HandlerResult<GetCarOfferResponse, IErrorResult>>
{
    public string IdOrUrl { get; init; } = string.Empty;
}