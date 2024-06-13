using Blog.API.Mappers;
using Blog.API.Utilities;
using BlogModule.Application.Contract.Tags.GetTags;
using MediatR;
using Microsoft.AspNetCore.Mvc;

namespace BlogModule.API.Controllers;

[ApiController]
[Route("api/[controller]")]
public class TagsController : Controller
{
    private readonly IMediator _mediator;
    private readonly IBlogPostApiMapper _mapper;

    public TagsController(IMediator mediator, IBlogPostApiMapper mapper)
    {
        _mediator = mediator;
        _mapper = mapper;
    }

    [HttpGet("asdasdasdasdasdas")]
    public async Task<IActionResult> GetTagsAsync([FromQuery] GetTagsRequest request, CancellationToken cancellationToken) =>
        (await _mediator.Send(_mapper.MapToMessage(request), cancellationToken)).Match(Ok, this.ErrorResult);
}