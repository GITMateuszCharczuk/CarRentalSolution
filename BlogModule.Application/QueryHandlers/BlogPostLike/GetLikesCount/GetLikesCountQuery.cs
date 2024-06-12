using BlogModule.Application.Contract.BlogPostLikes.GetLikesCount;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.QueryHandlers;

namespace BlogModule.Application.QueryHandlers.BlogPostLike.GetLikesCount;

public record GetLikesCountQuery : IQuery<HandlerResult<GetLikesCountResponse, IErrorResult>>
{
    public string BlogPostId { get; init; } = string.Empty;
}