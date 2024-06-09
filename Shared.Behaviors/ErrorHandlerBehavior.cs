using MediatR;
using Results.Domain;
using Results.Infrastructure;

namespace Shared.Behaviors;

public class ErrorHandlerBehavior<TRequest, TResponse> : IPipelineBehavior<TRequest, TResponse>
    where TRequest: IRequest<TResponse>
    where TResponse : class
{
    public async Task<TResponse> Handle(TRequest request, RequestHandlerDelegate<TResponse> next, CancellationToken cancellationToken)
    {
        IErrorResult errorMessage;
        try
        {
            return await next();
        }
        catch (TaskCanceledException e)
        {
            errorMessage = new TaskCanceledErrorResult() {
                Title = "Task canceled",
                Message = "Task was canceled."
            };
        }
        catch (Exception e)
        {
            errorMessage = new UnknownErrorResult() {
                Title = "Unknown error",
                Message = "Unknown error has occurred. Please contact administrator."
            };
        }
        var response = Activator.CreateInstance(typeof(TResponse), errorMessage) as TResponse;
        return response ?? throw new Exception("Improperly configured error handler. ");
    }
}