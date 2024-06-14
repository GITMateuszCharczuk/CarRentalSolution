using RentalModule.Application.Contract.Tags.GetTags;
using RentalModule.Domain.RepositoryInterfaces.CarTag;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.QueryHandlers;

namespace RentalModule.Application.QueryHandlers.Tag.GetTags;

public class GetTagsQueryHandler : IQueryHandler<GetTagsQuery, HandlerResult<GetTagsResponse, IErrorResult>>
{
    private readonly ICarTagQueryRepository _repository;
        
    public GetTagsQueryHandler(ICarTagQueryRepository repository)
    {
        _repository = repository;
    }
        
    public async Task<HandlerResult<GetTagsResponse, IErrorResult>> Handle(GetTagsQuery request, CancellationToken cancellationToken)
    {
        var tags = await _repository.GetCollectionAsync(
            request.OrderBy, 
            request.OrderDirection,
            request.CarOfferId,
            cancellationToken);

        var response = new GetTagsResponse() {
            OrderBy = request.OrderBy,
            OrderDirection = request.OrderDirection,
            Items = tags
        };
            
        return response;
    }
}