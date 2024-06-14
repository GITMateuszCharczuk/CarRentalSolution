using System.Collections.Immutable;
using System.ComponentModel.DataAnnotations;
using RentalModule.Application.Contract.CarOffers.CreateCarOffer;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.CommandHandlers;

namespace RentalModule.Application.CommandHandlers.CarOffer.CreateCarOffer;

public class CreateCarOfferCommand : ICommand<HandlerResult<CreateCarOfferResponse, IErrorResult>>
{
    [StringLength(200)]
    [Required]
    public string Heading { get; init; } = string.Empty;

    [StringLength(500)]
    [Required]
    public string ShortDescription { get; init; } = string.Empty;

    [Url]
    [Required]
    public string FeaturedImageUrl { get; init; } = string.Empty;

    [RegularExpression(@"^[a-zA-Z0-9\-]+$", ErrorMessage = "UrlHandle must be alphanumeric with hyphens allowed.")]
    [StringLength(100)]
    [Required]
    public string UrlHandle { get; init; } = string.Empty;

    [StringLength(50)]
    [Required]
    public string Horsepower { get; init; } = string.Empty;

    [Range(1886, 2100)]
    [Required]
    public int YearOfProduction { get; init; }

    [StringLength(20)]
    [Required]
    public string EngineDetails { get; init; } = string.Empty;

    [StringLength(10)]
    [Required]
    public string DriveDetails { get; init; } = string.Empty;

    [StringLength(10)]
    [Required]
    public string GearboxDetails { get; init; } = string.Empty;

    [StringLength(50)]
    [Required]
    public string CarDeliverylocation { get; init; } = string.Empty;

    [StringLength(50)]
    [Required]
    public string CarReturnLocation { get; init; } = string.Empty;

    [Required]
    public DateTime PublishedDate { get; init; } = DateTime.Today;

    [Required]
    public bool Visible { get; init; }

    [Range(0, double.MaxValue)]
    [Required]
    public double OneNormalDayPrice { get; init; }

    [Range(0, double.MaxValue)]
    [Required]
    public double OneWeekendDayPrice { get; init; }

    [Range(0, double.MaxValue)]
    [Required]
    public double FullWeekendPrice { get; init; }

    [Range(0, double.MaxValue)]
    [Required]
    public double OneWeekPrice { get; init; }

    [Range(0, double.MaxValue)]
    [Required]
    public double OneMonthPrice { get; init; }
    
    public ImmutableArray<string>? Tags { get; init; }

    //[Urls]todo
    public ImmutableArray<string>? ImageUrls { get; init; }

}