using RentalModule.Domain.Models;
using RentalModule.Domain.Models.Ids;

namespace RentalModule.Application.Contract.CarOffers.UpdateCarOffer;

public record UpdateCarOfferRequest
{
    public CarOfferId? Id { get; init; }
    public string Heading { get; init; } = string.Empty;
    public string ShortDescription { get; init; }= string.Empty;
    public string FeaturedImageUrl { get; init; }= string.Empty;
    public string UrlHandle { get; init; }= string.Empty;
    public string Horsepower { get; init; }= string.Empty;
    public int YearOfProduction { get; init; }
    public string EngineDetails { get; init; }= string.Empty;
    public string DriveDetails { get; init; }= string.Empty;
    public string GearboxDetails { get; init; }= string.Empty;
    public string CarDeliverylocation { get; init; }= string.Empty;
    public string CarReturnLocation { get; init; }= string.Empty;
    public DateTime PublishedDate { get; init; }
    public bool Visible { get; init; }
    public Double OneNormalDayPrice { get; init; }
    public Double OneWeekendDayPrice { get; init; }
    public Double FullWeekendPrice { get; init; }
    public Double OneWeekPrice { get; init; }
    public Double OneMonthPrice { get; init; }
    public string[]? Tags { get; init; }
    public string[]? ImageUrls { get; init; }
    //public string[]? UnavailableDates { get; init; }
}