using System.Collections.Immutable;
using Microsoft.EntityFrameworkCore;
using RentalModule.Domain.Enums;
using RentalModule.Domain.Models;
using RentalModule.Domain.Models.Ids;
using RentalModule.Domain.RepositoryInterfaces.CarOrder;
using RentalModule.Infrastructure.DataBase.Context;
using RentalModule.Infrastructure.DataBase.Entities;
using Results.Contract;
using Shared.CQRS.Repository;
using Shared.Utilities;

namespace RentalModule.Infrastructure.DataBase.Repository.CarOrder
{
    public class CarOrderQueryRepository : QueryRepository<CarOrderEntity, CarOrderId, CarOrderModel, RentalDbContext>, ICarOrderQueryRepository
    {
        public CarOrderQueryRepository(RentalDbContext dbContext, IPersistenceMapper<CarOrderEntity, CarOrderModel> mapper) 
            : base(dbContext, mapper)
        {
        }

        public override async Task<CarOrderModel?> GetByIdAsync(CarOrderId id, CancellationToken cancellationToken = default) =>
            await base.GetByIdAsync(id, cancellationToken);

        public override async Task<int> GetTotalCountAsync(CancellationToken cancellationToken = default) =>
            await base.GetTotalCountAsync(cancellationToken);

        public async Task<ImmutableArray<CarOrderModel>> GetCollectionAsync(
            int? page,
            int? pageSize,
            CarOrderSortColumnEnum? orderBy,
            SortOrderEnum? orderDirection,
            ImmutableArray<DateTime>? dates,
            Guid userId,
            CarOfferId carOfferId,
            CancellationToken cancellationToken)
        {
            var queryableOrders = DbContext.CarOrders
                .AsNoTracking()
                .Where(co => co.UserId == userId && co.CarOfferId == carOfferId)
                .AsQueryable();

            if (dates is not null && dates.Value.Any())
            {
                var dateRange = dates.Value.ToArray();
                queryableOrders = queryableOrders.Where(co => dateRange.Contains(co.StartDate) || dateRange.Contains(co.EndDate));
            }

            if (orderBy is not null && orderDirection is not null)
            {
                var isOrderDirectionAscending = orderDirection == SortOrderEnum.Ascending;
                queryableOrders = orderBy switch
                {
                    CarOrderSortColumnEnum.StartDate => isOrderDirectionAscending
                        ? queryableOrders.OrderBy(co => co.StartDate)
                        : queryableOrders.OrderByDescending(co => co.StartDate),
                    CarOrderSortColumnEnum.EndDate => isOrderDirectionAscending
                        ? queryableOrders.OrderBy(co => co.EndDate)
                        : queryableOrders.OrderByDescending(co => co.EndDate),
                    CarOrderSortColumnEnum.TotalCost => isOrderDirectionAscending
                        ? queryableOrders.OrderBy(co => co.TotalCost)
                        : queryableOrders.OrderByDescending(co => co.TotalCost),
                    _ => queryableOrders
                };
            }

            if (page is not null && pageSize is not null)
            {
                queryableOrders = queryableOrders
                    .Skip((page.Value - 1) * pageSize.Value)
                    .Take(pageSize.Value);
            }

            return await queryableOrders
                .Select(co => Mapper.MapToModel(co))
                .ToImmutableArrayAsync(cancellationToken);
        }
    }
}
