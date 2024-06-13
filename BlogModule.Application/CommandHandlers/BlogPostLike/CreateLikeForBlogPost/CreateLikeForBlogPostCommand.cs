using BlogModule.Application.Contract.BlogPostLikes.CreateLikeForBlogPost;
using BlogModule.Domain.Models.Ids;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.CommandHandlers;

namespace BlogModule.Application.CommandHandlers.BlogPostLike.CreateLikeForBlogPost;

public class CreateLikeForBlogPostCommand : ICommand<HandlerResult<CreateLikeForBlogPostResponse, IErrorResult>>
{
    public BlogPostId BlogPostId { get; init; } 
    public Guid UserId { get; set; } 
}