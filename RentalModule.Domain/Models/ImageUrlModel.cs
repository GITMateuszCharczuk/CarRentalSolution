using RentalModule.Domain.Models.Ids;

namespace RentalModule.Domain.Models;

public record ImageUrlModel
{
    public ImageUrlId Id { get; init; }
    public CarOfferId CarOfferId { get; init; }
    public string Url { get; init; } = string.Empty;
}