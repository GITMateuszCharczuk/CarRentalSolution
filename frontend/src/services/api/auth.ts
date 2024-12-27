import { api } from './config';
import type { LoginRequest, RegisterRequest, AuthResponse } from '../../types/api';

export const authService = {
  login: async (credentials: LoginRequest): Promise<AuthResponse> => {
    const response = await api.post('/login', credentials);
    return response.data;
  },

  register: async (userData: RegisterRequest): Promise<AuthResponse> => {
    const response = await api.post('/register', userData);
    return response.data;
  },

  refreshToken: async (refreshToken: string): Promise<AuthResponse> => {
    const response = await api.post('/token/refresh-token', { refresh_token: refreshToken });
    return response.data;
  },
}; 