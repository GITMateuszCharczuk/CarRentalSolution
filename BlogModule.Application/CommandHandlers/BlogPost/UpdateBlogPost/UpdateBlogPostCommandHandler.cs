using BlogModule.Application.CommandHandlers.UpdateBlogPost;
using BlogModule.Application.Contract.BlogPosts.UpdateBlogPost;
using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using BlogModule.Domain.RepositoryInterfaces.BlogPost;
using CarRental.Web.Models.Domain.Blog;
using Results.Application;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.CommandHandlers;

namespace BlogModule.Application.CommandHandlers.BlogPost.UpdateBlogPost;

public class UpdateBlogPostCommandHandler : ICommandHandler<UpdateBlogPostCommand, HandlerResult<UpdateBlogPostResponse, IErrorResult>>
{
    private readonly IBlogPostCommandRepository _commandRepository;
    private readonly IBlogPostQueryRepository _queryRepository;
    
    public UpdateBlogPostCommandHandler(
        IBlogPostCommandRepository repository, 
        IBlogPostQueryRepository queryRepository)
    {
        _commandRepository = repository;
        _queryRepository = queryRepository;
    }
    
    public async Task<HandlerResult<UpdateBlogPostResponse, IErrorResult>> Handle(UpdateBlogPostCommand request, CancellationToken cancellationToken)
    {
        var blogPost = request.Id is null
            ? await _queryRepository.GetByUrlAsync(request.UrlHandle, cancellationToken) 
            : await _queryRepository.GetByIdAsync(new BlogPostId(request.Id.Value), cancellationToken);
        if (blogPost is null) return new EntityNotFoundErrorResult() {
            Title = "Cannot update blog post",
            Message = $"Blog post with url {request.UrlHandle} does not exist in the database."
        };
        
        var blogPostToUpdate = new BlogPostModel() {
            Id = blogPost.Id,
            Heading = request.Heading,
            PageTitle = request.PageTitle,
            Content = request.Content,
            ShortDescription = request.ShortDescription,
            FeaturedImageUrl = request.FeaturedImageUrl,
            UrlHandle = request.UrlHandle,
            PublishedDate = request.PublishedDate,
            Author = request.Author,
            Visible = request.Visible,
            Tags = ConvertToTagModels(request.Tags, blogPost.Id),
            Likes = blogPost.Likes,
            Comments = blogPost.Comments
        };
        
        var updatedTruck = await _commandRepository.UpdateAsync(blogPostToUpdate, cancellationToken);
        
        return new UpdateBlogPostResponse() {
            Title = "Blog post updated",
            Message = $"Truck with url {updatedTruck.UrlHandle} was updated in the database."
        };
    }
    
    private static ICollection<TagModel> ConvertToTagModels(string[]? tags, Guid blogPostId)
    {
        if (tags == null) return new List<TagModel>();

        return tags.Select(tag => new TagModel
        {
            Id = new TagId(Guid.NewGuid()),
            Name = tag,
            BlogPostId = blogPostId
        }).ToList();
    }
}