using System.Collections.Immutable;

namespace Results.Contract;

public interface IApiCollectionResult<TItem>
{
    public ImmutableArray<TItem> Items { get; init; }
}