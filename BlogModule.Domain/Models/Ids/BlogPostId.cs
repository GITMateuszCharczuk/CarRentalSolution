using System.Text.Json;
using System.Text.Json.Serialization;
using Shared.Utilities;

namespace BlogModule.Domain.Models.Ids;

// [JsonConverter(typeof(BaseIdJsonConverter))]
// public sealed class BlogPostId : BaseId<BlogPostId>
// {
//     public BlogPostId(Guid value) : base(value) { }
// }


[JsonConverter(typeof(BlogPostIdJsonConverter))]
public readonly record struct BlogPostId
{
    private Guid Value { get; init; }

    public BlogPostId(Guid value)
    {
        Value = value;
    }
    
    public static implicit operator Guid(BlogPostId id) => id.Value;
    
    public static implicit operator BlogPostId(Guid id) => new(id);

    public override string ToString() => Value.ToString();

    private class BlogPostIdJsonConverter : JsonConverter<BlogPostId>
    {
        public override BlogPostId Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
        {
            string? value;
            try
            {
                value = reader.GetString();
            }
            catch
            {
                value = string.Empty;
            }
            _ = Guid.TryParse(value, out var guid);
            return guid;
        }

        public override void Write(Utf8JsonWriter writer, BlogPostId value, JsonSerializerOptions options)
        {
            writer.WriteStringValue(value.ToString());
        }
    }
}