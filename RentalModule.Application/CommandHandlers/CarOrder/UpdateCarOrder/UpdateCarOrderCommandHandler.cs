using RentalModule.Application.Contract.CarOrders.UpdateCarOrder;
using RentalModule.Domain.Models;
using RentalModule.Domain.RepositoryInterfaces.CarOrder;
using Results.Application;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.CommandHandlers;

namespace RentalModule.Application.CommandHandlers.CarOrder.UpdateCarOrder;

public class UpdateCarOrderCommandHandler : ICommandHandler<UpdateCarOrderCommand, HandlerResult<UpdateCarOrderResponse, IErrorResult>>
    {
        private readonly ICarOrderCommandRepository _commandRepository;
        private readonly ICarOrderQueryRepository _queryRepository;
        
        public UpdateCarOrderCommandHandler(
            ICarOrderCommandRepository commandRepository, 
            ICarOrderQueryRepository queryRepository)
        {
            _commandRepository = commandRepository;
            _queryRepository = queryRepository;
        }
        
        public async Task<HandlerResult<UpdateCarOrderResponse, IErrorResult>> Handle(UpdateCarOrderCommand request, CancellationToken cancellationToken)
        {
            var carOrder =  await _queryRepository.GetByIdAsync(request.Id, cancellationToken);
            
            if (carOrder is null) return new EntityNotFoundErrorResult()
            {
                Title = "Cannot update car order",
                Message = $"Car order with ID {request.Id} does not exist in the database."
            };
            
            var carOrderToUpdate = new CarOrderModel()
            {
                Id = carOrder.Id,
                UserId = request.UserId,
                CarOfferId = request.CarOfferId,
                StartDate = request.StartDate,
                EndDate = request.EndDate,
                Notes = request.Notes ?? carOrder.Notes,
                NumOfDrivers = request.NumOfDrivers,
                TotalCost = request.TotalCost
            };
            
            var updatedCarOrder = await _commandRepository.UpdateAsync(carOrderToUpdate, cancellationToken);
            
            return new UpdateCarOrderResponse()
            {
                Title = "Car order updated",
                Message = $"Car order with ID {updatedCarOrder.Id} was updated in the database."
            };
        }
    }