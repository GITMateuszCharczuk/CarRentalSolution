import { api } from './config';
import type {
  UserInfo,
  GetAllUsersResponse,
  UserManagementQueryParams,
  ModifyUserRequest,
  ApiResponse,
} from '../../types/api';

export const userService = {
  // Get all users with pagination and sorting
  async getUsers(params?: UserManagementQueryParams): Promise<GetAllUsersResponse> {
    const response = await api.get('/users', { params });
    return response.data;
  },

  // Get user info by ID
  async getUserInfo(userId?: string, token?: string): Promise<{ user_info: UserInfo }> {
    const params: Record<string, string> = {};
    if (userId) params.id = userId;
    if (token) params.token = token;
    
    const response = await api.get('/user/info', { params });
    return response.data;
  },

  // Get current user's internal info
  async getUserInternalInfo(): Promise<{ user_info: UserInfo }> {
    const response = await api.get('/user/internal');
    return response.data;
  },

  // Modify user
  async modifyUser(userData: ModifyUserRequest): Promise<ApiResponse> {
    const response = await api.put('/user', userData);
    return response.data;
  },

  // Delete user
  async deleteUser(userId: string): Promise<ApiResponse> {
    const response = await api.delete(`/user/${userId}`);
    return response.data;
  },
}; 