using Results.Domain;

namespace BlogModule.Application.Contract.BlogPostLikes.GetLikesCountForBlogPost;

public class GetLikesCountForBlogPostResponse : ISuccessResult
{
    public int TotalCount { get; init; }
}