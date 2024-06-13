using BlogModule.Application.Contract.BlogPostLikes.CreateLikeForBlogPost;
using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using BlogModule.Domain.RepositoryInterfaces.BlogPostLike;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.CommandHandlers;

namespace BlogModule.Application.CommandHandlers.BlogPostLike.CreateLikeForBlogPost;

public class CreateLikeForBlogPostCommandHandler : ICommandHandler<CreateLikeForBlogPostCommand, HandlerResult<CreateLikeForBlogPostResponse, IErrorResult>>
{
    private readonly IBlogPostLikeCommandRepository _repository;

    public CreateLikeForBlogPostCommandHandler(IBlogPostLikeCommandRepository repository)
    {
        _repository = repository;
    }
    public async Task<HandlerResult<CreateLikeForBlogPostResponse, IErrorResult>> Handle(CreateLikeForBlogPostCommand request, CancellationToken cancellationToken)
    {//do poprawy z repozytorium dodać validacje ilości lików
        var newLike = new BlogPostLikeModel//todo
        {
            Id = new BlogPostLikeId(new Guid()),
            BlogPostId = request.BlogPostId,
            UserId = request.UserId
        };

        var addedLike = await _repository.AddAsync(newLike, cancellationToken);
        
        return new CreateLikeForBlogPostResponse
        {
            Id = addedLike.Id,
            Title = "Like added",
            Message = $"Like was added to blog post of id: {addedLike.BlogPostId}."//todo
        };

    }
}