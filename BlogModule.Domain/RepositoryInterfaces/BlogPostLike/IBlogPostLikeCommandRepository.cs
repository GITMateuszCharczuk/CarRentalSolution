using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;

namespace BlogModule.Domain.RepositoryInterfaces.BlogPostLike;

public interface IBlogPostLikeCommandRepository
{
    public Task<BlogPostLikeModel> AddAsync(BlogPostLikeModel model, CancellationToken cancellationToken);
    public Task<BlogPostLikeModel> UpdateAsync(BlogPostLikeModel model, CancellationToken cancellationToken);
    public Task DeleteAsync(BlogPostLikeId id, CancellationToken cancellationToken);

    public Task AddLikeForBlogAsync(BlogPostLikeId blogPostId, Guid userId);
    //public Task DeleteAsync(string code, CancellationToken cancellationToken);
}