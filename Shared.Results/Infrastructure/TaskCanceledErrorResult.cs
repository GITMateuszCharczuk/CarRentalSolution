using Results.Domain;

namespace Results.Infrastructure;

public record TaskCanceledErrorResult : INotificationResult, IErrorResult
{
    public string Title { get; init; } = string.Empty;
    public string Message { get; init; } = string.Empty;
}