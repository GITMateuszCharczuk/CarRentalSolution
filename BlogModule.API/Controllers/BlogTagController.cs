using Blog.API.Mappers;
using BlogModule.API.Utilities;
using BlogModule.Application.Contract.Tags.GetTags;
using MediatR;
using Microsoft.AspNetCore.Mvc;

namespace BlogModule.API.Controllers;

[ApiController]
[Route("api/[controller]")]
public class BlogTagsController : Controller
{
    private readonly IMediator _mediator;
    private readonly IBlogPostApiMapper _mapper;

    public BlogTagsController(IMediator mediator, IBlogPostApiMapper mapper)
    {
        _mediator = mediator;
        _mapper = mapper;
    }

    [HttpGet("")]
    public async Task<IActionResult> GetTagsAsync([FromQuery] GetTagsRequest request, CancellationToken cancellationToken) =>
        (await _mediator.Send(_mapper.MapToMessage(request), cancellationToken)).Match(Ok, this.ErrorResult);

}