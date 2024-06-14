using RentalModule.Domain.Models.Ids;

namespace RentalModule.Application.Contract.ImageUrls.AddUrlToCarOffer;

public class AddUrlToCarOfferRequest
{
    public CarOfferId CarOfferId { get; init; }
    public string Url { get; init; } = string.Empty;
}