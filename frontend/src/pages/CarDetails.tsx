import { useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { useQuery, useMutation } from '@tanstack/react-query';
import { carService } from '../services/api';
import { useSelector } from 'react-redux';
import { RootState } from '../store';

const CarDetails = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const { isAuthenticated } = useSelector((state: RootState) => state.auth);
  const [reservation, setReservation] = useState({
    startDate: '',
    endDate: '',
  });

  const { data: car, isLoading, error } = useQuery({
    queryKey: ['carOffer', id],
    queryFn: () => carService.getCarOfferById(id!),
    enabled: !!id,
  });

  const createOrderMutation = useMutation({
    mutationFn: (orderData: { startDate: string; endDate: string }) =>
      carService.createCarOrder({
        carOfferId: id!,
        startDate: orderData.startDate,
        endDate: orderData.endDate,
      }),
    onSuccess: () => {
      navigate('/profile');
    },
  });

  const handleReservation = (e: React.FormEvent) => {
    e.preventDefault();
    if (!isAuthenticated) {
      navigate('/login', { state: { from: `/cars/${id}` } });
      return;
    }
    createOrderMutation.mutate(reservation);
  };

  if (isLoading) {
    return (
      <div className="flex h-96 items-center justify-center">
        <div className="h-32 w-32 animate-spin rounded-full border-b-2 border-t-2 border-primary-500"></div>
      </div>
    );
  }

  if (error || !car) {
    return (
      <div className="flex h-96 items-center justify-center">
        <p className="text-red-500">Error loading car details. Please try again later.</p>
      </div>
    );
  }

  return (
    <div className="space-y-8">
      {/* Car Details */}
      <div className="overflow-hidden bg-white shadow sm:rounded-lg">
        <div className="px-4 py-5 sm:px-6">
          <h1 className="text-3xl font-bold leading-tight text-gray-900">{car.name}</h1>
          <p className="mt-1 max-w-2xl text-sm text-gray-500">{car.description}</p>
        </div>

        {/* Image Gallery */}
        <div className="border-t border-gray-200">
          <div className="aspect-h-3 aspect-w-4 overflow-hidden">
            <img
              src={car.imageUrl}
              alt={car.name}
              className="h-96 w-full object-cover object-center"
            />
          </div>
        </div>

        {/* Car Information */}
        <div className="border-t border-gray-200 px-4 py-5 sm:px-6">
          <dl className="grid grid-cols-1 gap-x-4 gap-y-8 sm:grid-cols-2">
            <div className="sm:col-span-1">
              <dt className="text-sm font-medium text-gray-500">Price per Day</dt>
              <dd className="mt-1 text-lg font-semibold text-primary-600">${car.pricePerDay}</dd>
            </div>
            <div className="sm:col-span-1">
              <dt className="text-sm font-medium text-gray-500">Brand</dt>
              <dd className="mt-1 text-sm text-gray-900">{car.brand}</dd>
            </div>
            <div className="sm:col-span-1">
              <dt className="text-sm font-medium text-gray-500">Model</dt>
              <dd className="mt-1 text-sm text-gray-900">{car.model}</dd>
            </div>
            <div className="sm:col-span-1">
              <dt className="text-sm font-medium text-gray-500">Year</dt>
              <dd className="mt-1 text-sm text-gray-900">{car.year}</dd>
            </div>
          </dl>
        </div>

        {/* Reservation Form */}
        <div className="border-t border-gray-200 px-4 py-5 sm:px-6">
          <h3 className="text-lg font-medium leading-6 text-gray-900">Make a Reservation</h3>
          <form onSubmit={handleReservation} className="mt-6 space-y-6">
            <div className="grid grid-cols-1 gap-6 sm:grid-cols-2">
              <div>
                <label htmlFor="startDate" className="block text-sm font-medium text-gray-700">
                  Start Date
                </label>
                <input
                  type="datetime-local"
                  id="startDate"
                  required
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm"
                  value={reservation.startDate}
                  onChange={(e) => setReservation({ ...reservation, startDate: e.target.value })}
                />
              </div>
              <div>
                <label htmlFor="endDate" className="block text-sm font-medium text-gray-700">
                  End Date
                </label>
                <input
                  type="datetime-local"
                  id="endDate"
                  required
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm"
                  value={reservation.endDate}
                  onChange={(e) => setReservation({ ...reservation, endDate: e.target.value })}
                />
              </div>
            </div>

            <div className="flex justify-end">
              <button
                type="submit"
                disabled={createOrderMutation.isPending}
                className="inline-flex justify-center rounded-md border border-transparent bg-primary-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              >
                {createOrderMutation.isPending ? 'Processing...' : 'Reserve Now'}
              </button>
            </div>

            {createOrderMutation.isError && (
              <div className="mt-2 text-sm text-red-600">
                An error occurred while processing your reservation. Please try again.
              </div>
            )}
          </form>
        </div>
      </div>
    </div>
  );
};

export default CarDetails; 