using System.Text.Json.Serialization;
using Shared.Utilities;

namespace BlogModule.Domain.Models.Ids;

[JsonConverter(typeof(BaseIdJsonConverter))]
public sealed class BlogPostLikeId : BaseId<BlogPostLikeId>
{
    public BlogPostLikeId(Guid value) : base(value) { }
}