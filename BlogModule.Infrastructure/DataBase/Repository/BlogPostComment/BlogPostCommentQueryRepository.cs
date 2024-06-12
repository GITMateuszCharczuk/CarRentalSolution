using System.Collections.Immutable;
using BlogModule.Domain.Enums;
using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using BlogModule.Domain.RepositoryInterfaces.BlogPostComment;
using BlogModule.Infrastructure.DataBase.Context;
using BlogModule.Infrastructure.DataBase.Entities;
using Microsoft.EntityFrameworkCore;
using Results.Contract;
using Shared.CQRS.Repository;
using Shared.Utilities;

namespace BlogModule.Infrastructure.DataBase.Repository.BlogPostComment;

public class BlogPostCommentQueryRepository :
    QueryRepository<BlogPostCommentEntity, BlogPostCommentId, BlogPostCommentModel, BlogDbContext>,
    IBlogPostCommentQueryRepository
{
    public BlogPostCommentQueryRepository(BlogDbContext dbContext,
        IPersistenceMapper<BlogPostCommentEntity, BlogPostCommentModel> mapper) : base(dbContext, mapper)
    {
    }

    public override async Task<BlogPostCommentModel?> GetByIdAsync(BlogPostCommentId id,
        CancellationToken cancellationToken = default) =>
        await base.GetByIdAsync(id, cancellationToken);

    public override async Task<ImmutableArray<BlogPostCommentModel>> GetByIdsAsync(
        ImmutableArray<BlogPostCommentId> ids, CancellationToken cancellationToken = default) =>
        await base.GetByIdsAsync(ids, cancellationToken);

    public async Task<int> GetTotalCommentsCountAsync(BlogPostId? blogPostId, ImmutableArray<BlogPostCommentId>? ids,
        ImmutableArray<DateTime>? dateTimes,
        ImmutableArray<Guid>? userIds, CancellationToken cancellationToken)
    {
        var query = DbContext.BlogPostComments
            .AsNoTracking()
            .AsQueryable();

        if (blogPostId is not null)
        {
            query = query.Where(c => c.BlogPostId == blogPostId.Value);
        }

        if (ids.HasValue)
        {
            query = query.Where(c => ids.Value.Contains(c.Id));
        }

        if (dateTimes.HasValue)
        {
            query = query.Where(c => dateTimes.Value.Contains(c.DateAdded));
        }

        if (userIds.HasValue)
        {
            query = query.Where(c => userIds.Value.Contains(c.UserId));
        }

        return await query.CountAsync(cancellationToken);
    }

    public async Task<ImmutableArray<BlogPostCommentModel>> GetByBlogPostIdAsync(BlogPostId blogPostId,
        CancellationToken cancellationToken) =>
        await DbContext.BlogPostComments
            .AsNoTracking()
            .Where(x => x.BlogPostId == blogPostId)
            .Select(bp => Mapper.MapToModel(bp))
            .ToImmutableArrayAsync(cancellationToken);

    public async Task<ImmutableArray<BlogPostCommentModel>> GetCollectionAsync(int? page, int? pageSize,
        BlogPostId? blogPostId, BlogPostCommentSortColumnEnum? orderBy,
        SortOrderEnum? orderDirection, ImmutableArray<BlogPostCommentId>? ids, ImmutableArray<DateTime>? dateTimes,
        ImmutableArray<Guid>? userIds,
        CancellationToken cancellationToken)
    {
        var queryableComments = DbContext.BlogPostComments
            .AsNoTracking()
            .AsQueryable();

        if (blogPostId is not null)
        {
            queryableComments = queryableComments.Where(x => x.BlogPostId == blogPostId.Value);
        }

        if (ids is not null && ids.Any<BlogPostCommentId>())
        {
            queryableComments = queryableComments.Where(x => ids.Contains(x.Id));
        }

        if (dateTimes is not null && dateTimes.Any<DateTime>())
        {
            queryableComments = queryableComments.Where(x => dateTimes.Contains(x.DateAdded));
        }

        if (userIds is not null && userIds.Any<Guid>())
        {
            queryableComments = queryableComments.Where(x => userIds.Contains(x.UserId));
        }

        if (orderBy is not null && orderDirection is not null)
        {
            var isOrderDirectionAscending = orderDirection == SortOrderEnum.Ascending;
            queryableComments = orderBy switch
            {
                BlogPostCommentSortColumnEnum.DateAdded => isOrderDirectionAscending
                    ? queryableComments.OrderBy(x => x.DateAdded)
                    : queryableComments.OrderByDescending(x => x.DateAdded),
                BlogPostCommentSortColumnEnum.UserId => isOrderDirectionAscending
                    ? queryableComments.OrderBy(x => x.UserId)
                    : queryableComments.OrderByDescending(x => x.UserId),
                BlogPostCommentSortColumnEnum.Description => isOrderDirectionAscending
                    ? queryableComments.OrderBy(x => x.Description)
                    : queryableComments.OrderByDescending(x => x.Description),
                _ => queryableComments
            };
        }

        if (page is not null && pageSize is not null)
        {
            queryableComments = queryableComments
                .Skip((page.Value - 1) * pageSize.Value)
                .Take(pageSize.Value);
        }

        return await queryableComments
            .Select(x => Mapper.MapToModel(x))
            .ToImmutableArrayAsync(cancellationToken);
    }
}