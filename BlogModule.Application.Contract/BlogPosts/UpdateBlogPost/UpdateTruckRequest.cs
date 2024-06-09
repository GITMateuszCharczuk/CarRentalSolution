using System.Text.Json.Serialization;
using TrucksModule.Domain.Enums;
using TrucksModule.Domain.Models;

namespace TrucksModule.Application.Contract.Trucks.UpdateTruck;

public record UpdateTruckRequest
{
    public TruckId? Id { get; init; }
    public string Code { get; init; } = string.Empty;
    public string Name { get; init; } = string.Empty;
    public string? Description { get; init; } = string.Empty;
    [JsonConverter(typeof(JsonStringEnumConverter))]
    public TruckStatusEnum Status { get; init; }
}