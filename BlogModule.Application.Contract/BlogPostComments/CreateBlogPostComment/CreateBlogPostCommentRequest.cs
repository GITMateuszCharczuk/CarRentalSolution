using BlogModule.Domain.Models.Ids;

namespace BlogModule.Application.Contract.BlogPostComments.CreateBlogPostComment;

public class CreateBlogPostCommentRequest
{
    public string Description { get; set; } = string.Empty;

    public BlogPostId BlogPostId { get; set; }

    public Guid UserId { get; set; }

    public DateTime DateAdded { get; set; } = DateTime.Now;
}