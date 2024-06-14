using System.Text.Json;
using System.Text.Json.Serialization;

namespace RentalModule.Domain.Models.Ids;

[JsonConverter(typeof(CarOfferIdJsonConverter))]
public readonly record struct CarOfferId
{
    private Guid Value { get; init; }

    public CarOfferId(Guid value)
    {
        Value = value;
    }
    
    public static implicit operator Guid(CarOfferId id) => id.Value;
    
    public static implicit operator CarOfferId(Guid id) => new(id);

    public override string ToString() => Value.ToString();

    private class CarOfferIdJsonConverter : JsonConverter<CarOfferId>
    {
        public override CarOfferId Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
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

        public override void Write(Utf8JsonWriter writer, CarOfferId value, JsonSerializerOptions options)
        {
            writer.WriteStringValue(value.ToString());
        }
    }
}
