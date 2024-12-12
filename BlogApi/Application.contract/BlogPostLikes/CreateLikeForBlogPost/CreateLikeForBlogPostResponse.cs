using BlogModule.Domain.Models.Ids;
using Results.Domain;

namespace BlogModule.Application.Contract.BlogPostLikes.CreateLikeForBlogPost;

public class CreateLikeForBlogPostResponse : ISuccessResult, INotificationResult
{
    public BlogPostLikeId Id { get; init; }
    public string Title { get; init; } = string.Empty;
    public string Message { get; init; } = string.Empty;
}