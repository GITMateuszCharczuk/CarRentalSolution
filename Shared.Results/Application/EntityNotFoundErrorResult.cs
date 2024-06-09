using Results.Domain;

namespace Results.Application;

public record EntityNotFoundErrorResult : INotificationResult, IErrorResult
{
    public string Title { get; init; } = string.Empty;
    public string Message { get; init; } = string.Empty;
}