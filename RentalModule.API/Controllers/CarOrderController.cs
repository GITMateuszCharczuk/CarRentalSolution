using MediatR;
using Microsoft.AspNetCore.Mvc;
using RentalModule.API.Mappers;
using RentalModule.Application.CommandHandlers.CarOrder.DeleteCarOrder;
using RentalModule.Application.Contract.CarOrders.CreateCarOrder;
using RentalModule.Application.Contract.CarOrders.GetCarOrders;
using RentalModule.Application.Contract.CarOrders.UpdateCarOrder;
using RentalModule.Application.QueryHandlers.CarOrder.GetCarOrder;
using RentalModule.API.Utilities;
namespace RentalModule.API.Controllers;

[ApiController]
[Route("api/[controller]")]
public class CarOrdersController : Controller
{
    private readonly IMediator _mediator;
    private readonly ICarApiMapper _mapper;

    public CarOrdersController(IMediator mediator, ICarApiMapper mapper)
    {
        _mediator = mediator;
        _mapper = mapper;
    }

    [HttpGet("{idOrUrl}")]
    public async Task<IActionResult> GetCarOrderAsync([FromRoute] string idOrUrl, CancellationToken cancellationToken) =>
        (await _mediator.Send(new GetCarOrderQuery { Id = idOrUrl }, cancellationToken)).Match(Ok, this.ErrorResult);

    [HttpGet]
    public async Task<IActionResult> GetCarOrdersAsync([FromQuery] GetCarOrdersRequest request, CancellationToken cancellationToken) =>
        (await _mediator.Send(_mapper.MapToMessage(request), cancellationToken)).Match(Ok, this.ErrorResult);

    [HttpPost]
    public async Task<IActionResult> CreateCarOrderAsync([FromBody] CreateCarOrderRequest request, CancellationToken cancellationToken) =>
        (await _mediator.Send(_mapper.MapToMessage(request), cancellationToken)).Match(Ok, this.ErrorResult);

    [HttpPut]
    public async Task<IActionResult> UpdateCarOrderAsync([FromBody] UpdateCarOrderRequest request, CancellationToken cancellationToken) =>
        (await _mediator.Send(_mapper.MapToMessage(request), cancellationToken)).Match(Ok, this.ErrorResult);

    [HttpDelete("{id}")]
    public async Task<IActionResult> DeleteCarOrderAsync([FromRoute] string id, CancellationToken cancellationToken) =>
        (await _mediator.Send(new DeleteCarOrderCommand { Id = id }, cancellationToken)).Match(Ok, this.ErrorResult);
}
