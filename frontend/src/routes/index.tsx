import { createBrowserRouter } from 'react-router-dom';
import MainLayout from '../layouts/MainLayout';
import AdminLayout from '../layouts/AdminLayout';
import Home from '../pages/Home';
import CarListing from '../pages/CarListing';
import CarDetails from '../pages/CarDetails';
import CarOrder from '../pages/CarOrder';
import Login from '../pages/Login';
import Register from '../pages/Register';
import Profile from '../pages/Profile';
import BlogListing from '../pages/BlogListing';
import BlogPost from '../pages/BlogPost';
import BlogManagement from '../pages/admin/BlogManagement';
import UserManagement from '../pages/admin/UserManagement';
import CarOfferManagement from '../pages/admin/CarOfferManagement';
import CarOrderManagement from '../pages/admin/CarOrderManagement';

const router = createBrowserRouter([
  {
    path: '/',
    element: <MainLayout />,
    children: [
      { index: true, element: <Home /> },
      { path: 'cars', element: <CarListing /> },
      { path: 'cars/:id', element: <CarDetails /> },
      { path: 'car-order/:id', element: <CarOrder /> },
      { path: 'login', element: <Login /> },
      { path: 'register', element: <Register /> },
      { path: 'profile', element: <Profile /> },
      { path: 'blog', element: <BlogListing /> },
      { path: 'blog/:id', element: <BlogPost /> },
    ],
  },
  {
    path: '/admin',
    element: <AdminLayout />,
    children: [
      { path: 'blog', element: <BlogManagement /> },
      { path: 'users', element: <UserManagement /> },
      { path: 'car-offers', element: <CarOfferManagement /> },
      { path: 'car-orders', element: <CarOrderManagement /> },
    ],
  },
]);

export default router; 