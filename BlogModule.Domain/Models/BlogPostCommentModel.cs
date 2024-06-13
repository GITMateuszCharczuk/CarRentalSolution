using BlogModule.Domain.Models.Ids;

namespace BlogModule.Domain.Models;

public record BlogPostCommentModel
{
    public BlogPostCommentId Id { get; init; }

    public string Description { get; set; }

    public BlogPostId BlogPostId { get; set; }

    public Guid UserId { get; set; }

    public DateTime DateAdded { get; set; }
}