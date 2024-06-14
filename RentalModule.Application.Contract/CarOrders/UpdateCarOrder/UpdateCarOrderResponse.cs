using Results.Domain;

namespace RentalModule.Application.Contract.CarOrders.UpdateCarOrder;

public record UpdateCarOrderResponse : INotificationResult, ISuccessResult
{
    public string Title { get; init; } = string.Empty;
    public string Message { get; init; } = string.Empty;
}