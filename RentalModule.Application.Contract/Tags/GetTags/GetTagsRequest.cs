using System.Text.Json.Serialization;
using BlogModule.Domain.Enums;
using Microsoft.AspNetCore.Mvc;
using RentalModule.Domain.Models.Ids;
using RentalModule.Infrastructure.Binders.CarOfferId;
using Results.Contract;

namespace RentalModule.Application.Contract.Tags.GetTags;

public class GetTagsRequest : ISortable<CarTagSortColumnEnum?>
{
    [JsonConverter(typeof(JsonStringEnumConverter))]
    public CarTagSortColumnEnum? OrderBy { get; init; } = CarTagSortColumnEnum.Name;
    [JsonConverter(typeof(JsonStringEnumConverter))]
    public SortOrderEnum? OrderDirection { get; init; } 
    [ModelBinder(BinderType = typeof(CarOfferIdModelBinder))]
    public CarOfferId? CarOfferId { get; init; } 
}