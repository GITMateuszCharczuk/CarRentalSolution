using RentalModule.Application.Contract.CarOffers.GetCarOffer;
using RentalModule.Application.Contract.CarOrders.GetCarOrder;
using RentalModule.Domain.Models.Ids;
using RentalModule.Domain.RepositoryInterfaces.CarOffer;
using RentalModule.Domain.RepositoryInterfaces.CarOrder;
using Results.Application;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.QueryHandlers;

namespace RentalModule.Application.QueryHandlers.CarOrder.GetCarOrder;

public class GetCarOrderQueryHandler : IQueryHandler<GetCarOrderQuery, HandlerResult<GetCarOrderResponse, IErrorResult>>
{
    private readonly ICarOrderQueryRepository _repository;

    public GetCarOrderQueryHandler(ICarOrderQueryRepository repository)
    {
        _repository = repository;
    }

    public async Task<HandlerResult<GetCarOrderResponse, IErrorResult>> Handle(GetCarOrderQuery request,
        CancellationToken cancellationToken)
    {
        Guid.TryParse(request.IdOrUrl, out var blogPostId);

        var carOrder = await _repository.GetByIdAsync(new CarOrderId(blogPostId), cancellationToken);

        return carOrder is null ? new EntityNotFoundErrorResult()
            {
                Title = "Blog post was not found",
                Message = $"Blog post with ID/Url {request.IdOrUrl} was not found in the database."
            } : new GetCarOrderResponse
        {
            Id = carOrder.Id,
            UserId = carOrder.UserId,
            CarOfferId = carOrder.CarOfferId,
            StartDate = carOrder.StartDate,
            EndDate = carOrder.EndDate,
            Notes = carOrder.Notes,
            NumOfDrivers = carOrder.NumOfDrivers,
            TotalCost = carOrder.TotalCost
        };
    }
}