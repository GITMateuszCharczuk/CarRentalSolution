using Results.Domain;

namespace BlogModule.Application.Contract.BlogPostLikes.GetLikesCountForBlogPost;

public class GetLikesCountResponse : ISuccessResult
{
    public int TotalCount { get; init; }
}