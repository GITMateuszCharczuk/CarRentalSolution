using Results.Domain;

namespace BlogModule.Application.Contract.BlogPosts.UpdateBlogPost;

public record UpdateBlogPostResponse : INotificationResult, ISuccessResult
{
    public string Title { get; init; } = string.Empty;
    public string Message { get; init; } = string.Empty;
}