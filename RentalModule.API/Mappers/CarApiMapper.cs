using System.Collections.Immutable;
using RentalModule.Application.CommandHandlers.CarOffer.CreateCarOffer;
using RentalModule.Application.CommandHandlers.CarOffer.UpdateCarOffer;
using RentalModule.Application.CommandHandlers.CarOrder.CreateCarOrder;
using RentalModule.Application.CommandHandlers.CarOrder.UpdateCarOrder;
using RentalModule.Application.Contract.CarOffers.CreateCarOffer;
using RentalModule.Application.Contract.CarOffers.GetCarOffers;
using RentalModule.Application.Contract.CarOffers.UpdateCarOffer;
using RentalModule.Application.Contract.CarOrders.CreateCarOrder;
using RentalModule.Application.Contract.CarOrders.GetCarOrders;
using RentalModule.Application.Contract.CarOrders.UpdateCarOrder;
using RentalModule.Application.Contract.Tags.GetTags;
using RentalModule.Application.QueryHandlers.CarOffer.GetCarOffers;
using RentalModule.Application.QueryHandlers.CarOrder.GetCarOrders;
using RentalModule.Application.QueryHandlers.Tag.GetTags;

namespace RentalModule.API.Mappers;

public class CarApiMapper : ICarApiMapper
{
    public GetCarOffersQuery MapToMessage(GetCarOffersRequest request) => new()
    {
        Page = request.Page,
        PageSize = request.PageSize,
        OrderBy = request.OrderBy,
        OrderDirection = request.OrderDirection,
        PossibleDates = request.PossibleDates is null ? null : ImmutableArray.Create(request.PossibleDates),
        Tags = request.Tags is null ? null : ImmutableArray.Create(request.Tags)
    };

    public CreateCarOfferCommand MapToMessage(CreateCarOfferRequest request) => new()
    {
        Heading = request.Heading,
        ShortDescription = request.ShortDescription,
        FeaturedImageUrl = request.FeaturedImageUrl,
        UrlHandle = request.UrlHandle,
        Horsepower = request.Horsepower,
        YearOfProduction = request.YearOfProduction,
        EngineDetails = request.EngineDetails,
        DriveDetails = request.DriveDetails,
        GearboxDetails = request.GearboxDetails,
        CarDeliverylocation = request.CarDeliverylocation,
        CarReturnLocation = request.CarReturnLocation,
        PublishedDate = request.PublishedDate,
        Visible = request.Visible,
        OneNormalDayPrice = request.OneNormalDayPrice,
        OneWeekendDayPrice = request.OneWeekendDayPrice,
        FullWeekendPrice = request.FullWeekendPrice,
        OneWeekPrice = request.OneWeekPrice,
        OneMonthPrice = request.OneMonthPrice,
        Tags = request.Tags is null ? null : ImmutableArray.Create(request.Tags),
        ImageUrls = request.ImageUrls is null ? null : ImmutableArray.Create(request.ImageUrls)
    };

    public UpdateCarOfferCommand MapToMessage(UpdateCarOfferRequest request) => new()
    {
        Id = request.Id,
        Heading = request.Heading,
        ShortDescription = request.ShortDescription,
        FeaturedImageUrl = request.FeaturedImageUrl,
        UrlHandle = request.UrlHandle,
        Horsepower = request.Horsepower,
        YearOfProduction = request.YearOfProduction,
        EngineDetails = request.EngineDetails,
        DriveDetails = request.DriveDetails,
        GearboxDetails = request.GearboxDetails,
        CarDeliverylocation = request.CarDeliverylocation,
        CarReturnLocation = request.CarReturnLocation,
        PublishedDate = request.PublishedDate,
        Visible = request.Visible,
        OneNormalDayPrice = request.OneNormalDayPrice,
        OneWeekendDayPrice = request.OneWeekendDayPrice,
        FullWeekendPrice = request.FullWeekendPrice,
        OneWeekPrice = request.OneWeekPrice,
        OneMonthPrice = request.OneMonthPrice,
        Tags = request.Tags is null ? null : ImmutableArray.Create(request.Tags),
        ImageUrls = request.ImageUrls is null ? null : ImmutableArray.Create(request.ImageUrls)
    };
    
    public GetCarOrdersQuery MapToMessage(GetCarOrdersRequest request) => new()
    {
        Page = request.Page,
        PageSize = request.PageSize,
        OrderBy = request.OrderBy,
        OrderDirection = request.OrderDirection,
        Dates = request.Dates is null ? null : ImmutableArray.Create(request.Dates),
        UserId = request.UserId,
        CarOfferId = request.CarOfferId
    };

    public CreateCarOrderCommand MapToMessage(CreateCarOrderRequest request) => new()
    {
        UserId = request.UserId,
        CarOfferId = request.CarOfferId,
        StartDate = request.StartDate,
        EndDate = request.EndDate,
        Notes = request.Notes,
        NumOfDrivers = request.NumOfDrivers,
        TotalCost = request.TotalCost
    };

    public UpdateCarOrderCommand MapToMessage(UpdateCarOrderRequest request) => new()
    {
        Id = request.Id,
        UserId = request.UserId,
        CarOfferId = request.CarOfferId,
        StartDate = request.StartDate,
        EndDate = request.EndDate,
        Notes = request.Notes,
        NumOfDrivers = request.NumOfDrivers,
        TotalCost = request.TotalCost
    };

    
    public GetTagsQuery MapToMessage(GetTagsRequest request) => new()
    {
        OrderBy = request.OrderBy,
        OrderDirection = request.OrderDirection,
        CarOfferId = request.CarOfferId
    };
}
