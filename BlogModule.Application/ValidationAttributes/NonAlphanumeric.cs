using System.Collections.Immutable;
using System.ComponentModel.DataAnnotations;
using System.Linq;

namespace BlogModule.Application.ValidationAttributes;


public class NonAlphanumeric : ValidationAttribute
{
    protected override ValidationResult IsValid(object? value, ValidationContext validationContext)
    {
        var tags = value is ImmutableArray<string> ? (ImmutableArray<string>)value : default;
        if (tags == null || tags.Length == 0)
        {
            return ValidationResult.Success!;
        }

        var invalidTags = tags.Where(tag => !System.Text.RegularExpressions.Regex.IsMatch(tag, @"^[a-zA-Z0-9\s]*$"))
            .ToList();

        if (invalidTags.Any())
        {
            return new ValidationResult(
                $"The following tags contain invalid characters: {string.Join(", ", invalidTags)}");
        }

        return ValidationResult.Success!;
    }
}