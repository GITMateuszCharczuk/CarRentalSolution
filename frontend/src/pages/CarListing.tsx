import { useState, useEffect, useMemo } from 'react';
import { Link } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { carService, fileService } from '../services/api';
import { Pagination } from '../components/Pagination';
import { SortSelect, type SortField } from '../components/SortSelect';
import { CarOffer, CarOfferTag } from '../types/api';
import { formatDateForApi, formatDateForInput } from '../utils/dateUtils';

const SORT_FIELDS: SortField[] = [
  { field: 'heading', label: 'Name' },
  { field: 'engine_details', label: 'Engine' },
  { field: 'horsepower', label: 'Horsepower' },
  { field: 'year_of_production', label: 'Year' },
  { field: 'one_normal_day_price', label: 'Price' },
  { field: 'published_date', label: 'Published Date' },
];

const CarListing = () => {
  const [currentPage, setCurrentPage] = useState(1);
  const [selectedTags, setSelectedTags] = useState<string[]>([]);
  const [sortFields, setSortFields] = useState<string[]>([]);
  const [filters, setFilters] = useState({
    dateTimeFrom: '',
    dateTimeTo: '',
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

  const { data: carOffers, isLoading } = useQuery({
    queryKey: ['carOffers', currentPage, selectedTags, sortFields, filters],
    queryFn: () => carService.getCarOffers({
      page_size: 9,
      current_page: currentPage,
      tags: selectedTags.length > 0 ? selectedTags : undefined,
      date_time_from: filters.dateTimeFrom ? formatDateForApi(filters.dateTimeFrom) : undefined,
      date_time_to: filters.dateTimeTo ? formatDateForApi(filters.dateTimeTo) : undefined,
      sort_fields: sortFields.length > 0 ? sortFields : undefined,
    }),
  });

  const { data: tags } = useQuery({
    queryKey: ['carTags'],
    queryFn: () => carService.getCarOfferTags(undefined),
  });

  const handleTagClick = (tagName: string) => {
    setSelectedTags(prev => 
      prev.includes(tagName) 
        ? prev.filter(t => t !== tagName)
        : [...prev, tagName]
    );
    setCurrentPage(1);
  };

  if (isLoading) {
    return (
      <div className="flex h-96 items-center justify-center">
        <div className="h-32 w-32 animate-spin rounded-full border-b-2 border-t-2 border-primary-500"></div>
      </div>
    );
  }

  return (
    <div className="container mx-auto px-4">
      {/* Header Section */}
      <div className="mb-8">
        <h1 className="text-4xl font-bold text-gray-900 mb-2">Available Cars</h1>
        <p className="text-lg text-gray-600">Find your perfect ride for any occasion</p>
      </div>

      {/* Filters Section */}
      <div className="bg-white rounded-3xl shadow-lg p-6 mb-8">
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-xl font-semibold text-gray-900">Search Availability</h2>
          <SortSelect
            availableFields={SORT_FIELDS}
            onChange={setSortFields}
            className="min-w-[200px] text-gray-900"
          />
        </div>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label htmlFor="dateFrom" className="block text-sm font-medium text-gray-700 mb-2">
              From Date
            </label>
            <input
              type="datetime-local"
              id="dateFrom"
              className="w-full rounded-xl border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-900"
              value={filters.dateTimeFrom}
              onChange={(e) => handleDateChange('dateTimeFrom', e.target.value)}
              min={minDateTime}
            />
          </div>
          <div>
            <label htmlFor="dateTo" className="block text-sm font-medium text-gray-700 mb-2">
              To Date
            </label>
            <input
              type="datetime-local"
              id="dateTo"
              className="w-full rounded-xl border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-900"
              value={filters.dateTimeTo}
              onChange={(e) => handleDateChange('dateTimeTo', e.target.value)}
              min={minDateTime}
            />
          </div>
        </div>
      </div>

      {/* Tags Section */}
      <div className="mb-8">
        <h2 className="text-xl font-semibold text-gray-900 mb-4">Categories</h2>
        <div className="flex flex-wrap gap-2">
          {tags?.Items && tags.Items.map((tag: CarOfferTag) => (
            <button
              key={tag.id}
              onClick={() => handleTagClick(tag.name)}
              className={`px-4 py-2 rounded-xl text-sm font-medium transition-all transform hover:scale-105
                ${selectedTags.includes(tag.name)
                  ? 'bg-gray-900 text-white shadow-lg'
                  : 'bg-white text-gray-900 shadow hover:shadow-md border border-gray-200'
                }`}
            >
              {tag.name}
            </button>
          ))}
        </div>
      </div>

      {/* Car Listings */}
      {carOffers?.Items && carOffers.Items.length > 0 ? (
        <div className="grid grid-cols-1 lg:grid-cols-3 md:grid-cols-2 gap-8">
          {carOffers.Items.map((car: CarOffer) => (
            <div key={car.id} className="bg-white rounded-3xl shadow-lg overflow-hidden transform transition-all duration-300 hover:scale-[1.02] hover:shadow-xl">
              <div className="relative">
                <img
                  src={fileService.getFileUrl(car.featured_image_url)}
                  alt={car.heading}
                  className="w-full h-56 object-cover"
                />
                <div className="absolute inset-0 bg-gradient-to-t from-black/50 to-transparent"></div>
              </div>
              <div className="p-6">
                <h3 className="text-2xl font-bold mb-4 text-gray-900 truncate" title={car.heading}>{car.heading}</h3>
                
                <div className="grid grid-cols-2 gap-4 mb-6">
                  <div className="bg-gray-50 p-3 rounded-xl">
                    <div className="flex items-center text-gray-800">
                      <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 min-w-[1.25rem] mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
                      </svg>
                      <span className="font-medium truncate" title={car.engine_details}>{car.engine_details}</span>
                    </div>
                  </div>
                  <div className="bg-gray-50 p-3 rounded-xl">
                    <div className="flex items-center text-gray-800">
                      <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 min-w-[1.25rem] mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
                      </svg>
                      <span className="font-medium truncate" title={`${car.horsepower} hp`}>{car.horsepower} hp</span>
                    </div>
                  </div>
                  <div className="bg-gray-50 p-3 rounded-xl">
                    <div className="flex items-center text-gray-800">
                      <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 min-w-[1.25rem] mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                      </svg>
                      <span className="font-medium truncate" title={car.gearbox_details}>{car.gearbox_details}</span>
                    </div>
                  </div>
                  <div className="bg-gray-50 p-3 rounded-xl">
                    <div className="flex items-center text-gray-800">
                      <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 min-w-[1.25rem] mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 12a4 4 0 100-8 4 4 0 000 8z" />
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 14c-6.627 0-12 1.343-12 3v2h24v-2c0-1.657-5.373-3-12-3z" />
                      </svg>
                      <span className="font-medium truncate" title={car.drive_details}>{car.drive_details}</span>
                    </div>
                  </div>
                </div>

                <div className="flex justify-between items-center mb-6">
                  <span className="text-gray-600">Daily rate</span>
                  <div className="text-right">
                    <span className="text-3xl font-bold text-gray-900">${car.one_normal_day_price}</span>
                    <span className="text-gray-600 ml-1">USD</span>
                    <div className="text-sm text-gray-500">gross</div>
                  </div>
                </div>

                <div className="grid grid-cols-2 gap-4">
                  <Link
                    to={`/cars/${car.id}`}
                    className="bg-gray-900 text-white py-3 px-6 rounded-xl text-center font-medium hover:bg-gray-800 transition-colors"
                  >
                    Details
                  </Link>
                  <Link
                    to={`/cars/${car.id}/book`}
                    className="bg-primary-600 text-white py-3 px-6 rounded-xl text-center font-medium hover:bg-primary-700 transition-colors"
                  >
                    Book Now
                  </Link>
                </div>
              </div>
            </div>
          ))}
        </div>
      ) : (
        <div className="text-center py-12">
          <h3 className="text-xl font-medium text-gray-900 mb-2">No cars available</h3>
          <p className="text-gray-600">Try adjusting your search criteria</p>
        </div>
      )}

      {/* Pagination */}
      {carOffers && carOffers.TotalPages > 1 && (
        <div className="mt-12">
          <Pagination
            currentPage={currentPage}
            totalPages={carOffers.TotalPages}
            onPageChange={setCurrentPage}
          />
        </div>
      )}
    </div>
  );
};

export default CarListing; 