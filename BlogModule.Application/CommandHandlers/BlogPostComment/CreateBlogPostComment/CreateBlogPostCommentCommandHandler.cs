using System.ComponentModel.DataAnnotations;
using BlogModule.Application.Contract.BlogPostComments.CreateBlogPostComment;
using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using BlogModule.Domain.RepositoryInterfaces.BlogPostComment;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.CommandHandlers;

namespace BlogModule.Application.CommandHandlers.BlogPostComment.CreateBlogPostComment;

public class CreateBlogPostCommentCommandHandler : ICommandHandler<CreateBlogPostCommentCommand, HandlerResult<CreateBlogPostCommentResponse, IErrorResult>>
{
    private readonly IBlogPostCommentCommandRepository _repository;

    public CreateBlogPostCommentCommandHandler(IBlogPostCommentCommandRepository repository)
    {
        _repository = repository;
    }

    public async Task<HandlerResult<CreateBlogPostCommentResponse,IErrorResult>> Handle(CreateBlogPostCommentCommand request, CancellationToken cancellationToken)
    {//dodać walidacje todo
        var context = new ValidationContext(request);
        Validator.ValidateObject(request, context, validateAllProperties: true);
        
        var newComment = new BlogPostCommentModel
        {
            Id = new BlogPostCommentId(Guid.NewGuid()),
            BlogPostId = request.BlogPostId,
            UserId = request.UserId,
            Description = request.Description,
            DateAdded = DateTime.Now
        };

        await _repository.AddAsync(newComment, cancellationToken);

        var response = new CreateBlogPostCommentResponse
        {
            Id = newComment.Id,
            Title = "Comment Created",
            Message = "The blog post comment has been successfully created."
        };

        return response;
    }
}