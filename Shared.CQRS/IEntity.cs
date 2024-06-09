using Shared.Utilities;

namespace Shared.CQRS;

public interface IEntity<TId>
    where TId : BaseId<TId>
{
    public TId Id { get; init; }
}
