// using System.Collections.Immutable;
// using System.Text.Json.Serialization;
// using TrucksModule.Application.Contract.Base;
// using TrucksModule.Domain.Enums;
// using TrucksModule.Domain.Models;
//
// namespace TrucksModule.Application.Contract.Trucks.GetBlogPosts;
//
// public record GetTrucksRequest : IPageableRequest, ISortable<TruckSortColumnEnum?>
// {
//     public int? Page { get; init; }
//     
//     public int? PageSize { get; init; }
//     
//     [JsonConverter(typeof(JsonStringEnumConverter))]
//     public TruckSortColumnEnum? OrderBy { get; init; }
//     
//     [JsonConverter(typeof(JsonStringEnumConverter))]
//     public SortOrderEnum? OrderDirection { get; init; }
//     
//     public TruckId[]? Ids { get; init; }
//     
//     public string[]? Codes { get; init; }
//     
//     public string[]? Names { get; init; }
//     
//     [JsonConverter(typeof(JsonStringEnumConverter))]
//     public TruckStatusEnum[]? Statuses { get; init; }
// }