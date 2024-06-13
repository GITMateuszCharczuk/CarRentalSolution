using BlogModule.Domain.Models.Ids;

namespace BlogModule.Domain.Models;

public record TagModel
{
    public TagId Id { get; init; }
    public string Name { get; set; }
    public BlogPostId BlogPostId { get; set; }
}