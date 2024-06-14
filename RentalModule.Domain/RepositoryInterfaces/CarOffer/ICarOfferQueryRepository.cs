using System.Collections.Immutable;
using RentalModule.Domain.Enums;
using RentalModule.Domain.Models;
using RentalModule.Domain.Models.Ids;
using Results.Contract;

namespace RentalModule.Domain.RepositoryInterfaces.CarOffer;

public interface ICarOfferQueryRepository
{
    public Task<CarOfferModel?> GetByIdAsync(CarOfferId id, CancellationToken cancellationToken);
    
    public Task<int> GetTotalCountAsync(CancellationToken cancellationToken);
    public Task<CarOfferModel?> GetByUrlAsync(string urlHandle, CancellationToken cancellationToken);
    //public Task<ImmutableArray<CarOfferModel>> GetByTagAsync(string tagName, CancellationToken cancellationToken);
    
    public Task<ImmutableArray<CarOfferModel>> GetCollectionAsync(
        int? page,
        int? pageSize,
        CarOfferSortColumnEnum? orderBy,
        SortOrderEnum? orderDirection,
        ImmutableArray<DateTime>? possibleDates,
        ImmutableArray<string>? tags,
        CancellationToken cancellationToken);
}