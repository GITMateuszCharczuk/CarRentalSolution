using TrucksModule.Domain.Results.Base;

namespace TrucksModule.Application.Contract.Trucks.UpdateTruck;

public record UpdateTruckResponse : INotificationResult, ISuccessResult
{
    public string Title { get; init; } = string.Empty;
    public string Message { get; init; } = string.Empty;
}