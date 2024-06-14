using RentalModule.Domain.Models.Ids;

namespace RentalModule.Domain.Models;

public record CarOrderModel
{
    public CarOrderId Id { get; init; }

    public Guid UserId { get; init; }
    
    public CarOfferId CarOfferId { get; init; }
    
    public DateTime StartDate { get; init; }
    
    public DateTime EndDate { get; init; }

    public string? Notes { get; init; } = string.Empty;
    
    public int NumOfDrivers { get; init; }
    
    public Double TotalCost { get; set; }
}