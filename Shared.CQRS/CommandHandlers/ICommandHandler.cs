using MediatR;

namespace Shared.CQRS.CommandHandlers;

public interface ICommandHandler<in TRequest, TResponse> : IRequestHandler<TRequest, TResponse>
    where TRequest : ICommand<TResponse>
{
    
}