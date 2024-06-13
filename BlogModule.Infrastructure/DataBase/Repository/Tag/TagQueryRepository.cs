using System.Collections.Immutable;
using BlogModule.Domain.Enums;
using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using BlogModule.Domain.RepositoryInterfaces.Tag;
using BlogModule.Infrastructure.DataBase.Context;
using BlogModule.Infrastructure.DataBase.Entities;
using Microsoft.EntityFrameworkCore;
using Results.Contract;
using Shared.CQRS.Repository;
using Shared.Utilities;

namespace BlogModule.Infrastructure.DataBase.Repository.Tag;

public class TagQueryRepository : QueryRepository<TagEntity, TagId, TagModel, BlogDbContext>, ITagQueryRepository
{
    public TagQueryRepository(BlogDbContext dbContext, IPersistenceMapper<TagEntity, TagModel> mapper) : base(dbContext, mapper)
    {
    }
    
    public override async Task<TagModel?> GetByIdAsync(TagId id, CancellationToken cancellationToken = default) => 
        await base.GetByIdAsync(id, cancellationToken);
    
    public override async Task<ImmutableArray<TagModel>> GetByIdsAsync(ImmutableArray<TagId> ids, CancellationToken cancellationToken = default) => 
        await base.GetByIdsAsync(ids, cancellationToken);

    public override async Task<int> GetTotalCountAsync(CancellationToken cancellationToken = default) => 
        await base.GetTotalCountAsync(cancellationToken);

    public async Task<ImmutableArray<TagModel>> GetAllDistinctAsync(TagSortColumnEnum? orderBy,
        SortOrderEnum? orderDirection,BlogPostId? blogPostId, CancellationToken cancellationToken)
    {
        var queryableTags = DbContext.Tags
            .AsNoTracking()
            .AsQueryable();
        
        if (blogPostId is not null)
        {
            queryableTags = queryableTags.Where(x => x.BlogPostId == blogPostId);
        }
        
        if (orderBy.HasValue && orderDirection.HasValue)
        {
            var isOrderDirectionAscending = orderDirection == SortOrderEnum.Ascending;
            queryableTags = orderBy switch
            {
                TagSortColumnEnum.Name => isOrderDirectionAscending
                    ? queryableTags.OrderBy(x => x.Name)
                    : queryableTags.OrderByDescending(x => x.Name),
                _ => queryableTags
            };
        }

        var tags = await queryableTags.ToListAsync(cancellationToken);

        var tagModels = tags
            .Select(tag => Mapper.MapToModel(tag))
            .ToImmutableArray();

        return tagModels;
    }
}