using System.Collections.Immutable;
using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using CarRental.Web.Models.Domain.Blog;
using Results.Contract;
using Results.Domain;

namespace BlogModule.Application.Contract.BlogPosts.GetBlogPost;

public record GetBlogPostResponse :  ISuccessResult//IApiCollectionResult<TagModel>
{
    public BlogPostId Id { get; init; }
    public string Heading { get; set; } = string.Empty;
    public string PageTitle { get; set; } = string.Empty;
    public string Content { get; set; } = string.Empty;
    public string ShortDescription { get; set; } = string.Empty;
    public string FeaturedImageUrl { get; set; } = string.Empty;
    public string UrlHandle { get; set; } = string.Empty;
    public DateTime PublishedDate { get; set; } = DateTime.Today;
    public string Author { get; set; } = string.Empty;
    public bool Visible { get; set; }
    
    public ImmutableArray<TagModel> Tags { get; set; }
    public ImmutableArray<BlogPostLikeModel> Likes { get; set; }
    public ImmutableArray<BlogPostCommentModel> Comments { get; set; }
}