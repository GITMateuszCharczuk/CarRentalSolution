using BlogModule.Domain.Models.Ids;
using CarRental.Web.Models.Domain.Blog;

namespace BlogModule.Domain.RepositoryInterfaces.Tag;

public interface ITagCommandRepository
{
    public Task<TagModel> AddAsync(TagModel model, CancellationToken cancellationToken);
    public Task<TagModel> UpdateAsync(TagModel model, CancellationToken cancellationToken);
    public Task DeleteAsync(TagId id, CancellationToken cancellationToken);
    //public Task DeleteAsync(string code, CancellationToken cancellationToken);
}