using RentalModule.Domain.Models.Ids;

namespace RentalModule.Domain.Models;

public record CarTariffModel
{
    public CarTariffId Id { get; init; }
    public CarOfferId CarOfferId { get; init; }
    public Double OneNormalDayPrice { get; init; }
    public Double OneWeekendDayPrice { get; init; }
    public Double FullWeekendPrice { get; init; }
    public Double OneWeekPrice { get; init; }
    public Double OneMonthPrice { get; init; }
}