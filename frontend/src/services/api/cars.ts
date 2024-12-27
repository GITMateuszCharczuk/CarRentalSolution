import { api } from './config';
import type {
  CarOffer,
  CarOrder,
  CarOffersQueryParams,
  CarOrdersQueryParams,
  PaginatedResponse,
  ApiResponse,
} from '../../types/api';

export const carService = {
  // Car Offers
  async getCarOffers(params?: CarOffersQueryParams): Promise<PaginatedResponse<CarOffer>> {
    const response = await api.get('/car-offers', { params });
    console.log(response.data);
    return response.data;
  },

  async getCarOfferById(id: string): Promise<{ car_offer: CarOffer }> {
    const response = await api.get(`/car-offers/${id}`);
    return response.data;
  },

  async createCarOffer(carOffer: Partial<CarOffer>): Promise<ApiResponse> {
    const response = await api.post('/car-offers', carOffer);
    return response.data;
  },

  async updateCarOffer(id: string, carOffer: Partial<CarOffer>): Promise<ApiResponse> {
    const response = await api.put(`/car-offers/${id}`, carOffer);
    return response.data;
  },

  async deleteCarOffer(id: string): Promise<ApiResponse> {
    const response = await api.delete(`/car-offers/${id}`);
    return response.data;
  },

  async getCarOfferTags(id: string, sortFields?: string[]): Promise<{ items: string[] }> {
    const response = await api.get(`/car-offers/tags/${id}`, {
      params: { sort_fields: sortFields },
    });
    return response.data;
  },

  // Car Images
  async addImageToCarOffer(offerId: string, imageId: string): Promise<ApiResponse> {
    const response = await api.post(`/car-offers/images/${offerId}/${imageId}`);
    return response.data;
  },

  async deleteImageFromCarOffer(carOfferId: string, imageId: string): Promise<ApiResponse> {
    const response = await api.delete(`/car-offers/images/${carOfferId}/${imageId}`);
    return response.data;
  },

  // Car Orders
  async getCarOrders(params?: CarOrdersQueryParams): Promise<PaginatedResponse<CarOrder>> {
    const response = await api.get('/car-orders', { params });
    return response.data;
  },

  async getCarOrderById(id: string): Promise<{ car_order: CarOrder }> {
    const response = await api.get(`/car-orders/${id}`);
    return response.data;
  },

  async createCarOrder(order: Partial<CarOrder>): Promise<ApiResponse> {
    const response = await api.post('/car-orders', order);
    return response.data;
  },

  async updateCarOrder(id: string, order: Partial<CarOrder>): Promise<ApiResponse> {
    const response = await api.put(`/car-orders/${id}`, order);
    return response.data;
  },

  async deleteCarOrder(id: string): Promise<ApiResponse> {
    const response = await api.delete(`/car-orders/${id}`);
    return response.data;
  },
}; 