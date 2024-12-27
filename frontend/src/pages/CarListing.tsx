import { useState, useEffect, useMemo } from 'react';
import { useQuery } from '@tanstack/react-query';
import { Link } from 'react-router-dom';
import { carService, fileService } from '../services/api';
import type { CarOffer } from '../types/api';
import { formatDateForApi, formatDateForInput } from '../utils/dateUtils';

const CarListing = () => {
  const [currentPage, setCurrentPage] = useState(1);
  const [filters, setFilters] = useState({
    dateTimeFrom: '',
    dateTimeTo: '',
    tags: [] as string[],
  });

  // Calculate minDateTime once and memoize it
  const minDateTime = useMemo(() => {
    const now = new Date();
    return formatDateForInput(now.toISOString());
  }, []);

  // Set initial dates when component mounts
  useEffect(() => {
    if (!filters.dateTimeFrom) {
      const now = new Date();
      now.setDate(now.getDate() + 1);
      setFilters(prev => ({
        ...prev,
        dateTimeFrom: formatDateForInput(now.toISOString())
      }));
    }
  }, []);

  // Validate and update dates when setting filters
  const handleDateChange = (field: 'dateTimeFrom' | 'dateTimeTo', value: string) => {
    const selectedDate = new Date(value);
    const currentDate = new Date();
    

    // If selected date is in the past, use current date/time
    if (selectedDate < currentDate) {
      value = formatDateForInput(currentDate.toISOString());
    }

    // If setting end date and it's before start date, adjust it
    if (field === 'dateTimeTo' && filters.dateTimeFrom && value < filters.dateTimeFrom) {
      value = filters.dateTimeFrom;
    }

    // If setting start date and it's after end date, adjust end date
    if (field === 'dateTimeFrom' && filters.dateTimeTo && value > filters.dateTimeTo) {
      setFilters(prev => ({
        ...prev,
        [field]: value,
        dateTimeTo: value
      }));
      return;
    }

    setFilters(prev => ({
      ...prev,
      [field]: value
    }));
  };

  const { data, isLoading, error } = useQuery({
    queryKey: ['carOffers', currentPage, filters],
    queryFn: () =>
      carService.getCarOffers({
        current_page: currentPage,
        page_size: 9,
        date_time_from: filters.dateTimeFrom ? formatDateForApi(filters.dateTimeFrom) : undefined,
        date_time_to: filters.dateTimeTo ? formatDateForApi(filters.dateTimeTo) : undefined,
        tags: filters.tags,
        visible: undefined,
      }),
  });

  if (isLoading) {
    return (
      <div className="flex h-96 items-center justify-center">
        <div className="h-32 w-32 animate-spin rounded-full border-b-2 border-t-2 border-primary-500"></div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex h-96 items-center justify-center">
        <p className="text-red-500">Error loading car offers. Please try again later.</p>
      </div>
    );
  }

  if (!data?.Items) {
    console.log(data);
    return (
      <div className="flex h-96 items-center justify-center">
        <p className="text-gray-500">No car offers available.</p>
      </div>
    );
  }

  return (
    <div className="space-y-8">
      {/* Header */}
      <div className="border-b border-gray-200 pb-5">
        <h1 className="text-3xl font-bold leading-tight text-gray-900">Available Cars</h1>
        <p className="mt-2 max-w-4xl text-sm text-gray-500">
          Browse our selection of premium vehicles available for rent.
        </p>
      </div>

      {/* Filters */}
      <div className="bg-white p-4 shadow sm:rounded-lg">
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
          <div>
            <label htmlFor="dateFrom" className="block text-sm font-medium text-gray-700">
              From Date
            </label>
            <input
              type="datetime-local"
              id="dateFrom"
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm"
              value={filters.dateTimeFrom}
              onChange={(e) => handleDateChange('dateTimeFrom', e.target.value)}
              min={minDateTime}
            />
          </div>
          <div>
            <label htmlFor="dateTo" className="block text-sm font-medium text-gray-700">
              To Date
            </label>
            <input
              type="datetime-local"
              id="dateTo"
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm"
              value={filters.dateTimeTo}
              onChange={(e) => handleDateChange('dateTimeTo', e.target.value)}
              min={minDateTime}
            />
          </div>
        </div>
      </div>

      {/* Car Grid */}
      <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
        {data.Items.map((car: CarOffer) => (
          <Link key={car.id} to={`/cars/${car.id}`} className="group">
            <div className="aspect-h-2 aspect-w-3 overflow-hidden rounded-lg">
              <img
                src={car.featured_image_url ? fileService.getFileUrl(car.featured_image_url) : '/placeholder-car.jpg'}
                alt={car.heading}
                className="h-48 w-full object-cover object-center"
              />
            </div>
            <div className="p-4">
              <h3 className="text-lg font-semibold text-gray-900">{car.heading}</h3>
              <p className="mt-1 text-sm text-gray-500">
                {car.short_description.length > 135 ? `${car.short_description.substring(0, 135)}...` : car.short_description}
              </p>
              <div className="mt-2 flex items-center justify-between">
                <p className="text-lg font-bold text-primary-600">${car.one_normal_day_price}/day</p>
                <Link
                  to={`/cars/${car.id}`}
                  className="rounded-md bg-primary-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-primary-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-600"
                >
                  View Details
                </Link>
              </div>
            </div>
          </Link>
        ))}
      </div>

      {/* Pagination */}
      {data && data.TotalPages > 1 && (
        <div className="flex items-center justify-center space-x-2">
          <button
            onClick={() => setCurrentPage((prev) => Math.max(prev - 1, 1))}
            disabled={currentPage === 1}
            className="rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 disabled:cursor-not-allowed disabled:opacity-50"
          >
            Previous
          </button>
          <span className="text-sm text-gray-700">
            Page {currentPage} of {data.TotalPages}
          </span>
          <button
            onClick={() => setCurrentPage((prev) => Math.min(prev + 1, data.TotalPages))}
            disabled={currentPage === data.TotalPages}
            className="rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 disabled:cursor-not-allowed disabled:opacity-50"
          >
            Next
          </button>
        </div>
      )}
    </div>
  );
};

export default CarListing; 