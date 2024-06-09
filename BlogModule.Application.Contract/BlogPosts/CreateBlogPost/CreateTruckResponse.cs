using Results.Domain;

namespace BlogModule.Application.Contract.BlogPosts.CreateBlogPost;

public record CreateTruckResponse : INotificationResult, ISuccessResult
{
    public TruckId Id { get; init; }
    public string Title { get; init; } = string.Empty;
    public string Message { get; init; } = string.Empty;
}