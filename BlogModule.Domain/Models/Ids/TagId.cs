using System.Text.Json.Serialization;
using Shared.Utilities;

namespace BlogModule.Domain.Models.Ids;

// [JsonConverter(typeof(BaseIdJsonConverter))]
// public sealed class TagId : BaseId<TagId>
// {
//     public TagId(Guid value) : base(value) { }
// }

using System;
using System.Text.Json;
using System.Text.Json.Serialization;

[JsonConverter(typeof(TagIdJsonConverter))]
public readonly record struct TagId
{
    private Guid Value { get; init; }

    public TagId(Guid value)
    {
        Value = value;
    }
    
    public static implicit operator Guid(TagId id) => id.Value;
    
    public static implicit operator TagId(Guid id) => new(id);

    public override string ToString() => Value.ToString();

    private class TagIdJsonConverter : JsonConverter<TagId>
    {
        public override TagId Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
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

        public override void Write(Utf8JsonWriter writer, TagId value, JsonSerializerOptions options)
        {
            writer.WriteStringValue(value.ToString());
        }
    }
}
