using Results.Domain;

namespace RentalModule.Application.Contract.CarOffers.UpdateCarOffer;

public record UpdateCarOfferResponse : INotificationResult, ISuccessResult
{
    public string Title { get; init; } = string.Empty;
    public string Message { get; init; } = string.Empty;
}