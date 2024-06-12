using BlogModule.Application.Contract.BlogPostComments.GetBlogPostComments;
using BlogModule.Domain.RepositoryInterfaces.BlogPostComment;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.QueryHandlers;

namespace BlogModule.Application.QueryHandlers.BlogPostComment.GetBlogPostComments;

public class GetBlogPostCommentsQueryHandler : IQueryHandler<GetBlogPostCommentsQuery, HandlerResult<GetBlogPostCommentsResponse, IErrorResult>>
{
    private readonly IBlogPostCommentQueryRepository _repository;

    public GetBlogPostCommentsQueryHandler(IBlogPostCommentQueryRepository repository)
    {
        _repository = repository;
    }

    public async Task<HandlerResult<GetBlogPostCommentsResponse, IErrorResult>> Handle(GetBlogPostCommentsQuery request, CancellationToken cancellationToken)
    {
        var comments = await _repository.GetCollectionAsync(
            request.Page,
            request.PageSize,
            request.BlogPostId,
            request.OrderBy,
            request.OrderDirection,
            request.Ids,
            request.DateTimes,
            request.UserIds,
            cancellationToken);

        var commentsCount = await _repository.GetTotalCommentsCountAsync(
            request.BlogPostId,
            request.Ids,
            request.DateTimes,
            request.UserIds,
            cancellationToken);
        
        var response = new GetBlogPostCommentsResponse
        {
            Page = request.Page,
            PageSize = request.PageSize,
            TotalCount = commentsCount,
            OrderBy = request.OrderBy,
            OrderDirection = request.OrderDirection,
            Items = comments
        };

        return response;
    }
}