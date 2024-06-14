using System.Text.Json.Serialization;
using Microsoft.AspNetCore.Mvc;
using RentalModule.Domain.Enums;
using RentalModule.Domain.Models.Ids;
using RentalModule.Infrastructure.Binders.CarOfferId;
using Results.Contract;

namespace RentalModule.Application.Contract.CarOrders.GetCarOrders;

public record GetCarOrdersRequest : IPageableRequest, ISortable<CarOrderSortColumnEnum?>
{
    public int? Page { get; init; }
    
    public int? PageSize { get; init; }
    
    [JsonConverter(typeof(JsonStringEnumConverter))]
    public CarOrderSortColumnEnum? OrderBy { get; init; }
    
    [JsonConverter(typeof(JsonStringEnumConverter))]
    public SortOrderEnum? OrderDirection { get; init; }
    
    public DateTime[]? Dates { get; init; }

    public Guid UserId { get; init; }
    [ModelBinder(BinderType = typeof(CarOfferIdModelBinder))]
    public CarOfferId CarOfferId { get; init; }

}