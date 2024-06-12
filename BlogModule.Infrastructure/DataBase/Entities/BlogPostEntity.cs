using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using Shared.CQRS;

namespace BlogModule.Infrastructure.DataBase.Entities;

public record BlogPostEntity : IEntity<BlogPostId>
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

    // Navigation Properties
    public ICollection<TagEntity> Tags { get; set; }
    public ICollection<BlogPostLikeEntity> Likes { get; set; }
    public ICollection<BlogPostCommentEntity> Comments { get; set; }
}