using BlogModule.Application.Contract.BlogPosts.DeleteBlogPost;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.CommandHandlers;

namespace BlogModule.Application.CommandHandlers.BlogPost.DeleteBlogPost;

public class DeleteBlogPostCommand : ICommand<HandlerResult<DeleteBlogPostResponse, IErrorResult>>
{
    public string Id { get; init; } = string.Empty;
}