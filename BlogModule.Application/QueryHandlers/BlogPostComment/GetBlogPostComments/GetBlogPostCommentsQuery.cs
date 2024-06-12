using System.Collections.Immutable;
using BlogModule.Application.Contract.BlogPostComments.GetBlogPostComments;
using BlogModule.Domain.Enums;
using BlogModule.Domain.Models.Ids;
using Results.Contract;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.QueryHandlers;

namespace BlogModule.Application.QueryHandlers.BlogPostComment.GetBlogPostComments;

public class GetBlogPostCommentsQuery : IQuery<HandlerResult<GetBlogPostCommentsResponse, IErrorResult>>
{
    public int? Page { get; init; }
    public int? PageSize { get; init; }
    public BlogPostCommentSortColumnEnum? OrderBy { get; init; }
    public SortOrderEnum? OrderDirection { get; init; }
    public BlogPostId? BlogPostId { get; init; }
    public ImmutableArray<BlogPostCommentId>? Ids{ get; init; }
    public ImmutableArray<DateTime>? DateTimes{ get; init; }
    public ImmutableArray<Guid>? UserIds{ get; init; }
}