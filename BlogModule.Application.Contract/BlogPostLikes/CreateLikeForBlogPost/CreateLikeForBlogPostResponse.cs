using Results.Domain;

namespace BlogModule.Application.Contract.BlogPostLikes.CreateLikeForBlogPost;

public class CreateLikeForBlogPostResponse : ISuccessResult, INotificationResult
{
    public string Title { get; init; } = string.Empty;
    public string Message { get; init; } = string.Empty;
}