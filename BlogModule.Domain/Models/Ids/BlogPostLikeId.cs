using System.Text.Json.Serialization;
using Shared.Utilities;

namespace BlogModule.Domain.Models.Ids;

// [JsonConverter(typeof(BaseIdJsonConverter))]
// public sealed class BlogPostLikeId : BaseId<BlogPostLikeId>
// {
//     public BlogPostLikeId(Guid value) : base(value) { }
// }

using System;
using System.Text.Json;
using System.Text.Json.Serialization;

[JsonConverter(typeof(BlogPostLikeIdJsonConverter))]
public readonly record struct BlogPostLikeId
{
    private Guid Value { get; init; }

    public BlogPostLikeId(Guid value)
    {
        Value = value;
    }
    
    public static implicit operator Guid(BlogPostLikeId id) => id.Value;
    
    public static implicit operator BlogPostLikeId(Guid id) => new(id);

    public override string ToString() => Value.ToString();

    private class BlogPostLikeIdJsonConverter : JsonConverter<BlogPostLikeId>
    {
        public override BlogPostLikeId Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
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

        public override void Write(Utf8JsonWriter writer, BlogPostLikeId value, JsonSerializerOptions options)
        {
            writer.WriteStringValue(value.ToString());
        }
    }
}
