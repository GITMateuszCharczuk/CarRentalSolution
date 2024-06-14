using System.Collections.Immutable;
using BlogModule.Domain.Enums;
using Microsoft.EntityFrameworkCore;
using RentalModule.Domain.Models;
using RentalModule.Domain.Models.Ids;
using RentalModule.Domain.RepositoryInterfaces.CarTag;
using RentalModule.Infrastructure.DataBase.Context;
using RentalModule.Infrastructure.DataBase.Entities;
using Results.Contract;
using Shared.CQRS.Repository;
using Shared.Utilities;

namespace RentalModule.Infrastructure.DataBase.Repository.CarTag;

public class CarTagQueryRepository : QueryRepository<CarTagEntity, CarTagId, CarTagModel, RentalDbContext>, ICarTagQueryRepository
{
    public CarTagQueryRepository(RentalDbContext dbContext, IPersistenceMapper<CarTagEntity, CarTagModel> mapper)
        : base(dbContext, mapper)
    {
    }

    public async Task<ImmutableArray<CarTagModel>> GetCollectionAsync(
        CarTagSortColumnEnum? orderBy,
        SortOrderEnum? orderDirection,
        CarOfferId? carOfferId,
        CancellationToken cancellationToken)
    {
        var queryableTags = DbContext.CarTags
            .AsNoTracking()
            .AsQueryable();

        if (carOfferId is not null)
        {
            queryableTags = queryableTags.Where(ct => ct.CarOfferId == carOfferId);
        }

        if (orderBy is not null && orderDirection is not null)
        {
            var isOrderDirectionAscending = orderDirection == SortOrderEnum.Ascending;
            queryableTags = orderBy switch
            {
                CarTagSortColumnEnum.Name => isOrderDirectionAscending
                    ? queryableTags.OrderBy(ct => ct.Name)
                    : queryableTags.OrderByDescending(ct => ct.Name),
                _ => queryableTags
            };
        }

        return await queryableTags
            .Select(ct => Mapper.MapToModel(ct))
            .ToImmutableArrayAsync(cancellationToken);
    }
}