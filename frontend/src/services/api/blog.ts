import { api } from './config';
import type {
  BlogPost,
  BlogComment,
  PaginatedResponse,
  ApiResponse,
  PaginationParams,
  SortParams,
} from '../../types/api';

interface BlogPostsQueryParams extends PaginationParams, SortParams {
  ids?: string[];
  'date-time-from'?: string;
  'date-time-to'?: string;
  'author-ids'?: string[];
  tags?: string[];
  visible?: boolean;
}

interface BlogCommentsQueryParams extends PaginationParams, SortParams {
  ids?: string[];
  user_ids?: string[];
  date_time_from?: string;
  date_time_to?: string;
}

export const blogService = {
  // Blog Posts
  async getBlogPosts(params?: BlogPostsQueryParams): Promise<PaginatedResponse<BlogPost>> {
    const response = await api.get('/posts', { params });
    return response.data;
  },

  async getBlogPostById(id: string): Promise<{ blog_post: BlogPost }> {
    const response = await api.get(`/posts/${id}`);
    return response.data;
  },

  async createBlogPost(post: Partial<BlogPost>): Promise<ApiResponse> {
    const response = await api.post('/posts', post);
    return response.data;
  },

  async updateBlogPost(id: string, post: Partial<BlogPost>): Promise<ApiResponse> {
    const response = await api.put(`/posts/${id}`, post);
    return response.data;
  },

  async deleteBlogPost(id: string): Promise<ApiResponse> {
    const response = await api.delete(`/posts/${id}`);
    return response.data;
  },

  async getBlogPostTags(id: string, sortFields?: string[]): Promise<{ items: string[] }> {
    const response = await api.get(`/posts/tags/${id}`, {
      params: { sort_fields: sortFields },
    });
    return response.data;
  },

  // Comments
  async getBlogPostComments(
    postId: string,
    params?: BlogCommentsQueryParams
  ): Promise<PaginatedResponse<BlogComment>> {
    const response = await api.get(`/posts/${postId}/comments`, { params });
    return response.data;
  },

  async createBlogPostComment(
    postId: string,
    comment: { description: string }
  ): Promise<ApiResponse> {
    const response = await api.post(`/posts/${postId}/comments`, comment);
    return response.data;
  },

  async deleteBlogPostComment(commentId: string): Promise<ApiResponse> {
    const response = await api.delete(`/comments/${commentId}`);
    return response.data;
  },

  // Likes
  async getBlogPostLikes(postId: string): Promise<{ totalCount: number }> {
    const response = await api.get(`/posts/${postId}/likes`);
    return response.data;
  },

  async likeBlogPost(postId: string): Promise<ApiResponse> {
    const response = await api.post(`/posts/${postId}/likes`);
    return response.data;
  },

  async unlikeBlogPost(postId: string): Promise<ApiResponse> {
    const response = await api.delete(`/posts/${postId}/likes`);
    return response.data;
  },
}; 