using System;
using System.Text.Json;
using System.Text.Json.Serialization;

namespace Shared.Utilities
{
    // public abstract class BaseId<T> where T : BaseId<T>
    // {
    //     public Guid Value { get; init; }
    //
    //     protected BaseId(Guid value)
    //     {
    //         Value = value;
    //     }
    //
    //     public override string ToString() => Value.ToString();
    //
    //     public override bool Equals(object? obj)
    //     {
    //         if (obj is BaseId<T> other)
    //             return Value.Equals(other.Value);
    //         return false;
    //     }
    //
    //     public override int GetHashCode() => Value.GetHashCode();
    //
    //     public static implicit operator Guid(BaseId<T> id) => id.Value;
    //     public static implicit operator BaseId<T>(Guid id) => (T)Activator.CreateInstance(typeof(T), id)!;
    //
    //     protected class BaseIdJsonConverter : JsonConverter<T>
    //     {
    //         public override T Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
    //         {
    //             var value = reader.GetString();
    //             if (Guid.TryParse(value, out var guid))
    //             {
    //                 return (T)Activator.CreateInstance(typeof(T), guid)!;
    //             }
    //             throw new JsonException("Invalid GUID format.");
    //         }
    //
    //         public override void Write(Utf8JsonWriter writer, T value, JsonSerializerOptions options)
    //         {
    //             writer.WriteStringValue(value.ToString());
    //         }
    //     }
    // }
}
