using System.Collections.Immutable;
using Microsoft.EntityFrameworkCore;

namespace Shared.Utilities
{
    public static class QueryableExtension
    {
        public static async Task<ImmutableArray<TSource>> ToImmutableArrayAsync<TSource>(
            this IQueryable<TSource> source, CancellationToken cancellationToken = default) => 
            (await source.ToListAsync(cancellationToken).ConfigureAwait(false)).ToImmutableArray();
    }
}