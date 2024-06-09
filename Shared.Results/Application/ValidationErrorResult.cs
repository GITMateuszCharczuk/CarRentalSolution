using Results.Domain;

namespace Results.Application;

public record ValidationErrorResult : INotificationResult, IErrorResult
{
    public string Title { get; init; } = string.Empty;
    public string Message { get; init; } = string.Empty;
    public ValidationError[] Errors { get; init; } = Array.Empty<ValidationError>();
    
    public record ValidationError
    {
        public string PropertyName { get; init; } = string.Empty;
        public string ErrorMessage { get; init; } = string.Empty;
    }
}