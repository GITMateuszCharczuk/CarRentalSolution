import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { useMutation } from '@tanstack/react-query';
import { authService } from '../services/api';
import { setCredentials } from '../store/slices/authSlice';
import type { RegisterRequest, AuthResponse } from '../types/api';

const Register = () => {
  const navigate = useNavigate();
  const dispatch = useDispatch();
  const [formData, setFormData] = useState({
    email_address: '',
    password: '',
    confirmPassword: '',
    name: '',
    surname: '',
    address: '',
    city: '',
    postal_code: '',
    phone_number: '',
  });

  const [errors, setErrors] = useState<Record<string, string>>({});

  const validateForm = () => {
    const newErrors: Record<string, string> = {};

    // Email validation
    if (!formData.email_address.match(/^[^\s@]+@[^\s@]+\.[^\s@]+$/)) {
      newErrors.email_address = 'Please enter a valid email address';
    }

    // Password validation
    if (formData.password.length < 8) {
      newErrors.password = 'Password must be at least 8 characters long';
    }

    if (formData.password !== formData.confirmPassword) {
      newErrors.confirmPassword = 'Passwords do not match';
    }

    // Phone number validation
    if (!formData.phone_number.match(/^\+?[\d\s-]{9,}$/)) {
      newErrors.phone_number = 'Please enter a valid phone number';
    }

    // Postal code validation
    if (!formData.postal_code.match(/^[\d\s-]{4,10}$/)) {
      newErrors.postal_code = 'Please enter a valid postal code';
    }

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const registerMutation = useMutation({
    mutationFn: (userData: Omit<typeof formData, 'confirmPassword'>) =>
      authService.register(userData),
    onSuccess: () => {
      navigate('/login', { replace: true });
    },
    onError: (error: any) => {
      if (error.response?.data?.message) {
        setErrors({ submit: error.response.data.message });
      } else {
        setErrors({ submit: 'An error occurred during registration. Please try again.' });
      }
    },
  });

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (!validateForm()) return;
    
    const { confirmPassword, ...registrationData } = formData;
    registerMutation.mutate(registrationData);
  };

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="max-w-2xl mx-auto">
        <div className="bg-white rounded-3xl shadow-lg p-8">
          <div className="text-center mb-8">
            <h1 className="text-4xl font-bold text-gray-900 mb-4">Create your account</h1>
            <p className="text-lg text-gray-600">
              Already have an account?{' '}
              <Link to="/login" className="text-primary-600 hover:text-primary-500 font-medium ">
                Sign in
              </Link>
            </p>
          </div>

          <form onSubmit={handleSubmit} className="space-y-6">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              {/* Email */}
              <div className="md:col-span-2">
                <label htmlFor="email_address" className="block text-sm font-medium text-gray-700 mb-2">
                  Email address
                </label>
                <input
                  id="email_address"
                  name="email_address"
                  type="email"
                  autoComplete="email"
                  required
                  className={`w-full rounded-xl border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-600 ${
                    errors.email_address ? 'border-red-500' : ''
                  }`}
                  value={formData.email_address}
                  onChange={(e) => {
                    setFormData({ ...formData, email_address: e.target.value });
                    if (errors.email_address) {
                      const { email_address, ...rest } = errors;
                      setErrors(rest);
                    }
                  }}
                />
                {errors.email_address && (
                  <p className="mt-1 text-sm text-red-600">{errors.email_address}</p>
                )}
              </div>

              {/* Password */}
              <div>
                <label htmlFor="password" className="block text-sm font-medium text-gray-700 mb-2">
                  Password
                </label>
                <input
                  id="password"
                  name="password"
                  type="password"
                  autoComplete="new-password"
                  required
                  className={`w-full rounded-xl border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-600 ${
                    errors.password ? 'border-red-500' : ''
                  }`}
                  value={formData.password}
                  onChange={(e) => {
                    setFormData({ ...formData, password: e.target.value });
                    if (errors.password) {
                      const { password, ...rest } = errors;
                      setErrors(rest);
                    }
                  }}
                />
                {errors.password && (
                  <p className="mt-1 text-sm text-red-600">{errors.password}</p>
                )}
              </div>

              {/* Confirm Password */}
              <div>
                <label htmlFor="confirmPassword" className="block text-sm font-medium text-gray-700 mb-2">
                  Confirm Password
                </label>
                <input
                  id="confirmPassword"
                  name="confirmPassword"
                  type="password"
                  autoComplete="new-password"
                  required
                  className={`w-full rounded-xl border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-600 ${
                    errors.confirmPassword ? 'border-red-500' : ''
                  }`}
                  value={formData.confirmPassword}
                  onChange={(e) => {
                    setFormData({ ...formData, confirmPassword: e.target.value });
                    if (errors.confirmPassword) {
                      const { confirmPassword, ...rest } = errors;
                      setErrors(rest);
                    }
                  }}
                />
                {errors.confirmPassword && (
                  <p className="mt-1 text-sm text-red-600">{errors.confirmPassword}</p>
                )}
              </div>

              {/* Name */}
              <div>
                <label htmlFor="name" className="block text-sm font-medium text-gray-700 mb-2">
                  First Name
                </label>
                <input
                  id="name"
                  name="name"
                  type="text"
                  autoComplete="given-name"
                  required
                  className="w-full rounded-xl border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-600"
                  value={formData.name}
                  onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                />
              </div>

              {/* Surname */}
              <div>
                <label htmlFor="surname" className="block text-sm font-medium text-gray-700 mb-2">
                  Last Name
                </label>
                <input
                  id="surname"
                  name="surname"
                  type="text"
                  autoComplete="family-name"
                  required
                  className="w-full rounded-xl border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-600"
                  value={formData.surname}
                  onChange={(e) => setFormData({ ...formData, surname: e.target.value })}
                />
              </div>

              {/* Address */}
              <div className="md:col-span-2">
                <label htmlFor="address" className="block text-sm font-medium text-gray-700 mb-2">
                  Address
                </label>
                <input
                  id="address"
                  name="address"
                  type="text"
                  autoComplete="street-address"
                  required
                  className="w-full rounded-xl border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-600"
                  value={formData.address}
                  onChange={(e) => setFormData({ ...formData, address: e.target.value })}
                />
              </div>

              {/* City */}
              <div>
                <label htmlFor="city" className="block text-sm font-medium text-gray-700 mb-2">
                  City
                </label>
                <input
                  id="city"
                  name="city"
                  type="text"
                  autoComplete="address-level2"
                  required
                  className="w-full rounded-xl border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-600"
                  value={formData.city}
                  onChange={(e) => setFormData({ ...formData, city: e.target.value })}
                />
              </div>

              {/* Postal Code */}
              <div>
                <label htmlFor="postal_code" className="block text-sm font-medium text-gray-700 mb-2">
                  Postal Code
                </label>
                <input
                  id="postal_code"
                  name="postal_code"
                  type="text"
                  autoComplete="postal-code"
                  required
                  className={`w-full rounded-xl border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-600 ${
                    errors.postal_code ? 'border-red-500' : ''
                  }`}
                  value={formData.postal_code}
                  onChange={(e) => {
                    setFormData({ ...formData, postal_code: e.target.value });
                    if (errors.postal_code) {
                      const { postal_code, ...rest } = errors;
                      setErrors(rest);
                    }
                  }}
                />
                {errors.postal_code && (
                  <p className="mt-1 text-sm text-red-600">{errors.postal_code}</p>
                )}
              </div>

              {/* Phone Number */}
              <div>
                <label htmlFor="phone_number" className="block text-sm font-medium text-gray-700 mb-2">
                  Phone Number
                </label>
                <input
                  id="phone_number"
                  name="phone_number"
                  type="tel"
                  autoComplete="tel"
                  required
                  className={`w-full rounded-xl border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 text-gray-600 ${
                    errors.phone_number ? 'border-red-500' : ''
                  }`}
                  value={formData.phone_number}
                  onChange={(e) => {
                    setFormData({ ...formData, phone_number: e.target.value });
                    if (errors.phone_number) {
                      const { phone_number, ...rest } = errors;
                      setErrors(rest);
                    }
                  }}
                />
                {errors.phone_number && (
                  <p className="mt-1 text-sm text-red-600">{errors.phone_number}</p>
                )}
              </div>
            </div>

            {errors.submit && (
              <div className="rounded-xl bg-red-50 p-4 text-sm text-red-600">
                {errors.submit}
              </div>
            )}

            <div>
              <button
                type="submit"
                disabled={registerMutation.isPending}
                className="w-full rounded-xl bg-primary-600 py-3 px-4 text-white font-medium hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
              >
                {registerMutation.isPending ? (
                  <div className="flex items-center justify-center">
                    <div className="h-5 w-5 border-2 border-white border-t-transparent rounded-full animate-spin mr-2"></div>
                    Creating account...
                  </div>
                ) : (
                  'Create account'
                )}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Register; 