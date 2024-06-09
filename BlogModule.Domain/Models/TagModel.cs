using BlogModule.Domain.Models.Ids;

namespace CarRental.Web.Models.Domain.Blog;

public record TagModel
{
    public TagId Id { get; init; }
    public string Name { get; set; }
    public Guid BlogPostId { get; set; }
}