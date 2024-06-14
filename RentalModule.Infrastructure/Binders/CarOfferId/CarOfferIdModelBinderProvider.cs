using Microsoft.AspNetCore.Mvc.ModelBinding;
using Microsoft.AspNetCore.Mvc.ModelBinding.Binders;

namespace RentalModule.Infrastructure.Binders.CarOfferId;

public class CarOfferIdModelBinderProvider : IModelBinderProvider
{
    public IModelBinder GetBinder(ModelBinderProviderContext context)
    {
        if (context.Metadata.ModelType == typeof(Domain.Models.Ids.CarOfferId))
        {
            return new BinderTypeModelBinder(typeof(CarOfferIdModelBinder));
        }
        return null;
    }
}
