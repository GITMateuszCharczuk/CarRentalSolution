import { api } from './config';
import type { LoginRequest, LoginResponse, RegisterRequest, ApiResponse } from '../../types/api';

export const authService = {
  async login(credentials: LoginRequest): Promise<LoginResponse> {
    const response = await api.post<LoginResponse>('/login', credentials);
    return response.data;
  },

  async register(userData: RegisterRequest): Promise<ApiResponse> {
    const response = await api.post<ApiResponse>('/register', userData);
    return response.data;
  },

  async validateToken(token: string): Promise<{ valid: boolean; roles: string[] }> {
    const response = await api.get('/token/validate', {
      params: { token },
    });
    return response.data;
  },

  async refreshToken(token: string): Promise<{ token: string }> {
    const response = await api.post('/token/refresh', null, {
      params: { token },
    });
    return response.data;
  },

  async getUserInfo(token: string, userId?: string): Promise<ApiResponse> {
    const response = await api.get('/user/info', {
      params: { token, id: userId },
    });
    return response.data;
  },

  async modifyUser(token: string, userData: Partial<RegisterRequest>): Promise<ApiResponse> {
    const response = await api.put('/user', userData, {
      params: { token },
    });
    return response.data;
  },
}; 