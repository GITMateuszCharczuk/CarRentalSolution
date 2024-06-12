using BlogModule.Application.Contract.BlogPosts.CreateBlogPost;
using BlogModule.Domain.Models;
using BlogModule.Domain.RepositoryInterfaces.BlogPost;
using Results.Application;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.CommandHandlers;

namespace BlogModule.Application.CommandHandlers.BlogPost.CreateBlogPost;

public class CreateBlogPostCommandHandler : ICommandHandler<CreateBlogPostCommand, HandlerResult<CreateBlogPostResponse, IErrorResult>>
{
    private readonly IBlogPostCommandRepository _commandRepository;
    private readonly IBlogPostQueryRepository _queryRepository;
    
    public CreateBlogPostCommandHandler(
        IBlogPostCommandRepository repository, 
        IBlogPostQueryRepository queryRepository)
    {
        _commandRepository = repository;
        _queryRepository = queryRepository;
    }
    
    public async Task<HandlerResult<CreateBlogPostResponse, IErrorResult>> Handle(CreateBlogPostCommand request, CancellationToken cancellationToken)
    {
        var blogPost = await _queryRepository.GetByUrlAsync(request.UrlHandle, cancellationToken);
        if (blogPost is not null) return new EntityAlreadyExistsErrorResult() {
            Title = "Cannot create blog post",
            Message = $"Blog post with url handle {request.UrlHandle} already exists in the database."
        };
        
        var blogPostToAdd = new BlogPostModel
        {
            Heading = request.Heading,
            PageTitle = request.PageTitle,
            Content = request.Content,
            ShortDescription = request.ShortDescription,
            FeaturedImageUrl = request.FeaturedImageUrl,
            UrlHandle = request.UrlHandle,
            PublishedDate = request.PublishedDate,
            Author = request.Author,
            Visible = request.Visible,
            Tags = request.Tags,
            Likes = blogPost.Likes,
            Comments = blogPost.Comments
        };

        var addedBlogPost = await _commandRepository.AddAsync(blogPostToAdd, cancellationToken);

        return new CreateBlogPostResponse() {
            Id = addedBlogPost.Id,
            Title = "Blog post created",
            Message = $"Blog post '{addedBlogPost.Heading}' was created in the database."
        };
    }
}