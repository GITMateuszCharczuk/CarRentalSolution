using Microsoft.AspNetCore.Mvc.ModelBinding;

namespace BlogModule.Infrastructure.Binders.TagId;

public class TagIdModelBinder : IModelBinder
{
    public Task BindModelAsync(ModelBindingContext bindingContext)
    {
        if (bindingContext == null)
        {
            throw new ArgumentNullException(nameof(bindingContext));
        }

        var valueProviderResult = bindingContext.ValueProvider.GetValue(bindingContext.ModelName);

        if (valueProviderResult == ValueProviderResult.None)
        {
            return Task.CompletedTask;
        }

        bindingContext.ModelState.SetModelValue(bindingContext.ModelName, valueProviderResult);

        var value = valueProviderResult.FirstValue;

        if (string.IsNullOrEmpty(value))
        {
            return Task.CompletedTask;
        }

        if (Guid.TryParse(value, out var guid))
        {
            bindingContext.Result = ModelBindingResult.Success(new Domain.Models.Ids.TagId(guid));
        }
        else
        {
            bindingContext.ModelState.TryAddModelError(bindingContext.ModelName, "Invalid GUID format.");
        }

        return Task.CompletedTask;
    }
}
