using System.Text.Json.Serialization;

namespace BlogModule.Application.Contract.BlogPosts.UpdateBlogPost;

public record UpdateBlogPostRequest
{
    public string Heading { get; set; } = string.Empty;
    public string PageTitle { get; set; } = string.Empty;
    public string Content { get; set; } = string.Empty;
    public string ShortDescription { get; set; } = string.Empty;
    public string FeaturedImageUrl { get; set; } = string.Empty;
    public string UrlHandle { get; set; } = string.Empty;
    public DateTime PublishedDate { get; set; } = DateTime.Today;
    public string Author { get; set; } = string.Empty;
    public bool Visible { get; set; }
    public string[]? Tags { get; set; }
}