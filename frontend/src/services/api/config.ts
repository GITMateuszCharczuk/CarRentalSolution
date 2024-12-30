import axios from 'axios';
import { store } from '../../store';
import { setToken, logout, selectAuthToken, selectRefreshToken } from '../../store/slices/authSlice';

const BASE_URL = 'http://localhost:8000/car-rental/api';

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

    if (error.response.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;

      try {
        const state = store.getState();
        const refreshToken = selectRefreshToken(state);
        const response = await axios.post(`${BASE_URL}/token/refresh`, null, {
          params: { token: refreshToken },
        });

        const { Token } = response.data;
        store.dispatch(setToken(Token));

        originalRequest.params = {
          ...originalRequest.params,
          token: Token,
        };
        return api(originalRequest);
      } catch (refreshError) {
        store.dispatch(logout());
        return Promise.reject(refreshError);
      }
    }

    return Promise.reject(error);
  }
); 