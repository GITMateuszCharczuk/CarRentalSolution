using System.Collections.Immutable;
using RentalModule.Domain.Models.Ids;
using Results.Domain;

namespace RentalModule.Application.Contract.CarOrders.GetCarOrder;

public record GetCarOrderResponse :  ISuccessResult
{
    public CarOrderId Id { get; init; }

    public Guid UserId { get; init; }
    
    public CarOfferId CarOfferId { get; init; }
    
    public DateTime StartDate { get; init; }
    
    public DateTime EndDate { get; init; }

    public string Notes { get; init; } = string.Empty;
    
    public int NumOfDrivers { get; init; }
    
    public Double TotalCost { get; set; }
}