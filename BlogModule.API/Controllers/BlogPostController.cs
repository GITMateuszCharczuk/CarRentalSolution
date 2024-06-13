using Blog.API.Mappers;
using Blog.API.Utilities;
using BlogModule.Application.CommandHandlers.BlogPost.DeleteBlogPost;
using BlogModule.Application.Contract.BlogPosts.CreateBlogPost;
using BlogModule.Application.Contract.BlogPosts.GetBlogPosts;
using BlogModule.Application.Contract.BlogPosts.UpdateBlogPost;
using BlogModule.Application.QueryHandlers.BlogPost.GetBlogPost;
using MediatR;
using Microsoft.AspNetCore.Mvc;

namespace BlogModule.API.Controllers;

[ApiController]
[Route("api/[controller]")]
public class BlogPostsController : Controller
{
    private readonly IMediator _mediator;
    private readonly IBlogPostApiMapper _mapper;

    public BlogPostsController(IMediator mediator, IBlogPostApiMapper mapper)
    {
        _mediator = mediator;
        _mapper = mapper;
    }

    [HttpGet("{idOrUrl}")]
    public async Task<IActionResult> GetBlogPostAsync([FromRoute] string idOrUrl, CancellationToken cancellationToken) =>
        (await _mediator.Send(new GetBlogPostQuery() { IdOrUrl = idOrUrl }, cancellationToken)).Match(Ok, this.ErrorResult);
    
    [HttpGet]
    public async Task<IActionResult> GetBlogPostsAsync([FromQuery] GetBlogPostsRequest request, CancellationToken cancellationToken) =>
        (await _mediator.Send(_mapper.MapToMessage(request), cancellationToken)).Match(Ok, this.ErrorResult);

    [HttpPost]
    public async Task<IActionResult> CreateBlogPostAsync([FromBody] CreateBlogPostRequest request, CancellationToken cancellationToken) =>
        (await _mediator.Send(_mapper.MapToMessage(request), cancellationToken)).Match(Ok, this.ErrorResult);
    
    [HttpPut]
    public async Task<IActionResult> UpdateBlogPostAsync([FromBody] UpdateBlogPostRequest request, CancellationToken cancellationToken) =>
        (await _mediator.Send(_mapper.MapToMessage(request), cancellationToken)).Match(Ok, this.ErrorResult);
    
    [HttpDelete("{id}")]
    public async Task<IActionResult> DeleteBlogPostAsync([FromRoute] string id, CancellationToken cancellationToken) =>
        (await _mediator.Send(new DeleteBlogPostCommand() { Id = id }, cancellationToken)).Match(Ok, this.ErrorResult);
}
