import { Link, Outlet, useLocation, Navigate } from 'react-router-dom';
import { useSelector } from 'react-redux';
import { RootState } from '../store';

const AdminLayout = () => {
  const location = useLocation();
  const { user } = useSelector((state: RootState) => state.auth);

  // Redirect non-admin users
  if (!user?.is_admin) {
    return <Navigate to="/login" state={{ from: location }} replace />;
  }

  const navigation = [
    { name: 'Blog Posts', href: '/admin/blog' },
    { name: 'Users', href: '/admin/users' },
    { name: 'Car Offers', href: '/admin/car-offers' },
    { name: 'Car Orders', href: '/admin/car-orders' },
  ];

  return (
    <div className="min-h-screen">
      <div className="flex">
        {/* Sidebar */}
        <div className="w-64 min-h-screen bg-gray-800">
          <div className="flex flex-col h-full">
            <div className="flex items-center h-16 px-4 bg-gray-900">
              <span className="text-lg font-medium text-white">Admin Dashboard</span>
            </div>
            <nav className="flex-1 px-2 py-4 space-y-1">
              {navigation.map((item) => {
                const isActive = location.pathname === item.href;
                return (
                  <Link
                    key={item.name}
                    to={item.href}
                    className={`${
                      isActive
                        ? 'bg-gray-900 text-white'
                        : 'text-gray-300 hover:bg-gray-700 hover:text-white'
                    } group flex items-center px-2 py-2 text-sm font-medium rounded-md`}
                  >
                    {item.name}
                  </Link>
                );
              })}
            </nav>
          </div>
        </div>

        {/* Main content */}
        <div className="flex-1 min-h-screen bg-gray-100">
          <main className="py-6 px-4 sm:px-6 lg:px-8">
            <Outlet />
          </main>
        </div>
      </div>
    </div>
  );
};

export default AdminLayout; 