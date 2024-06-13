using System.Text.Json.Serialization;
using Shared.Utilities;

namespace BlogModule.Domain.Models.Ids;

// [JsonConverter(typeof(BaseIdJsonConverter))]
// public sealed class BlogPostCommentId : BaseId<BlogPostCommentId>
// {
//     public BlogPostCommentId(Guid value) : base(value){ }
// }


using System;
using System.Text.Json;
using System.Text.Json.Serialization;

[JsonConverter(typeof(BlogPostCommentIdJsonConverter))]
public readonly record struct BlogPostCommentId
{
    private Guid Value { get; init; }

    public BlogPostCommentId(Guid value)
    {
        Value = value;
    }
    
    public static implicit operator Guid(BlogPostCommentId id) => id.Value;
    
    public static implicit operator BlogPostCommentId(Guid id) => new(id);

    public override string ToString() => Value.ToString();

    private class BlogPostCommentIdJsonConverter : JsonConverter<BlogPostCommentId>
    {
        public override BlogPostCommentId Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
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

        public override void Write(Utf8JsonWriter writer, BlogPostCommentId value, JsonSerializerOptions options)
        {
            writer.WriteStringValue(value.ToString());
        }
    }
}
