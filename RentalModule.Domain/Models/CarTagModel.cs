using RentalModule.Domain.Models.Ids;

namespace RentalModule.Domain.Models;

public record CarTagModel
{
    public CarTagId Id { get; init; }
    public string Name { get; init; } = string.Empty;
    
    public CarOfferId CarOfferId { get; init; }
}