using System.Collections.Immutable;
using System.Text.RegularExpressions;
using BlogModule.Application.Contract.BlogPosts.GetBlogPost;
using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using BlogModule.Domain.RepositoryInterfaces.BlogPost;
using Results.Application;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.QueryHandlers;

namespace BlogModule.Application.QueryHandlers.BlogPost.GetBlogPost;

public class GetBlogPostQueryHandler : IQueryHandler<GetBlogPostQuery, HandlerResult<GetBlogPostResponse, IErrorResult>>
{
    private readonly IBlogPostQueryRepository _repository;

    public GetBlogPostQueryHandler(IBlogPostQueryRepository repository)
    {
        _repository = repository;
    }

    public async Task<HandlerResult<GetBlogPostResponse, IErrorResult>> Handle(GetBlogPostQuery request,
        CancellationToken cancellationToken)
    {
        var isSuccess = Guid.TryParse(request.IdOrUrl, out var blogPostId);

        var blogPost = isSuccess
            ? await _repository.GetByIdAsync(new BlogPostId(blogPostId), cancellationToken)
            : await _repository.GetByUrlAsync(request.IdOrUrl, cancellationToken);


        return blogPost is null
            ? new EntityNotFoundErrorResult()
            {
                Title = "Blog post was not found",
                Message = $"Blog post with ID/Url {request.IdOrUrl} was not found in the database."
            }
            : new GetBlogPostResponse
            {
                Id = blogPost.Id,
                Heading = blogPost.Heading,
                PageTitle = blogPost.PageTitle,
                Author = blogPost.Author,
                Visible = blogPost.Visible,
                Content = blogPost.Content,
                ShortDescription = blogPost.ShortDescription,
                FeaturedImageUrl = blogPost.FeaturedImageUrl,
                UrlHandle = blogPost.UrlHandle,
                PublishedDate = blogPost.PublishedDate,
                Tags = blogPost.Tags.ToImmutableArray(),
                Likes = blogPost.Likes.ToImmutableArray(),
                Comments = blogPost.Comments.ToImmutableArray(),
            };
    }
}