using Blog.API.Mappers;
using BlogModule.API.Utilities;
using BlogModule.Application.CommandHandlers.BlogPostLike.CreateLikeForBlogPost;
using BlogModule.Application.Contract.BlogPostLikes.CreateLikeForBlogPost;
using BlogModule.Application.QueryHandlers.BlogPostLike.GetLikesCount;
using BlogModule.Application.QueryHandlers.BlogPostLike.GetLikesForBlogPost;
using MediatR;
using Microsoft.AspNetCore.Mvc;

namespace BlogModule.API.Controllers;

[ApiController]
[Route("api/[controller]")]
public class BlogPostLikeController : Controller
{
    private readonly IMediator _mediator;
    private readonly IBlogPostApiMapper _mapper;

    public BlogPostLikeController(IMediator mediator, IBlogPostApiMapper mapper)
    {
        _mediator = mediator;
        _mapper = mapper;
    }

    [HttpPost]
    public async Task<IActionResult> CreateBlogPostLikeAsync([FromBody] CreateLikeForBlogPostRequest request, CancellationToken cancellationToken) =>
        (await _mediator.Send(_mapper.MapToMessage(request), cancellationToken)).Match(Ok, this.ErrorResult);
      
    [HttpGet("getcount/{idOrUrl}")]
    public async Task<IActionResult> GetLikesCountForBlogPost([FromRoute] string idOrUrl, CancellationToken cancellationToken) =>
        (await _mediator.Send(new GetLikesCountQuery { BlogPostId = idOrUrl }, cancellationToken)).Match(Ok, this.ErrorResult);
    
    [HttpGet("{idOrUrl}")]
    public async Task<IActionResult> GetLikesForBlogPost([FromRoute] string idOrUrl, CancellationToken cancellationToken) =>
        (await _mediator.Send(new GetLikesForBlogPostQuery() { BlogPostId = idOrUrl }, cancellationToken)).Match(Ok, this.ErrorResult);
    // [HttpDelete("{id}")]
    // public async Task<IActionResult> DeleteBlogPostLikeAsync([FromRoute] BlogPostLikeId id, CancellationToken cancellationToken) =>
    //     (await _mediator.Send(new DeleteBlogPostLikeCommand() { Id = id }, cancellationToken)).Match(Ok, this.ErrorResult);
}
