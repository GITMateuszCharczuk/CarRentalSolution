using BlogModule.Application.Contract.BlogPostLikes.GetLikesForBlogPost;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.QueryHandlers;

namespace BlogModule.Application.QueryHandlers.BlogPostLike.GetLikesForBlogPost;

public record GetLikesForBlogPostQuery : IQuery<HandlerResult<GetLikesForBlogPostResponse, IErrorResult>>
{
    public string BlogPostId { get; init; } = string.Empty;
}