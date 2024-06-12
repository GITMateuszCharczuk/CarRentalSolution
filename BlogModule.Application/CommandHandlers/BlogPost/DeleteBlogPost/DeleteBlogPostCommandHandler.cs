using BlogModule.Application.CommandHandlers.DeleteBlogPost;
using BlogModule.Application.Contract.BlogPosts.DeleteBlogPost;
using BlogModule.Domain.Models.Ids;
using BlogModule.Domain.RepositoryInterfaces.BlogPost;
using Results.Application;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.CommandHandlers;

namespace BlogModule.Application.CommandHandlers.BlogPost.DeleteBlogPost;

public class DeleteBlogPostCommandHandler : ICommandHandler<DeleteBlogPostCommand, HandlerResult<DeleteBlogPostResponse, IErrorResult>>
{
    private readonly IBlogPostCommandRepository _commandRepository;
    private readonly IBlogPostQueryRepository _queryRepository;

    public DeleteBlogPostCommandHandler(
        IBlogPostCommandRepository repository, 
        IBlogPostQueryRepository queryRepository)
    {
        _commandRepository = repository;
        _queryRepository = queryRepository;
    }

    public async Task<HandlerResult<DeleteBlogPostResponse, IErrorResult>> Handle(DeleteBlogPostCommand request, CancellationToken cancellationToken)
    {
        var isSuccess = Guid.TryParse(request.Id, out var blogPostId);
        var blogPost = isSuccess ? await _queryRepository.GetByIdAsync(new BlogPostId(blogPostId), cancellationToken) : null;
        
        if (blogPost is null) return new EntityNotFoundErrorResult() {
            Title = "Cannot delete blog post",
            Message = $"Blog post with ID {request.Id} was not found in the database."
        };
        
        await _commandRepository.DeleteAsync(blogPost.Id, cancellationToken);

        return new DeleteBlogPostResponse() {
            Title = "Blog post deleted",
            Message = $"Blog post with ID {request.Id} was deleted from the database."
        };
    }
}