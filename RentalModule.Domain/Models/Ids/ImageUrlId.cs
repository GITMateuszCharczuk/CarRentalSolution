using System.Text.Json;
using System.Text.Json.Serialization;

namespace RentalModule.Domain.Models.Ids;

[JsonConverter(typeof(ImageUrlIdJsonConverter))]
public readonly record struct ImageUrlId
{
    private Guid Value { get; init; }

    public ImageUrlId(Guid value)
    {
        Value = value;
    }
    
    public static implicit operator Guid(ImageUrlId id) => id.Value;
    
    public static implicit operator ImageUrlId(Guid id) => new(id);

    public override string ToString() => Value.ToString();

    private class ImageUrlIdJsonConverter : JsonConverter<ImageUrlId>
    {
        public override ImageUrlId Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
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

        public override void Write(Utf8JsonWriter writer, ImageUrlId value, JsonSerializerOptions options)
        {
            writer.WriteStringValue(value.ToString());
        }
    }
}