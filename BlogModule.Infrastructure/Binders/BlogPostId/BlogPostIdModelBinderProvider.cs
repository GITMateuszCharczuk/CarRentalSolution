using Microsoft.AspNetCore.Mvc.ModelBinding;
using Microsoft.AspNetCore.Mvc.ModelBinding.Binders;

namespace BlogModule.Infrastructure.Binders.BlogPostId;

public class BlogPostIdModelBinderProvider : IModelBinderProvider
{
    public IModelBinder GetBinder(ModelBinderProviderContext context)
    {
        if (context.Metadata.ModelType == typeof(Domain.Models.Ids.BlogPostId))
        {
            return new BinderTypeModelBinder(typeof(BlogPostIdModelBinder));
        }

        return null;
    }
}