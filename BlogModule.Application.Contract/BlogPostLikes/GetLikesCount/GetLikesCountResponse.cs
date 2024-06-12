using Results.Domain;

namespace BlogModule.Application.Contract.BlogPostLikes.GetLikesCount;

public class GetLikesCountResponse : ISuccessResult
{
    public int TotalCount { get; init; }
}