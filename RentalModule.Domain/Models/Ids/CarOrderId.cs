using System.Text.Json;
using System.Text.Json.Serialization;

namespace RentalModule.Domain.Models.Ids;


[JsonConverter(typeof(CarOrderIdJsonConverter))]
public readonly record struct CarOrderId
{
    private Guid Value { get; init; }

    public CarOrderId(Guid value)
    {
        Value = value;
    }
    
    public static implicit operator Guid(CarOrderId id) => id.Value;
    
    public static implicit operator CarOrderId(Guid id) => new(id);

    public override string ToString() => Value.ToString();

    private class CarOrderIdJsonConverter : JsonConverter<CarOrderId>
    {
        public override CarOrderId Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
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

        public override void Write(Utf8JsonWriter writer, CarOrderId value, JsonSerializerOptions options)
        {
            writer.WriteStringValue(value.ToString());
        }
    }
}