using System.ComponentModel.DataAnnotations;
using BlogModule.Application.Contract.BlogPosts.CreateBlogPost;
using BlogModule.Application.ValidationAttributes;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.CommandHandlers;

namespace BlogModule.Application.CommandHandlers.BlogPost.CreateBlogPost;

public class CreateBlogPostCommand : ICommand<HandlerResult<CreateBlogPostResponse, IErrorResult>>
{
    [StringLength(200)]
    [Required]
    public string Heading { get; set; } = string.Empty;

    [StringLength(200)]
    [Required]
    public string PageTitle { get; set; } = string.Empty;

    [Required]
    public string Content { get; set; } = string.Empty;

    [StringLength(500)]
    [Required]
    public string ShortDescription { get; set; } = string.Empty;

    [Url]
    [Required]
    public string FeaturedImageUrl { get; set; } = string.Empty;

    [RegularExpression(@"^[a-zA-Z0-9\-]+$", ErrorMessage = "UrlHandle must be alphanumeric with hyphens allowed.")]
    [StringLength(100)]
    [Required]
    public string UrlHandle { get; set; } = string.Empty;

    [Required]
    public DateTime PublishedDate { get; set; } = DateTime.Today;

    [StringLength(100)]
    [Required]
    public string Author { get; set; } = string.Empty;

    [Required]
    public bool Visible { get; set; }
    
    [NonAlphanumeric]
    public string[]? Tags { get; set; }
}