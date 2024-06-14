using RentalModule.Domain.Models.Ids;
using Results.Domain;

namespace RentalModule.Application.Contract.ImageUrls.AddUrlToCarOffer;

public class AddUrlToCarOfferResponse : ISuccessResult, INotificationResult
{
    public ImageUrlId Id { get; init; }
    public string Title { get; init; } = string.Empty;
    public string Message { get; init; } = string.Empty;
}