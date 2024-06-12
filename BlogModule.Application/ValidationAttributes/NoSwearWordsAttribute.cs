using System.ComponentModel.DataAnnotations;
using DotnetBadWordDetector;


namespace BlogModule.Application.ValidationAttributes;

public class NoSwearWordsAttribute : ValidationAttribute
{
    protected override ValidationResult IsValid(object? value, ValidationContext validationContext)
    {
        if (value is string description)
        {
            var detector = new ProfanityDetector();
            if (detector.IsProfane(description))
            {
                return new ValidationResult("The description contains inappropriate language.");
            }
        }

        return ValidationResult.Success!;
    }
}