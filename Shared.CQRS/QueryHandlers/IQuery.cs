using MediatR;

namespace Shared.CQRS.QueryHandlers;

public interface IQuery<out TResponse> : IRequest<TResponse>
{
}