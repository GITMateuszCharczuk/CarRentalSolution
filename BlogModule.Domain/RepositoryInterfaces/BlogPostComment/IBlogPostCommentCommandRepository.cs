using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;

namespace BlogModule.Domain.RepositoryInterfaces.BlogPostComment;

public interface IBlogPostCommentCommandRepository
{
    public Task<BlogPostCommentModel> AddAsync(BlogPostCommentModel model, CancellationToken cancellationToken);
    public Task<BlogPostCommentModel> UpdateAsync(BlogPostCommentModel model, CancellationToken cancellationToken);
    public Task DeleteAsync(BlogPostCommentId id, CancellationToken cancellationToken);
    //public Task DeleteAsync(string code, CancellationToken cancellationToken);
}