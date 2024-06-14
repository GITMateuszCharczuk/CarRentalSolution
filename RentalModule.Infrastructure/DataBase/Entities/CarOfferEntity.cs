using RentalModule.Domain.Models;
using RentalModule.Domain.Models.Ids;
using Shared.CQRS;

namespace RentalModule.Infrastructure.DataBase.Entities;

public record CarOfferEntity : CarOfferModel, IEntity<CarOfferId>
{
    public CarOfferId Id { get; init; }
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
    public CarTariffEntity Tarrif { get; init; }

    // Navigation Property
    public ICollection<CarTagEntity> Tags { get; init; }
    public ICollection<ImageUrlEntity> ImageUrls { get; init; }
    public ICollection<TimePeriodEntity> UnavailableDates { get; init; }
}