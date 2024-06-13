using Microsoft.AspNetCore.Mvc.ModelBinding;
using Microsoft.AspNetCore.Mvc.ModelBinding.Binders;

namespace BlogModule.Infrastructure.Binders.BlogPostLikeId;

public class BlogPostLikeIdModelBinderProvider : IModelBinderProvider
{
    public IModelBinder GetBinder(ModelBinderProviderContext context)
    {
        if (context.Metadata.ModelType == typeof(Domain.Models.Ids.BlogPostLikeId))
        {
            return new BinderTypeModelBinder(typeof(BlogPostLikeIdModelBinder));
        }
        return null;
    }
}
