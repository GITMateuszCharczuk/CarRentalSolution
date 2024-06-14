using RentalModule.Domain.Models.Ids;
using Results.Domain;

namespace RentalModule.Application.Contract.CarOffers.CreateCarOffer;

public record CreateCarOfferResponse : INotificationResult, ISuccessResult
{
    public CarOfferId Id { get; init; }
    public string Title { get; init; } = string.Empty;
    public string Message { get; init; } = string.Empty;
}