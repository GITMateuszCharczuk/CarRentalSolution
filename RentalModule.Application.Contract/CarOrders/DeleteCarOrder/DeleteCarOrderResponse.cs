using Results.Domain;

namespace RentalModule.Application.Contract.CarOrders.DeleteCarOrder;

public record DeleteCarOrderResponse : INotificationResult, ISuccessResult
{
    public string Title { get; init; } = string.Empty;
    public string Message { get; init; } = string.Empty;
}