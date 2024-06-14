using MediatR;
using Microsoft.AspNetCore.Mvc;
using RentalModule.API.Mappers;
using RentalModule.Application.CommandHandlers.CarOffer.DeleteCarOffer;
using RentalModule.Application.Contract.CarOffers.CreateCarOffer;
using RentalModule.Application.Contract.CarOffers.GetCarOffers;
using RentalModule.Application.Contract.CarOffers.UpdateCarOffer;
using RentalModule.Application.QueryHandlers.CarOffer.GetCarOffer;
using RentalModule.API.Utilities;
namespace RentalModule.API.Controllers;

[ApiController]
[Route("api/[controller]")]
public class CarOffersController : Controller
{
    private readonly IMediator _mediator;
    private readonly ICarApiMapper _mapper;

    public CarOffersController(IMediator mediator, ICarApiMapper mapper)
    {
        _mediator = mediator;
        _mapper = mapper;
    }

    [HttpGet("{idOrUrl}")]
    public async Task<IActionResult> GetCarOfferAsync([FromRoute] string idOrUrl, CancellationToken cancellationToken) =>
        (await _mediator.Send(new GetCarOfferQuery { IdOrUrl = idOrUrl }, cancellationToken)).Match(Ok, this.ErrorResult);

    [HttpGet]
    public async Task<IActionResult> GetCarOffersAsync([FromQuery] GetCarOffersRequest request, CancellationToken cancellationToken) =>
        (await _mediator.Send(_mapper.MapToMessage(request), cancellationToken)).Match(Ok, this.ErrorResult);

    [HttpPost]
    public async Task<IActionResult> CreateCarOfferAsync([FromBody] CreateCarOfferRequest request, CancellationToken cancellationToken) =>
        (await _mediator.Send(_mapper.MapToMessage(request), cancellationToken)).Match(Ok, this.ErrorResult);

    [HttpPut]
    public async Task<IActionResult> UpdateCarOfferAsync([FromBody] UpdateCarOfferRequest request, CancellationToken cancellationToken) =>
        (await _mediator.Send(_mapper.MapToMessage(request), cancellationToken)).Match(Ok, this.ErrorResult);

    [HttpDelete("{id}")]
    public async Task<IActionResult> DeleteCarOfferAsync([FromRoute] string id, CancellationToken cancellationToken) =>
        (await _mediator.Send(new DeleteCarOfferCommand { Id = id }, cancellationToken)).Match(Ok, this.ErrorResult);
}
