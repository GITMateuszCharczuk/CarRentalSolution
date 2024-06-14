using System.Collections.Immutable;
using RentalModule.Domain.Enums;
using RentalModule.Domain.Models;
using RentalModule.Domain.Models.Ids;
using Results.Contract;

namespace RentalModule.Domain.RepositoryInterfaces.CarOrder;

public interface ICarOrderQueryRepository
{
    public Task<CarOrderModel?> GetByIdAsync(CarOrderId id, CancellationToken cancellationToken);
    
    public Task<int> GetTotalCountAsync(CancellationToken cancellationToken);
    // public Task<CarOrderModel?> GetByUrlAsync(string urlHandle, CancellationToken cancellationToken);
    //public Task<ImmutableArray<CarOfferModel>> GetByTagAsync(string tagName, CancellationToken cancellationToken);
    
    public Task<ImmutableArray<CarOrderModel>> GetCollectionAsync(
        int? page,
        int? pageSize,
        CarOrderSortColumnEnum? orderBy,
        SortOrderEnum? orderDirection,
        ImmutableArray<DateTime>? dates,
        Guid userId,
        CarOfferId carOfferId,
        CancellationToken cancellationToken);
}    