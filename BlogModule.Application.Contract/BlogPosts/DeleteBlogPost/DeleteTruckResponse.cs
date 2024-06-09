using Results.Domain;

namespace BlogModule.Application.Contract.BlogPosts.DeleteBlogPost;

public record DeleteTruckResponse : INotificationResult, ISuccessResult
{
    public string Title { get; init; } = string.Empty;
    public string Message { get; init; } = string.Empty;
}