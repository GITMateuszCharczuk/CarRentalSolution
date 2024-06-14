using RentalModule.Domain.Models.Ids;
using Results.Domain;

namespace RentalModule.Application.Contract.CarOrders.CreateCarOrder;

public record CreateCarOrderResponse : INotificationResult, ISuccessResult
{
    public CarOrderId Id { get; init; }
    public string Title { get; init; } = string.Empty;
    public string Message { get; init; } = string.Empty;
}