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


    public interface ICarApiMapper
    {
        GetCarOffersQuery MapToMessage(GetCarOffersRequest request);
        CreateCarOfferCommand MapToMessage(CreateCarOfferRequest request);
        UpdateCarOfferCommand MapToMessage(UpdateCarOfferRequest request);

        GetCarOrdersQuery MapToMessage(GetCarOrdersRequest request);
        CreateCarOrderCommand MapToMessage(CreateCarOrderRequest request);
        UpdateCarOrderCommand MapToMessage(UpdateCarOrderRequest request);

        GetTagsQuery MapToMessage(GetTagsRequest request);
    }
