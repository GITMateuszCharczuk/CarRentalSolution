using BlogModule.Domain.Models.Ids;
using Results.Domain;

namespace BlogModule.Application.Contract.BlogPostComments.CreateBlogPostComment;

public class CreateBlogPostCommentResponse : ISuccessResult, INotificationResult
{
    public BlogPostCommentId Id { get; init; }
    public string Title { get; init; } = string.Empty;
    public string Message { get; init; } = string.Empty;
}