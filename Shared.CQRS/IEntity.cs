using Shared.Utilities;

namespace Shared.CQRS;

public interface IEntity<TId>
    where TId : struct//BaseId<TId>
{
    public TId Id { get; init; }
}
