using System.Text.Json;
using System.Text.Json.Serialization;

namespace RentalModule.Domain.Models.Ids;

[JsonConverter(typeof(CarTariffIdJsonConverter))]
public readonly record struct CarTariffId
{
    private Guid Value { get; init; }

    public CarTariffId(Guid value)
    {
        Value = value;
    }
    
    public static implicit operator Guid(CarTariffId id) => id.Value;
    
    public static implicit operator CarTariffId(Guid id) => new(id);

    public override string ToString() => Value.ToString();

    private class CarTariffIdJsonConverter : JsonConverter<CarTariffId>
    {
        public override CarTariffId Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
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

        public override void Write(Utf8JsonWriter writer, CarTariffId value, JsonSerializerOptions options)
        {
            writer.WriteStringValue(value.ToString());
        }
    }
}
