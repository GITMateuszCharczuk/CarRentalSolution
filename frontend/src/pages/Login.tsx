import { useState } from 'react';
import { Link, useNavigate, useLocation } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { useMutation } from '@tanstack/react-query';
import { authService, userService } from '../services/api';
import { setCredentials } from '../store/slices/authSlice';

const Login = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const dispatch = useDispatch();
  const [formData, setFormData] = useState({
    email: '',
    password: '',
  });

  const from = location.state?.from || '/';

  const loginMutation = useMutation({
    mutationFn: async (credentials: { email: string; password: string }) => {
      const loginResponse = await authService.login(credentials);
      const userInfoResponse = await userService.getUserInfo(undefined, loginResponse.Token);
      return {
        ...loginResponse,
        user: userInfoResponse.user_info
      };
    },
    onSuccess: (response) => {
      dispatch(setCredentials({
        token: response.Token,
        refresh_token: response.RefreshToken,
        roles: response.Roles,
        user: response.user
      }));
      navigate(from, { replace: true });
    },
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    loginMutation.mutate(formData);
  };

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="max-w-2xl mx-auto">
        <div className="bg-white rounded-3xl shadow-lg p-8">
          <div className="text-center mb-8">
            <h1 className="text-4xl font-bold text-gray-900 mb-4">Sign in to your account</h1>
            <p className="text-lg text-gray-600">
              Don't have an account?{' '}
              <Link to="/register" className="text-primary-600 hover:text-primary-500 font-medium">
                Create one now
              </Link>
            </p>
          </div>

          <form onSubmit={handleSubmit} className="space-y-6">
            <div>
              <label htmlFor="email" className="block text-sm font-medium text-gray-700 mb-2">
                Email address
              </label>
              <input
                id="email"
                name="email"
                type="email"
                autoComplete="email"
                required
                className="w-full rounded-xl border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-600"
                value={formData.email}
                onChange={(e) => setFormData({ ...formData, email: e.target.value })}
              />
            </div>

            <div>
              <label htmlFor="password" className="block text-sm font-medium text-gray-700 mb-2">
                Password
              </label>
              <input
                id="password"
                name="password"
                type="password"
                autoComplete="current-password"
                required
                className="w-full rounded-xl border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-600"
                value={formData.password}
                onChange={(e) => setFormData({ ...formData, password: e.target.value })}
              />
            </div>

            {loginMutation.isError && (
              <div className="rounded-xl bg-red-50 p-4 text-sm text-red-600">
                Invalid email or password. Please try again.
              </div>
            )}

            <button
              type="submit"
              disabled={loginMutation.isPending}
              className="w-full rounded-xl bg-primary-600 py-3 px-4 text-white font-medium hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
            >
              {loginMutation.isPending ? (
                <div className="flex items-center justify-center">
                  <div className="h-5 w-5 border-2 border-white border-t-transparent rounded-full animate-spin mr-2"></div>
                  Signing in...
                </div>
              ) : (
                'Sign in'
              )}
            </button>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Login; 