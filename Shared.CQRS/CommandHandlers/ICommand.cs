using MediatR;

namespace Shared.CQRS.CommandHandlers;

public interface ICommand<out TResponse> : IRequest<TResponse>
{
}