using BlogModule.Domain.Models.Ids;

namespace BlogModule.Application.Contract.BlogPostLikes.CreateLikeForBlogPost;

public class CreateLikeForBlogPostRequest
{
    public BlogPostId BlogPostId { get; init; }
    public Guid UserId { get; set; }
}