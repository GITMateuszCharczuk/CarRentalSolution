import axios from 'axios';
import { store } from '../../store';
import { setToken, logout, selectAuthToken, selectRefreshToken } from '../../store/slices/authSlice';

const BASE_URL = 'http://localhost:8000/car-rental/api'; // Change this to your API URL

export const api = axios.create({
  baseURL: BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor for API calls
api.interceptors.request.use(
  (config) => {
    const state = store.getState();
    const token = selectAuthToken(state);
    
    if (token) {
      config.params = {
        ...config.params,
        token,
      };
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Response interceptor for API calls
api.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;

    // If the error status is 401 and there is no originalRequest._retry flag,
    // it means the token has expired and we need to refresh it
    if (error.response.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;

      try {
        const state = store.getState();
        const refreshToken = selectRefreshToken(state);
        console.log(refreshToken);
        const response = await axios.post(`${BASE_URL}/token/refresh`, null, {
          params: { token: refreshToken },
        });

        const { Token } = response.data;
        store.dispatch(setToken(Token));

        // Retry the original request with the new token
        originalRequest.params = {
          ...originalRequest.params,
          token: Token,
        };
        console.log(originalRequest.params);
        return api(originalRequest);
      } catch (refreshError) {
        // If refresh token fails, logout user and redirect to login
        store.dispatch(logout());
        return Promise.reject(refreshError);
      }
    }

    return Promise.reject(error);
  }
); 