using System.Text.Json.Serialization;

namespace BlogModule.Application.Contract.BlogPosts.CreateBlogPost;

public record CreateTruckRequest
{
    public string Code { get; init; } = string.Empty;
    public string Name { get; init; } = string.Empty;
    public string? Description { get; init; } = string.Empty;
    [JsonConverter(typeof(JsonStringEnumConverter))]
    public TruckStatusEnum Status { get; init; }
}