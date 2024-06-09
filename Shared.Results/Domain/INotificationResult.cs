namespace Results.Domain;

public interface INotificationResult
{
    public string Title { get; init; }
    public string Message { get; init; }
}