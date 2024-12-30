import { useState, useEffect } from 'react';
import { useParams, useNavigate, useLocation, useSearchParams } from 'react-router-dom';
import { useQuery, useMutation } from '@tanstack/react-query';
import { useSelector } from 'react-redux';
import { carService, fileService } from '../services/api';
import { selectCurrentUser, selectIsAuthenticated } from '../store/slices/authSlice';
import type { CreateCarOrderRequest } from '../types/api';

const CarOrderPage = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const { id } = useParams();
  const user = useSelector(selectCurrentUser);
  const searchParams = useSearchParams()[0];
  const isAuthenticated = useSelector(selectIsAuthenticated);

  // Redirect if not authenticated
  useEffect(() => {
    if (!isAuthenticated) {
      navigate('/login');
    }
  }, [isAuthenticated, navigate]);

  const [formData, setFormData] = useState({
    name: user?.name || '',
    surname: user?.surname || '',
    phoneNumber: user?.phone_number || '',
    emailAddress: user?.email_address || '',
    address: '',
    postcode: '',
    city: '',
    startDate: searchParams.get('startDate') || '',
    endDate: searchParams.get('endDate') || '',
    driversLicenseNumber: '',
    numOfDrivers: 1,
  });

  const { data: carOffer, isLoading } = useQuery({
    queryKey: ['carOffer', id],
    queryFn: () => carService.getCarOfferById(id!),
    enabled: !!id,
  });

  const createOrderMutation = useMutation({
    mutationFn: (orderData: CreateCarOrderRequest) => carService.createCarOrder(orderData),
    onSuccess: () => {
      navigate('/car-offers');
    },
  });

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: value
    }));
  };

  const calculateTotalPrice = () => {
    if (!formData.startDate || !formData.endDate || !carOffer?.car_offer) return 0;
    
    const start = new Date(formData.startDate);
    const end = new Date(formData.endDate);
    const days = Math.ceil((end.getTime() - start.getTime()) / (1000 * 60 * 60 * 24));
    const driverSurcharge = 200;
    
    return (days * carOffer.car_offer.one_normal_day_price) + 
           (driverSurcharge * (formData.numOfDrivers - 1));
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (!id || !carOffer?.car_offer) return;

    const totalPrice = calculateTotalPrice();
    
    createOrderMutation.mutate({
      carOfferId: id,
      startDate: formData.startDate,
      endDate: formData.endDate,
      numOfDrivers: formData.numOfDrivers,
      totalCost: totalPrice,
      deliveryLocation: `${formData.address}, ${formData.postcode} ${formData.city}`,
      returnLocation: `${formData.address}, ${formData.postcode} ${formData.city}`,
    });
  };

  if (isLoading) {
    return (
      <div className="flex h-96 items-center justify-center">
        <div className="h-32 w-32 animate-spin rounded-full border-b-2 border-t-2 border-primary-500"></div>
      </div>
    );
  }

  if (!carOffer?.car_offer) {
    return (
      <div className="flex h-96 items-center justify-center">
        <p className="text-red-500">Car offer not found.</p>
      </div>
    );
  }

  return (
    <form onSubmit={handleSubmit} className="container mx-auto px-4">
      <div className="mt-5 grid grid-cols-1 gap-6 lg:grid-cols-3">
        {/* Personal Information */}
        <div className="rounded-3xl border border-gray-900 bg-white p-6 shadow-lg">
          <h2 className="mb-6 text-2xl font-bold text-gray-900">Your personal information</h2>
          
          <div className="space-y-4">
            <div>
              <label htmlFor="name" className="block text-sm font-medium text-gray-900">name</label>
              <input
                type="text"
                id="name"
                name="name"
                required
                className="mt-1 block w-full rounded-xl border-2 border-gray-200 bg-gray-25 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-900"
                value={formData.name}
                onChange={handleInputChange}
              />
            </div>

            <div>
              <label htmlFor="surname" className="block text-sm font-medium text-gray-900">surname</label>
              <input
                type="text"
                id="surname"
                name="surname"
                required
                className="mt-1 block w-full rounded-xl border-2 border-gray-200 bg-gray-25 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-900"
                value={formData.surname}
                onChange={handleInputChange}
              />
            </div>

            <div>
              <label htmlFor="phoneNumber" className="block text-sm font-medium text-gray-900">phone number</label>
              <input
                type="tel"
                id="phoneNumber"
                name="phoneNumber"
                required
                className="mt-1 block w-full rounded-xl border-2 border-gray-200 bg-gray-25 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-900"
                value={formData.phoneNumber}
                onChange={handleInputChange}
              />
            </div>

            <div>
              <label htmlFor="emailAddress" className="block text-sm font-medium text-gray-900">email address</label>
              <input
                type="email"
                id="emailAddress"
                name="emailAddress"
                required
                className="mt-1 block w-full rounded-xl border-2 border-gray-200 bg-gray-25 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-900"
                value={formData.emailAddress}
                onChange={handleInputChange}
              />
            </div>

            <div>
              <label htmlFor="address" className="block text-sm font-medium text-gray-900">address</label>
              <input
                type="text"
                id="address"
                name="address"
                required
                className="mt-1 block w-full rounded-xl border-2 border-gray-200 bg-gray-25 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-900"
                value={formData.address}
                onChange={handleInputChange}
              />
            </div>

            <div className="grid grid-cols-2 gap-4">
              <div>
                <label htmlFor="postcode" className="block text-sm font-medium text-gray-900">postcode</label>
                <input
                  type="text"
                  id="postcode"
                  name="postcode"
                  required
                  className="mt-1 block w-full rounded-xl border-2 border-gray-200 bg-gray-25 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-900"
                  value={formData.postcode}
                  onChange={handleInputChange}
                />
              </div>
              <div>
                <label htmlFor="city" className="block text-sm font-medium text-gray-900">city</label>
                <input
                  type="text"
                  id="city"
                  name="city"
                  required
                  className="mt-1 block w-full rounded-xl border-2 border-gray-200 bg-gray-25 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-900"
                  value={formData.city}
                  onChange={handleInputChange}
                />
              </div>
            </div>
          </div>
        </div>

        {/* Reservation Details */}
        <div className="rounded-3xl border border-gray-900 bg-white p-6 shadow-lg">
          <h2 className="mb-6 text-2xl font-bold text-gray-900">Your reservation</h2>
          
          <div className="space-y-4">
            <div className="grid grid-cols-2 gap-4">
              <div>
                <label htmlFor="startDate" className="block text-sm font-medium text-gray-900">Start Date</label>
                <input
                  type="date"
                  id="startDate"
                  name="startDate"
                  required
                  className="mt-1 block w-full rounded-xl border-2 border-gray-200 bg-gray-25 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-900"
                  value={formData.startDate}
                  onChange={handleInputChange}
                />
              </div>
              <div>
                <label htmlFor="endDate" className="block text-sm font-medium text-gray-900">End Date</label>
                <input
                  type="date"
                  id="endDate"
                  name="endDate"
                  required
                  className="mt-1 block w-full rounded-xl border-2 border-gray-200 bg-gray-25 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-900"
                  value={formData.endDate}
                  onChange={handleInputChange}
                />
              </div>
            </div>

            <div>
              <label htmlFor="driversLicenseNumber" className="block text-sm font-medium text-gray-900">
                driver licence number
              </label>
              <input
                type="text"
                id="driversLicenseNumber"
                name="driversLicenseNumber"
                required
                className="mt-1 block w-full rounded-xl border-2 border-gray-200 bg-gray-25 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-900"
                value={formData.driversLicenseNumber}
                onChange={handleInputChange}
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-900">Number of drivers</label>
              <div className="mt-1 flex items-center space-x-2">
                <button
                  type="button"
                  onClick={() => setFormData(prev => ({
                    ...prev,
                    numOfDrivers: Math.max(1, prev.numOfDrivers - 1)
                  }))}
                  className="rounded-xl border border-gray-300 p-2 hover:bg-gray-50 text-gray-900"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path fillRule="evenodd" d="M3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z" clipRule="evenodd" />
                  </svg>
                </button>
                <input
                  type="number"
                  id="numOfDrivers"
                  name="numOfDrivers"
                  min="1"
                  className="block w-20 rounded-xl border-2 border-gray-200 bg-gray-25 text-center shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-900"
                  value={formData.numOfDrivers}
                  onChange={handleInputChange}
                />
                <button
                  type="button"
                  onClick={() => setFormData(prev => ({
                    ...prev,
                    numOfDrivers: prev.numOfDrivers + 1
                  }))}
                  className="rounded-xl border border-gray-300 p-2 hover:bg-gray-50 text-gray-900"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path fillRule="evenodd" d="M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z" clipRule="evenodd" />
                  </svg>
                </button>
              </div>
              <p className="mt-1 text-sm text-gray-900">The additional driver surcharge is 200 USD.</p>
            </div>
          </div>
        </div>

        {/* Summary and Payment */}
        <div className="rounded-3xl border border-gray-900 bg-white p-6 shadow-lg">
          <h2 className="mb-6 text-2xl font-bold text-gray-900">Summary and payment</h2>
          
          <div className="space-y-6">
            <div>
              <h3 className="text-lg font-medium text-gray-900">Vehicle</h3>
              <div className="mt-2 flex items-center space-x-4">
                <img
                  src={carOffer.car_offer.featured_image_url ? fileService.getFileUrl(carOffer.car_offer.featured_image_url) : '/placeholder-car.jpg'}
                  alt={carOffer.car_offer.heading}
                  className="h-16 w-24 rounded-lg object-cover"
                />
                <span className="font-medium text-gray-900">{carOffer.car_offer.heading}</span>
              </div>
            </div>

            <div>
              <h3 className="text-lg font-medium text-gray-900">General Terms and Conditions</h3>
              <ul className="mt-2 space-y-2 text-sm text-gray-900">
                <li className="flex items-center">
                  <svg className="mr-2 h-5 w-5 text-green-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 13l4 4L19 7" />
                  </svg>
                  The security deposit - 1500 USD.
                </li>
                <li className="flex items-center">
                  <svg className="mr-2 h-5 w-5 text-green-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 13l4 4L19 7" />
                  </svg>
                  Daily mileage limit: 300 km.
                </li>
                <li className="flex items-center">
                  <svg className="mr-2 h-5 w-5 text-green-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 13l4 4L19 7" />
                  </svg>
                  The fee for exceeding the mileage limit is 2 USD/KM.
                </li>
              </ul>
              <div className="mt-4">
                <label className="flex items-center">
                  <input
                    type="checkbox"
                    required
                    className="rounded border-gray-300 text-primary-600 focus:ring-primary-500"
                  />
                  <span className="ml-2 text-sm text-gray-900">I agree</span>
                </label>
              </div>
            </div>

            <div className="rounded-xl bg-gray-50 p-4">
              <div className="flex items-center justify-between">
                <span className="text-xl font-bold text-gray-900">Rent price</span>
                <div className="text-right">
                  <span className="text-2xl font-bold text-gray-900">{calculateTotalPrice()} USD</span>
                  <div className="text-sm text-gray-900">gross</div>
                </div>
              </div>
            </div>

            <button
              type="submit"
              disabled={createOrderMutation.isPending}
              className="w-full rounded-full bg-primary-600 py-3 text-white hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
            >
              {createOrderMutation.isPending ? 'Processing...' : 'Pay online'}
            </button>
          </div>
        </div>
      </div>
    </form>
  );
};

export default CarOrderPage; 