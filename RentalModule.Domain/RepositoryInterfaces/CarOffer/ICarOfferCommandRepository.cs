using RentalModule.Domain.Models;
using RentalModule.Domain.Models.Ids;

namespace RentalModule.Domain.RepositoryInterfaces.CarOffer;

public interface ICarOfferCommandRepository
{
    public Task<CarOfferModel> AddAsync(CarOfferModel model, CancellationToken cancellationToken);
    public Task<CarOfferModel> UpdateAsync(CarOfferModel model, CancellationToken cancellationToken);
    public Task DeleteAsync(CarOfferId id, CancellationToken cancellationToken);
}