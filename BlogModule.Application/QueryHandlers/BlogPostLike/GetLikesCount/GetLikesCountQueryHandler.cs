using BlogModule.Application.Contract.BlogPostLikes.GetLikesCount;
using BlogModule.Domain.Models.Ids;
using BlogModule.Domain.RepositoryInterfaces.BlogPostLike;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.QueryHandlers;

namespace BlogModule.Application.QueryHandlers.BlogPostLike.GetLikesCount;

public class
    GetLikesCountQueryHandler : IQueryHandler<GetLikesCountQuery, HandlerResult<GetLikesCountResponse, IErrorResult>>
{
    private readonly IBlogPostLikeQueryRepository _repository;

    public GetLikesCountQueryHandler(IBlogPostLikeQueryRepository repository)
    {
        _repository = repository;
    }

    public async Task<HandlerResult<GetLikesCountResponse, IErrorResult>> Handle(GetLikesCountQuery request,
        CancellationToken cancellationToken)
    {
        var blogPostId = new BlogPostId(Guid.Parse(request.BlogPostId));

        var likesCount = await _repository.GetTotalLikesForBlogPostAsync(blogPostId, cancellationToken);

        var response = new GetLikesCountResponse
        {
            TotalCount = likesCount
        };

        return response;
    }
}