using RentalModule.Application.Contract.CarOrders.GetCarOrders;
using RentalModule.Domain.RepositoryInterfaces.CarOrder;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.QueryHandlers;

namespace RentalModule.Application.QueryHandlers.CarOrder.GetCarOrders;

public class GetCarOrdersQueryHandler : IQueryHandler<GetCarOrdersQuery, HandlerResult<GetCarOrdersResponse, IErrorResult>>
{
    private readonly ICarOrderQueryRepository _repository;
    
    public GetCarOrdersQueryHandler(ICarOrderQueryRepository repository)
    {
        _repository = repository;
    }
    
    public async Task<HandlerResult<GetCarOrdersResponse, IErrorResult>> Handle(GetCarOrdersQuery request, CancellationToken cancellationToken)
    {
        var carOffers = await _repository.GetCollectionAsync(
            request.Page, 
            request.PageSize, 
            request.OrderBy, 
            request.OrderDirection,
            request.Dates,
            request.UserId,
            request.CarOfferId,
            cancellationToken);

        var carOffersCount = await _repository.GetTotalCountAsync(cancellationToken);
        
        var response = new GetCarOrdersResponse() {
            Page = request.Page,
            PageSize = request.PageSize,
            TotalCount = carOffersCount,
            OrderBy = request.OrderBy,
            OrderDirection = request.OrderDirection,
            Items = carOffers
        };
        
        return response;
    }
}