using BlogModule.Application.Contract.BlogPosts.GetBlogPost;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.QueryHandlers;

namespace BlogModule.Application.QueryHandlers.BlogPost.GetBlogPost;

public record GetBlogPostQuery : IQuery<HandlerResult<GetBlogPostResponse, IErrorResult>>
{
    public string IdOrUrl { get; init; } = string.Empty;
}