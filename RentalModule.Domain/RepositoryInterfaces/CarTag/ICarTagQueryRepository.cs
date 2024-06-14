using System.Collections.Immutable;
using BlogModule.Domain.Enums;
using RentalModule.Domain.Models;
using RentalModule.Domain.Models.Ids;
using Results.Contract;

namespace RentalModule.Domain.RepositoryInterfaces.CarTag;

public interface ICarTagQueryRepository
{ 
        public Task<ImmutableArray<CarTagModel>> GetCollectionAsync(CarTagSortColumnEnum? orderBy,
        SortOrderEnum? orderDirection, CarOfferId? carOfferId, CancellationToken cancellationToken);
}