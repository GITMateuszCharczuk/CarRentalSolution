using RentalModule.Application.CommandHandlers.CarOrder.DeleteCarOrder;
using RentalModule.Application.Contract.CarOffers.DeleteCarOffer;
using RentalModule.Domain.Models.Ids;
using RentalModule.Domain.RepositoryInterfaces.CarOffer;
using Results.Application;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.CommandHandlers;

namespace RentalModule.Application.CommandHandlers.CarOffer.DeleteCarOffer;

public class DeleteCarOfferCommandHandler : ICommandHandler<DeleteCarOfferCommand, HandlerResult<DeleteCarOfferResponse, IErrorResult>>
{
    private readonly ICarOfferCommandRepository _commandRepository;
    private readonly ICarOfferQueryRepository _queryRepository;

    public DeleteCarOfferCommandHandler(
        ICarOfferCommandRepository repository, 
        ICarOfferQueryRepository queryRepository)
    {
        _commandRepository = repository;
        _queryRepository = queryRepository;
    }

    public async Task<HandlerResult<DeleteCarOfferResponse, IErrorResult>> Handle(DeleteCarOfferCommand request, CancellationToken cancellationToken)
    {
        var isSuccess = Guid.TryParse(request.Id, out var carOfferId);
        var carOffer = isSuccess ? await _queryRepository.GetByIdAsync(new CarOfferId(carOfferId), cancellationToken) : null;
        
        if (carOffer is null) return new EntityNotFoundErrorResult() {
            Title = "Cannot delete car offer",
            Message = $"Car offer with ID {request.Id} was not found in the database."
        };
        
        await _commandRepository.DeleteAsync(carOffer.Id, cancellationToken);

        return new DeleteCarOfferResponse() {
            Title = "Car offer deleted",
            Message = $"Car offer with ID {request.Id} was deleted from the database."
        };
    }
}