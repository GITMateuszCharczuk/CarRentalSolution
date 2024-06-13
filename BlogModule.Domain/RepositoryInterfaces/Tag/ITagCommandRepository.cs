using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;

namespace BlogModule.Domain.RepositoryInterfaces.Tag;

public interface ITagCommandRepository
{
    public Task<TagModel> AddAsync(TagModel model, CancellationToken cancellationToken);
    public Task<TagModel> UpdateAsync(TagModel model, CancellationToken cancellationToken);
    public Task DeleteAsync(TagId id, CancellationToken cancellationToken);
}