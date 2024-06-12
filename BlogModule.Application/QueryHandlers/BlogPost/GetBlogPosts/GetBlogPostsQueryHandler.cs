using BlogModule.Application.Contract.BlogPosts.GetBlogPosts;
using BlogModule.Domain.RepositoryInterfaces.BlogPost;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.QueryHandlers;

namespace BlogModule.Application.QueryHandlers.BlogPost.GetBlogPosts;

public class GetBlogPostsQueryHandler : IQueryHandler<GetBlogPostsQuery, HandlerResult<GetBlogPostsResponse, IErrorResult>>
{
    private readonly IBlogPostQueryRepository _repository;
    
    public GetBlogPostsQueryHandler(IBlogPostQueryRepository repository)
    {
        _repository = repository;
    }
    
    public async Task<HandlerResult<GetBlogPostsResponse, IErrorResult>> Handle(GetBlogPostsQuery request, CancellationToken cancellationToken)
    {
        var blogPosts = await _repository.GetCollectionAsync(
            request.Page, 
            request.PageSize, 
            request.OrderBy, 
            request.OrderDirection,
            request.Ids,
            request.PublishedDates,
            request.Authors,
            cancellationToken);

        var blogPostsCount = await _repository.GetTotalCountAsync(cancellationToken);
        
        var response = new GetBlogPostsResponse() {
            Page = request.Page,
            PageSize = request.PageSize,
            TotalCount = blogPostsCount,
            OrderBy = request.OrderBy,
            OrderDirection = request.OrderDirection,
            Items = blogPosts
        };
        
        return response;
    }
}