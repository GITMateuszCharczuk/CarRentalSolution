using System.Text.Json.Serialization;
using Shared.Utilities;

namespace BlogModule.Domain.Models.Ids;

[JsonConverter(typeof(BaseIdJsonConverter))]
public sealed class BlogPostCommentId : BaseId<BlogPostCommentId>
{
    public BlogPostCommentId(Guid value) : base(value){ }
}