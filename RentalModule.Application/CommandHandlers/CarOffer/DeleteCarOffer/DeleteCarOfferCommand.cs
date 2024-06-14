using RentalModule.Application.Contract.CarOffers.DeleteCarOffer;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.CommandHandlers;

namespace RentalModule.Application.CommandHandlers.CarOffer.DeleteCarOffer;

public class DeleteCarOfferCommand : ICommand<HandlerResult<DeleteCarOfferResponse, IErrorResult>>
{
    public string Id { get; init; } = string.Empty;
}