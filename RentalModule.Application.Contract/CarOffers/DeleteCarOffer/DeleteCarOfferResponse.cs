using Results.Domain;

namespace RentalModule.Application.Contract.CarOffers.DeleteCarOffer;

public record DeleteCarOfferResponse : INotificationResult, ISuccessResult
{
    public string Title { get; init; } = string.Empty;
    public string Message { get; init; } = string.Empty;
}