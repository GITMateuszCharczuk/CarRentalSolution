using Microsoft.AspNetCore.Mvc.ModelBinding;
using Microsoft.AspNetCore.Mvc.ModelBinding.Binders;

namespace BlogModule.Infrastructure.Binders.TagId;

public class TagIdModelBinderProvider : IModelBinderProvider
{
    public IModelBinder GetBinder(ModelBinderProviderContext context)
    {
        if (context.Metadata.ModelType == typeof(Domain.Models.Ids.TagId))
        {
            return new BinderTypeModelBinder(typeof(TagIdModelBinder));
        }
        return null;
    }
}
