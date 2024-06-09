using System.Text.Json.Serialization;
using Shared.Utilities;

namespace BlogModule.Domain.Models.Ids;

[JsonConverter(typeof(BaseIdJsonConverter))]
public sealed class BlogPostId : BaseId<BlogPostId>
{
    public BlogPostId(Guid value) : base(value) { }
}