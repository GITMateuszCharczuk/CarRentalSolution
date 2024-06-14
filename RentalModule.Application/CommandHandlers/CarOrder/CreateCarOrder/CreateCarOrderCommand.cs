using System.Collections.Immutable;
using System.ComponentModel.DataAnnotations;
using RentalModule.Application.Contract.CarOrders.CreateCarOrder;
using RentalModule.Domain.Models.Ids;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.CommandHandlers;

namespace RentalModule.Application.CommandHandlers.CarOrder.CreateCarOrder;

public class CreateCarOrderCommand : ICommand<HandlerResult<CreateCarOrderResponse, IErrorResult>>
{
    [Required]
    public Guid UserId { get; init; }

    [Required]
    public CarOfferId CarOfferId { get; init; }

    [Required]
    public DateTime StartDate { get; init; }

    [Required]
    public DateTime EndDate { get; init; }

    [StringLength(1000)]
    public string? Notes { get; init; } = string.Empty;

    [Range(1, int.MaxValue, ErrorMessage = "Number of drivers must be at least 1.")]
    public int NumOfDrivers { get; init; }

    [Range(0, double.MaxValue, ErrorMessage = "Total cost must be a positive value.")]
    [Required]
    public double TotalCost { get; set; }
}