using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;

namespace BlogModule.Domain.RepositoryInterfaces.BlogPost;

public interface IBlogPostCommandRepository
{
    public Task<BlogPostModel> AddAsync(BlogPostModel model, CancellationToken cancellationToken);
    public Task<BlogPostModel> UpdateAsync(BlogPostModel model, CancellationToken cancellationToken);
    public Task DeleteAsync(BlogPostId id, CancellationToken cancellationToken);
    //public Task DeleteAsync(string code, CancellationToken cancellationToken);
}