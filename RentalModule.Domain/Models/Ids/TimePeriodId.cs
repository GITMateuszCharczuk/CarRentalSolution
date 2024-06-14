using System.Text.Json;
using System.Text.Json.Serialization;

namespace RentalModule.Domain.Models.Ids;


[JsonConverter(typeof(BlogPostLikeIdJsonConverter))]
public readonly record struct TimePeriodId
{
    private Guid Value { get; init; }

    public TimePeriodId(Guid value)
    {
        Value = value;
    }
    
    public static implicit operator Guid(TimePeriodId id) => id.Value;
    
    public static implicit operator TimePeriodId(Guid id) => new(id);

    public override string ToString() => Value.ToString();

    private class BlogPostLikeIdJsonConverter : JsonConverter<TimePeriodId>
    {
        public override TimePeriodId Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
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

        public override void Write(Utf8JsonWriter writer, TimePeriodId value, JsonSerializerOptions options)
        {
            writer.WriteStringValue(value.ToString());
        }
    }
}