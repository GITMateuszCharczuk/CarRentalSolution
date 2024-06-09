using System.Collections.Immutable;
using BlogModule.Domain.Models.Ids;
using BlogModule.Domain.RepositoryInterfaces.BlogPost;
using BlogModule.Infrastructure.DataBase.Context;
using BlogModule.Infrastructure.DataBase.Entities;
using CarRental.Web.Models.Domain.Blog;
using Microsoft.EntityFrameworkCore;
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

    public async Task<ImmutableArray<TagModel>> GetAllDistinctAsync(CancellationToken cancellationToken) =>
        await DbContext.Tags
            .AsNoTracking()
            .GroupBy(x => x.Name.ToLower())
            .Select(x => Mapper.MapToModel(x.First()))
            .ToImmutableArrayAsync(cancellationToken);
}