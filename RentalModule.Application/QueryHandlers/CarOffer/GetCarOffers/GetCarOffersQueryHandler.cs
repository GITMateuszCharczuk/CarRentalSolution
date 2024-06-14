using RentalModule.Application.Contract.CarOffers.GetCarOffers;
using RentalModule.Domain.RepositoryInterfaces.CarOffer;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.QueryHandlers;

namespace RentalModule.Application.QueryHandlers.CarOffer.GetCarOffers;

public class GetCarOffersQueryHandler : IQueryHandler<GetCarOffersQuery, HandlerResult<GetCarOffersResponse, IErrorResult>>
{
    private readonly ICarOfferQueryRepository _repository;
    
    public GetCarOffersQueryHandler(ICarOfferQueryRepository repository)
    {
        _repository = repository;
    }
    
    public async Task<HandlerResult<GetCarOffersResponse, IErrorResult>> Handle(GetCarOffersQuery request, CancellationToken cancellationToken)
    {
        var carOffers = await _repository.GetCollectionAsync(
            request.Page, 
            request.PageSize, 
            request.OrderBy, 
            request.OrderDirection,
            request.PossibleDates,
            request.Tags,
            cancellationToken);

        var carOffersCount = await _repository.GetTotalCountAsync(cancellationToken);
        
        var response = new GetCarOffersResponse() {
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