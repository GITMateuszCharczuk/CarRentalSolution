using RentalModule.Application.Contract.CarOrders.CreateCarOrder;
using RentalModule.Domain.Models;
using RentalModule.Domain.Models.Ids;
using RentalModule.Domain.RepositoryInterfaces.CarOrder;
using Results.Application;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.CommandHandlers;

namespace RentalModule.Application.CommandHandlers.CarOrder.CreateCarOrder;

public class CreateCarOrderCommandHandler : ICommandHandler<CreateCarOrderCommand, HandlerResult<CreateCarOrderResponse, IErrorResult>> {
    
    private readonly ICarOrderCommandRepository _commandRepository;
    private readonly ICarOrderQueryRepository _queryRepository;

    public CreateCarOrderCommandHandler(
        ICarOrderCommandRepository commandRepository,
        ICarOrderQueryRepository queryRepository)
    {
        _commandRepository = commandRepository;
        _queryRepository = queryRepository;
    }

    public async Task<HandlerResult<CreateCarOrderResponse, IErrorResult>> Handle(CreateCarOrderCommand request,
        CancellationToken cancellationToken)
    {
        //todo check na to czy już jest jakaś oferta na tą fure
        // var carOrder =
        //     await _queryRepository.GetByIdAsync(request.CarOfferId, cancellationToken);
        // if (carOrder != null)
        //     return new EntityAlreadyExistsErrorResult()
        //     {
        //         Title = "Cannot create car order",
        //         Message =
        //             $"Car order for the car offer {request.CarOfferId} and user {request.UserId} already exists in the database."
        //     };

        var newCarOrderId = new CarOrderId(Guid.NewGuid());
        var carOrderToAdd = new CarOrderModel
        {
            Id = newCarOrderId,
            UserId = request.UserId,
            CarOfferId = request.CarOfferId,
            StartDate = request.StartDate,
            EndDate = request.EndDate,
            Notes = request.Notes ,
            NumOfDrivers = request.NumOfDrivers,
            TotalCost = request.TotalCost
        };

        var addedCarOrder = await _commandRepository.AddAsync(carOrderToAdd, cancellationToken);

        return new CreateCarOrderResponse()
        {
            Id = addedCarOrder.Id,
            Title = "Car order created",
            Message = $"Car order for '{addedCarOrder.CarOfferId}' was created in the database."
        };
    }
}