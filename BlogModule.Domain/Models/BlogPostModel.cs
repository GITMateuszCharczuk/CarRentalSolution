using BlogModule.Domain.Models.Ids;
using CarRental.Web.Models.Domain.Blog;

namespace BlogModule.Domain.Models;

public record BlogPostModel
{
    public BlogPostId Id { get; init; }
    public string Heading { get; set; }
    public string PageTitle { get; set; }
    public string Content { get; set; }
    public string ShortDescription { get; set; }
    public string FeaturedImageUrl { get; set; }
    public string UrlHandle { get; set; }
    public DateTime PublishedDate { get; set; }
    public string Author { get; set; }

    public bool Visible { get; set; }

    // Navigation Property
    public ICollection<TagModel> Tags { get; set; }
    public ICollection<BlogPostLikeModel> Likes { get; set; }
    public ICollection<BlogPostCommentModel> Comments { get; set; }
}