using BlogModule.Domain.Models.Ids;
using Results.Domain;

namespace BlogModule.Application.Contract.BlogPosts.CreateBlogPost;

public record CreateBlogPostResponse : INotificationResult, ISuccessResult
{
    public BlogPostId Id { get; init; }
    public string Title { get; init; } = string.Empty;
    public string Message { get; init; } = string.Empty;
}