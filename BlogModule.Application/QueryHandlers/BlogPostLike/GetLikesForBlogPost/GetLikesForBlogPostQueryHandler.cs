using BlogModule.Application.Contract.BlogPostLikes.GetLikesForBlogPost;
using BlogModule.Domain.Models.Ids;
using BlogModule.Domain.RepositoryInterfaces.BlogPostLike;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.QueryHandlers;

namespace BlogModule.Application.QueryHandlers.BlogPostLike.GetLikesForBlogPost;

public class GetLikesForBlogPostQueryHandler : IQueryHandler<GetLikesForBlogPostQuery, HandlerResult<GetLikesForBlogPostResponse, IErrorResult>>
{
    private readonly IBlogPostLikeQueryRepository _repository;

    public GetLikesForBlogPostQueryHandler(IBlogPostLikeQueryRepository repository)
    {
        _repository = repository;
    }

    public async Task<HandlerResult<GetLikesForBlogPostResponse, IErrorResult>> Handle(GetLikesForBlogPostQuery request, CancellationToken cancellationToken)
    {
        var blogPostId = new BlogPostId(Guid.Parse(request.BlogPostId));

        var likes = await _repository.GetLikesForBlogPostAsync(blogPostId, cancellationToken);

        var response = new GetLikesForBlogPostResponse
        {
            Items = likes
        };

        return response;
    }
}