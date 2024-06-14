using RentalModule.Domain.Models.Ids;

namespace RentalModule.Domain.Models;


public record TimePeriodModel
{
    public TimePeriodId Id { get; init; }
    public CarOfferId CarOfferId { get; init; }
    public DateTime StartDate { get; init; }
    public DateTime EndDate { get; init; }
}