using MediatR;
using Microsoft.AspNetCore.Mvc;
using RentalModule.API.Mappers;
using RentalModule.Application.Contract.Tags.GetTags;
using RentalModule.API.Utilities;
namespace RentalModule.API.Controllers;

[ApiController]
[Route("api/[controller]")]
public class CarTagsController : Controller
{
    private readonly IMediator _mediator;
    private readonly ICarApiMapper _mapper;

    public CarTagsController(IMediator mediator, ICarApiMapper mapper)
    {
        _mediator = mediator;
        _mapper = mapper;
    }

    [HttpGet]
    public async Task<IActionResult> GetTagsAsync([FromQuery] GetTagsRequest request, CancellationToken cancellationToken) =>
        (await _mediator.Send(_mapper.MapToMessage(request), cancellationToken)).Match(Ok, this.ErrorResult);
}