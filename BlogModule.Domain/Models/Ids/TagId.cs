using System.Text.Json.Serialization;
using Shared.Utilities;

namespace BlogModule.Domain.Models.Ids;

[JsonConverter(typeof(BaseIdJsonConverter))]
public sealed class TagId : BaseId<TagId>
{
    public TagId(Guid value) : base(value) { }
}