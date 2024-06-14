using System.Collections.Immutable;
using Microsoft.EntityFrameworkCore;
using RentalModule.Domain.Enums;
using RentalModule.Domain.Models;
using RentalModule.Domain.Models.Ids;
using RentalModule.Domain.RepositoryInterfaces.CarOffer;
using RentalModule.Infrastructure.DataBase.Context;
using RentalModule.Infrastructure.DataBase.Entities;
using Results.Contract;
using Shared.CQRS.Repository;
using Shared.Utilities;

namespace RentalModule.Infrastructure.DataBase.Repository.CarOffer
{
    public class CarOfferQueryRepository : QueryRepository<CarOfferEntity, CarOfferId, CarOfferModel, RentalDbContext>, ICarOfferQueryRepository
    {
        public CarOfferQueryRepository(RentalDbContext dbContext, IPersistenceMapper<CarOfferEntity, CarOfferModel> mapper) 
            : base(dbContext, mapper)
        {
        }

        public override async Task<CarOfferModel?> GetByIdAsync(CarOfferId id, CancellationToken cancellationToken = default) =>
            await base.GetByIdAsync(id, cancellationToken);

        public override async Task<int> GetTotalCountAsync(CancellationToken cancellationToken = default) =>
            await base.GetTotalCountAsync(cancellationToken);

        public async Task<CarOfferModel?> GetByUrlAsync(string urlHandle, CancellationToken cancellationToken) =>
            await DbContext.CarOffers.AsNoTracking()
                .FirstOrDefaultAsync(x => x.UrlHandle == urlHandle, cancellationToken)
                .ContinueWith(x => x.Result is null ? null : Mapper.MapToModel(x.Result), cancellationToken);

        public async Task<ImmutableArray<CarOfferModel>> GetCollectionAsync(
            int? page,
            int? pageSize,
            CarOfferSortColumnEnum? orderBy,
            SortOrderEnum? orderDirection,
            ImmutableArray<DateTime>? possibleDates,
            ImmutableArray<string>? tags,
            CancellationToken cancellationToken)
        {
            var queryableOffers = DbContext.CarOffers
                .Include(co => co.Tags)
                .Include(co => co.ImageUrls)
                .Include(co => co.UnavailableDates)
                .AsNoTracking()
                .AsQueryable();

            if (tags is not null && tags.Value.Any())
            {
                queryableOffers = queryableOffers.Where(co => co.Tags.Any(tag => tags.Contains(tag.Name)));
            }

            if (possibleDates is not null && possibleDates.Value.Any())
            {
                var dates = possibleDates.Value.Select(dt => dt.Date).ToHashSet();
                queryableOffers = queryableOffers.Where(co => co.UnavailableDates.Any(date => dates.Contains(date.StartDate.Date) || dates.Contains(date.EndDate.Date)));
            }

            if (orderBy is not null && orderDirection is not null)
            {
                var isOrderDirectionAscending = orderDirection == SortOrderEnum.Ascending;
                queryableOffers = orderBy switch
                {
                    CarOfferSortColumnEnum.PublishedDate => isOrderDirectionAscending
                        ? queryableOffers.OrderBy(co => co.PublishedDate)
                        : queryableOffers.OrderByDescending(co => co.PublishedDate),
                    CarOfferSortColumnEnum.Heading => isOrderDirectionAscending
                        ? queryableOffers.OrderBy(co => co.Heading)
                        : queryableOffers.OrderByDescending(co => co.Heading),
                    _ => queryableOffers
                };
            }

            if (page is not null && pageSize is not null)
            {
                queryableOffers = queryableOffers
                    .Skip((page.Value - 1) * pageSize.Value)
                    .Take(pageSize.Value);
            }

            return await queryableOffers
                .Select(co => Mapper.MapToModel(co))
                .ToImmutableArrayAsync(cancellationToken);
        }
    }
}
