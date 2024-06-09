using BlogModule.Domain.Models.Ids;
using BlogModule.Domain.RepositoryInterfaces.Tag;
using BlogModule.Infrastructure.DataBase.Context;
using BlogModule.Infrastructure.DataBase.Entities;
using CarRental.Web.Models.Domain.Blog;
using Shared.CQRS.Repository;
using Shared.Utilities;

namespace BlogModule.Infrastructure.DataBase.Repository.Tag;

public class TagCommandRepository : CommandRepository<TagEntity, TagId, TagModel, BlogDbContext>, ITagCommandRepository
{
    public TagCommandRepository(BlogDbContext dbContext, IPersistenceMapper<TagEntity, TagModel> mapper) : base(dbContext, mapper)
    {
    }

    public override async Task<TagModel> AddAsync(TagModel model, CancellationToken cancellationToken = default) =>
        await base.AddAsync(model, cancellationToken);

    public override async Task<TagModel> UpdateAsync(TagModel model, CancellationToken cancellationToken = default) =>
        await base.UpdateAsync(model, cancellationToken);

    public override async Task DeleteAsync(TagId id, CancellationToken cancellationToken = default) =>
        await base.DeleteAsync(id, cancellationToken);
}