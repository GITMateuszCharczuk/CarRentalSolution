import { api } from './config';
import type { ApiResponse } from '../../types/api';

export const fileService = {
  async uploadFile(file: File): Promise<{ id: string; message: string }> {
    const formData = new FormData();
    formData.append('file', file);

    const response = await api.post('/files', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    return response.data;
  },

  async getFile(fileId: string, download?: boolean): Promise<Blob> {
    const response = await api.get('/files/get', {
      params: { file_id: fileId, download },
      responseType: 'blob',
    });
    return response.data;
  },

  async deleteFile(fileId: string): Promise<ApiResponse> {
    const response = await api.delete(`/files/delete/${fileId}`);
    return response.data;
  },

  // Helper function to get a temporary URL for a file
  getFileUrl(fileId: string): string {
    return `${api.defaults.baseURL}/files/get?file_id=${fileId}`;
  },

  // Helper function to get a download URL for a file
  getDownloadUrl(fileId: string): string {
    return `${api.defaults.baseURL}/files/get?file_id=${fileId}&download=true`;
  },
}; 