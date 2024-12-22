import { useState } from 'react';
import { useQuery } from '@tanstack/react-query';
import { useSelector } from 'react-redux';
import { carService } from '../services/api';
import { RootState } from '../store';
import { Link } from 'react-router-dom';

const Profile = () => {
  const { user } = useSelector((state: RootState) => state.auth);
  const [activeTab, setActiveTab] = useState<'orders' | 'profile'>('orders');

  const { data: orders, isLoading } = useQuery({
    queryKey: ['userOrders'],
    queryFn: () =>
      carService.getCarOrders({
        user_id: user?.id,
        page_size: 10,
        current_page: 1,
      }),
  });

  if (!user) {
    return (
      <div className="flex h-96 items-center justify-center">
        <p className="text-gray-500">Please log in to view your profile.</p>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {/* Profile Header */}
      <div className="border-b border-gray-200">
        <div className="pb-5">
          <h1 className="text-3xl font-bold leading-tight text-gray-900">My Profile</h1>
          <p className="mt-2 max-w-4xl text-sm text-gray-500">
            Manage your account and view your rental history.
          </p>
        </div>

        {/* Tabs */}
        <div className="mt-4">
          <nav className="-mb-px flex space-x-8">
            <button
              onClick={() => setActiveTab('orders')}
              className={`${
                activeTab === 'orders'
                  ? 'border-primary-500 text-primary-600'
                  : 'border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700'
              } whitespace-nowrap border-b-2 px-1 pb-4 text-sm font-medium`}
            >
              My Orders
            </button>
            <button
              onClick={() => setActiveTab('profile')}
              className={`${
                activeTab === 'profile'
                  ? 'border-primary-500 text-primary-600'
                  : 'border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700'
              } whitespace-nowrap border-b-2 px-1 pb-4 text-sm font-medium`}
            >
              Profile Information
            </button>
          </nav>
        </div>
      </div>

      {/* Tab Content */}
      <div className="mt-6">
        {activeTab === 'orders' ? (
          <div className="space-y-6">
            <h2 className="text-lg font-medium text-gray-900">Order History</h2>
            {isLoading ? (
              <div className="flex h-48 items-center justify-center">
                <div className="h-8 w-8 animate-spin rounded-full border-b-2 border-t-2 border-primary-500"></div>
              </div>
            ) : orders?.items.length ? (
              <div className="overflow-hidden bg-white shadow sm:rounded-lg">
                <ul className="divide-y divide-gray-200">
                  {orders.items.map((order) => (
                    <li key={order.id} className="p-4 hover:bg-gray-50">
                      <div className="flex items-center justify-between">
                        <div className="flex-1">
                          <h3 className="text-sm font-medium text-gray-900">
                            Order #{order.id.slice(0, 8)}
                          </h3>
                          <p className="mt-1 text-sm text-gray-500">
                            {new Date(order.startDate).toLocaleDateString()} -{' '}
                            {new Date(order.endDate).toLocaleDateString()}
                          </p>
                        </div>
                        <div className="ml-4">
                          <span
                            className={`inline-flex rounded-full px-2 text-xs font-semibold leading-5 ${
                              order.status === 'Completed'
                                ? 'bg-green-100 text-green-800'
                                : order.status === 'Pending'
                                ? 'bg-yellow-100 text-yellow-800'
                                : order.status === 'Cancelled'
                                ? 'bg-red-100 text-red-800'
                                : 'bg-gray-100 text-gray-800'
                            }`}
                          >
                            {order.status}
                          </span>
                        </div>
                      </div>
                      <div className="mt-2">
                        <Link
                          to={`/cars/${order.carOfferId}`}
                          className="text-sm font-medium text-primary-600 hover:text-primary-500"
                        >
                          View Car Details →
                        </Link>
                      </div>
                    </li>
                  ))}
                </ul>
              </div>
            ) : (
              <p className="text-center text-gray-500">No orders found.</p>
            )}
          </div>
        ) : (
          <div className="space-y-6">
            <h2 className="text-lg font-medium text-gray-900">Profile Information</h2>
            <div className="overflow-hidden bg-white shadow sm:rounded-lg">
              <div className="px-4 py-5 sm:p-6">
                <dl className="grid grid-cols-1 gap-x-4 gap-y-8 sm:grid-cols-2">
                  <div className="sm:col-span-1">
                    <dt className="text-sm font-medium text-gray-500">Full name</dt>
                    <dd className="mt-1 text-sm text-gray-900">
                      {user.name} {user.surname}
                    </dd>
                  </div>
                  <div className="sm:col-span-1">
                    <dt className="text-sm font-medium text-gray-500">Email address</dt>
                    <dd className="mt-1 text-sm text-gray-900">{user.email_address}</dd>
                  </div>
                  <div className="sm:col-span-1">
                    <dt className="text-sm font-medium text-gray-500">Phone number</dt>
                    <dd className="mt-1 text-sm text-gray-900">{user.phone_number}</dd>
                  </div>
                  <div className="sm:col-span-1">
                    <dt className="text-sm font-medium text-gray-500">Address</dt>
                    <dd className="mt-1 text-sm text-gray-900">
                      {user.address}, {user.city} {user.postal_code}
                    </dd>
                  </div>
                </dl>
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default Profile; 