import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { carService } from '../../services/api';
import type { CarOffer } from '../../types/api';

const CarOfferManagement = () => {
  const queryClient = useQueryClient();
  const [currentPage, setCurrentPage] = useState(1);
  const [pageSize, setPageSize] = useState(10);
  const [sortField, setSortField] = useState<keyof CarOffer>('name');
  const [sortOrder, setSortOrder] = useState<'asc' | 'desc'>('asc');
  const [searchTerm, setSearchTerm] = useState('');
  const [priceRange, setPriceRange] = useState({ min: '', max: '' });
  const [selectedTags, setSelectedTags] = useState<string[]>([]);

  const { data, isLoading } = useQuery({
    queryKey: ['adminCarOffers', currentPage, pageSize, sortField, sortOrder, searchTerm, priceRange, selectedTags],
    queryFn: () =>
      carService.getCarOffers({
        current_page: currentPage,
        page_size: pageSize,
        sort_by: sortField,
        sort_order: sortOrder,
        search: searchTerm,
        min_price: priceRange.min ? Number(priceRange.min) : undefined,
        max_price: priceRange.max ? Number(priceRange.max) : undefined,
        tags: selectedTags,
      }),
  });

  const deleteMutation = useMutation({
    mutationFn: (offerId: string) => carService.deleteCarOffer(offerId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['adminCarOffers'] });
    },
  });

  const toggleVisibilityMutation = useMutation({
    mutationFn: (offerId: string) => carService.toggleCarOfferVisibility(offerId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['adminCarOffers'] });
    },
  });

  const handleSort = (field: keyof CarOffer) => {
    if (field === sortField) {
      setSortOrder(sortOrder === 'asc' ? 'desc' : 'asc');
    } else {
      setSortField(field);
      setSortOrder('asc');
    }
  };

  const handleDelete = async (offerId: string) => {
    if (window.confirm('Are you sure you want to delete this car offer?')) {
      await deleteMutation.mutateAsync(offerId);
    }
  };

  return (
    <div className="space-y-6">
      <div className="sm:flex sm:items-center sm:justify-between">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Car Offers Management</h1>
          <p className="mt-2 text-sm text-gray-700">
            Manage car offers, including creation, editing, and deletion.
          </p>
        </div>
        <button
          type="button"
          onClick={() => {/* TODO: Implement create/edit modal */}}
          className="inline-flex items-center rounded-md bg-primary-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-primary-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-600"
        >
          Add New Car Offer
        </button>
      </div>

      {/* Filters */}
      <div className="bg-white p-4 shadow sm:rounded-lg">
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
          <div>
            <input
              type="text"
              placeholder="Search cars..."
              className="block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm"
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
            />
          </div>
          <div>
            <input
              type="number"
              placeholder="Min price"
              className="block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm"
              value={priceRange.min}
              onChange={(e) => setPriceRange({ ...priceRange, min: e.target.value })}
            />
          </div>
          <div>
            <input
              type="number"
              placeholder="Max price"
              className="block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm"
              value={priceRange.max}
              onChange={(e) => setPriceRange({ ...priceRange, max: e.target.value })}
            />
          </div>
          {/* Add tag selection if needed */}
        </div>
      </div>

      {/* Table */}
      <div className="overflow-hidden bg-white shadow sm:rounded-lg">
        <div className="overflow-x-auto">
          <table className="min-w-full divide-y divide-gray-300">
            <thead className="bg-gray-50">
              <tr>
                <th
                  scope="col"
                  className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900 cursor-pointer"
                  onClick={() => handleSort('name')}
                >
                  Name
                </th>
                <th
                  scope="col"
                  className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900 cursor-pointer"
                  onClick={() => handleSort('brand')}
                >
                  Brand
                </th>
                <th
                  scope="col"
                  className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900 cursor-pointer"
                  onClick={() => handleSort('pricePerDay')}
                >
                  Price/Day
                </th>
                <th
                  scope="col"
                  className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900"
                >
                  Status
                </th>
                <th scope="col" className="relative py-3.5 pl-3 pr-4 sm:pr-6">
                  <span className="sr-only">Actions</span>
                </th>
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-200 bg-white">
              {data?.items.map((offer) => (
                <tr key={offer.id}>
                  <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-900">
                    {offer.name}
                  </td>
                  <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                    {offer.brand} {offer.model}
                  </td>
                  <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                    ${offer.pricePerDay}
                  </td>
                  <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                    <span
                      className={`inline-flex rounded-full px-2 text-xs font-semibold leading-5 ${
                        offer.visible
                          ? 'bg-green-100 text-green-800'
                          : 'bg-yellow-100 text-yellow-800'
                      }`}
                    >
                      {offer.visible ? 'Active' : 'Hidden'}
                    </span>
                  </td>
                  <td className="relative whitespace-nowrap py-4 pl-3 pr-4 text-right text-sm font-medium sm:pr-6">
                    <button
                      onClick={() => {/* TODO: Implement edit */}}
                      className="text-primary-600 hover:text-primary-900 mr-4"
                    >
                      Edit
                    </button>
                    <button
                      onClick={() => toggleVisibilityMutation.mutate(offer.id)}
                      className="text-yellow-600 hover:text-yellow-900 mr-4"
                    >
                      {offer.visible ? 'Hide' : 'Show'}
                    </button>
                    <button
                      onClick={() => handleDelete(offer.id)}
                      className="text-red-600 hover:text-red-900"
                    >
                      Delete
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>

      {/* Pagination */}
      {data && (
        <div className="flex items-center justify-between bg-white px-4 py-3 sm:px-6">
          <div className="flex flex-1 justify-between sm:hidden">
            <button
              onClick={() => setCurrentPage((prev) => Math.max(prev - 1, 1))}
              disabled={currentPage === 1}
              className="relative inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50"
            >
              Previous
            </button>
            <button
              onClick={() => setCurrentPage((prev) => Math.min(prev + 1, data.totalPages))}
              disabled={currentPage === data.totalPages}
              className="relative ml-3 inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50"
            >
              Next
            </button>
          </div>
          <div className="hidden sm:flex sm:flex-1 sm:items-center sm:justify-between">
            <div>
              <p className="text-sm text-gray-700">
                Showing <span className="font-medium">{(currentPage - 1) * pageSize + 1}</span> to{' '}
                <span className="font-medium">
                  {Math.min(currentPage * pageSize, data.totalItems)}
                </span>{' '}
                of <span className="font-medium">{data.totalItems}</span> results
              </p>
            </div>
            <div>
              <select
                value={pageSize}
                onChange={(e) => setPageSize(Number(e.target.value))}
                className="rounded-md border-gray-300 py-2 pl-3 pr-10 text-base focus:border-primary-500 focus:outline-none focus:ring-primary-500 sm:text-sm"
              >
                <option value="5">5 per page</option>
                <option value="10">10 per page</option>
                <option value="20">20 per page</option>
                <option value="50">50 per page</option>
              </select>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default CarOfferManagement; 