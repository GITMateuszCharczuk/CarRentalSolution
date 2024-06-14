using System.Text.Json;
using System.Text.Json.Serialization;

namespace RentalModule.Domain.Models.Ids;

// [JsonConverter(typeof(BaseIdJsonConverter))]
// public sealed class CarTagId : BaseId<CarTagId>
// {
//     public CarTagId(Guid value) : base(value) { }
// }

[JsonConverter(typeof(CarTagIdJsonConverter))]
public readonly record struct CarTagId
{
    private Guid Value { get; init; }

    public CarTagId(Guid value)
    {
        Value = value;
    }
    
    public static implicit operator Guid(CarTagId id) => id.Value;
    
    public static implicit operator CarTagId(Guid id) => new(id);

    public override string ToString() => Value.ToString();

    private class CarTagIdJsonConverter : JsonConverter<CarTagId>
    {
        public override CarTagId Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
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

        public override void Write(Utf8JsonWriter writer, CarTagId value, JsonSerializerOptions options)
        {
            writer.WriteStringValue(value.ToString());
        }
    }
}
