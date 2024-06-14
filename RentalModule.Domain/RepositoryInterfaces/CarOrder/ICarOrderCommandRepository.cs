using RentalModule.Domain.Models;
using RentalModule.Domain.Models.Ids;

namespace RentalModule.Domain.RepositoryInterfaces.CarOrder;

public interface ICarOrderCommandRepository
{
    public Task<CarOrderModel> AddAsync(CarOrderModel model, CancellationToken cancellationToken);
    public Task<CarOrderModel> UpdateAsync(CarOrderModel model, CancellationToken cancellationToken);
    public Task DeleteAsync(CarOrderId id, CancellationToken cancellationToken);
}