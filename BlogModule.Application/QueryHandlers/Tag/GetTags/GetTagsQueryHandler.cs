using BlogModule.Application.Contract.Tags.GetTags;
using BlogModule.Domain.RepositoryInterfaces.Tag;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.QueryHandlers;

namespace BlogModule.Application.QueryHandlers.Tag.GetTags;

public class GetTagsQueryHandler : IQueryHandler<GetTagsQuery, HandlerResult<GetTagsResponse, IErrorResult>>
{
    private readonly ITagQueryRepository _repository;
        
    public GetTagsQueryHandler(ITagQueryRepository repository)
    {
        _repository = repository;
    }
        
    public async Task<HandlerResult<GetTagsResponse, IErrorResult>> Handle(GetTagsQuery request, CancellationToken cancellationToken)
    {
        var tags = await _repository.GetAllDistinctAsync(
            request.OrderBy, 
            request.OrderDirection,
            request.BlogPostId,
            cancellationToken);

        var response = new GetTagsResponse() {
            OrderBy = request.OrderBy,
            OrderDirection = request.OrderDirection,
            Items = tags
        };
            
        return response;
    }
}