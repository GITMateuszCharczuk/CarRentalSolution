using Results.Domain;

namespace BlogModule.Application.Contract.BlogPosts.GetBlogPost;

public record GetTruckResponse : ISuccessResult
{
    public TruckId Id { get; init; }
    public string Code { get; init; } = string.Empty;
    public string Name { get; init; } = string.Empty;
    public string? Description { get; init; } = string.Empty;
    public TruckStatusEnum Status { get; init; }
}