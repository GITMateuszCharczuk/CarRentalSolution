using System.ComponentModel.DataAnnotations;
using MediatR;
using Results.Application;

namespace Shared.Behaviors;

    public class ValidationHandlerBehaviour<TRequest, TResponse> : IPipelineBehavior<TRequest, TResponse>
        where TRequest : IRequest<TResponse>
        where TResponse : class
    {
        public async Task<TResponse> Handle(TRequest request, RequestHandlerDelegate<TResponse> next, CancellationToken cancellationToken)
        {
            var validationContext = new ValidationContext(request);
            var validationResults = new List<ValidationResult>();
            var isValid = Validator.TryValidateObject(request, validationContext, validationResults, validateAllProperties: true);

            if (isValid)
            {
                return await next();
            }

            var errorResult = new ValidationErrorResult()
            {
                Title = "Validation error",
                Message = "Validation error has occurred. Please check the request and try again.",
                Errors = validationResults.Select(x => new ValidationErrorResult.ValidationError()
                {
                    PropertyName = x.MemberNames.FirstOrDefault() ?? string.Empty,
                    ErrorMessage = x.ErrorMessage ?? string.Empty
                }).ToArray()
            };

            var response = Activator.CreateInstance(typeof(TResponse), errorResult) as TResponse;
            return response ?? throw new Exception("Improperly configured error handler.");
        }
    }
