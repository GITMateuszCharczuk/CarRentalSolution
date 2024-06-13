using Microsoft.AspNetCore.Mvc.ModelBinding;
using Microsoft.AspNetCore.Mvc.ModelBinding.Binders;

namespace BlogModule.Infrastructure.Binders.BlogPostCommentId;

public class BlogPostCommentIdModelBinderProvider : IModelBinderProvider
{
    public IModelBinder GetBinder(ModelBinderProviderContext context)
    {
        if (context.Metadata.ModelType == typeof(Domain.Models.Ids.BlogPostCommentId))
        {
            return new BinderTypeModelBinder(typeof(BlogPostCommentIdModelBinder));
        }
        return null;
    }
}
