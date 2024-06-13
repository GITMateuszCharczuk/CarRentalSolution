using Blog.API.Mappers;
using BlogModule.Application.Contract.BlogPostComments.CreateBlogPostComment;
using BlogModule.Application.Contract.BlogPostComments.GetBlogPostComments;
using Blog.API.Utilities;
using MediatR;
using Microsoft.AspNetCore.Mvc;

namespace BlogModule.API.Controllers;

[ApiController]
[Route("api/[controller]")]
public class BlogPostCommentsController : Controller
{
    private readonly IMediator _mediator;
    private readonly IBlogPostApiMapper _mapper;

    public BlogPostCommentsController(IMediator mediator, IBlogPostApiMapper mapper)
    {
        _mediator = mediator;
        _mapper = mapper;
    }

    [HttpGet]
    public async Task<IActionResult> GetBlogPostCommentsAsync([FromQuery] GetBlogPostCommentsRequest request, CancellationToken cancellationToken) =>
        (await _mediator.Send(_mapper.MapToMessage(request), cancellationToken)).Match(Ok, this.ErrorResult);

    [HttpPost]
    public async Task<IActionResult> CreateBlogPostCommentAsync([FromBody] CreateBlogPostCommentRequest request, CancellationToken cancellationToken) =>
        (await _mediator.Send(_mapper.MapToMessage(request), cancellationToken)).Match(Ok, this.ErrorResult);
    
    // [HttpDelete("{id}")]
    // public async Task<IActionResult> DeleteBlogPostCommentAsync([FromRoute] BlogPostCommentId id, CancellationToken cancellationToken) =>
    //     (await _mediator.Send(new DeleteBlogPostCommentCommand() { Id = id }, cancellationToken)).Match(Ok, this.ErrorResult); todo
}
