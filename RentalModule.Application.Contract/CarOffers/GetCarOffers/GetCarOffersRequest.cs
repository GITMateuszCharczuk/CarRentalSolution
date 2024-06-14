using System.Text.Json.Serialization;
using RentalModule.Domain.Enums;
using RentalModule.Domain.Models.Ids;
using Results.Contract;

namespace RentalModule.Application.Contract.CarOffers.GetCarOffers;

public record GetCarOffersRequest : IPageableRequest, ISortable<CarOfferSortColumnEnum?>
{
    public int? Page { get; init; }
    
    public int? PageSize { get; init; }
    
    [JsonConverter(typeof(JsonStringEnumConverter))]
    public CarOfferSortColumnEnum? OrderBy { get; init; }
    
    [JsonConverter(typeof(JsonStringEnumConverter))]
    public SortOrderEnum? OrderDirection { get; init; }
    
    public DateTime[]? PossibleDates { get; init; }
    public string[]? Tags { get; init; }
}