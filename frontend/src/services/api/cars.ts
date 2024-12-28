import { api } from './config';
import type {
  CarOffer,
  CarOrder,
  CarOffersQueryParams,
  CarOrdersQueryParams,
  PaginatedResponse,
  ApiResponse,
  CarOfferTag,
  CarOfferImage,
  ListResponse,
  CreateCarOrderRequest,
} from '../../types/api';

export const carService = {
  // Car Offers
  async getCarOffers({
    page_size,
    current_page,
    tags,
    date_time_from,
    date_time_to,
    sort_fields,
  }: {
    page_size: number;
    current_page: number;
    tags?: string[];
    date_time_from?: string;
    date_time_to?: string;
    sort_fields?: string[];
  }): Promise<PaginatedResponse<CarOffer>> {
    const response = await api.get<PaginatedResponse<CarOffer>>('/car-offers', {
      params: {
        page_size,
        current_page,
        tags: tags?.join(','),
        date_time_from,
        date_time_to,
        sort_fields: sort_fields?.join(','),
      },
    });
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

  async getCarOfferTags(id?: string, sortFields?: string[]): Promise<ListResponse<CarOfferTag>> {
    const response = await api.get('/car-offers/tags', {
      params: { 
        car_offer_id: id !== '' ? id : undefined,
        sort_fields: sortFields?.join(',')
      },
    });
    return response.data;
  },

  // Car Images
  async addImageToCarOffer(offerId: string, imageId: string): Promise<ApiResponse> {
    const response = await api.post(`/car-offers/images/${offerId}/${imageId}`);
    return response.data;
  },

  async getCarOfferImages(offerId: string): Promise<ListResponse<CarOfferImage>> {
    const response = await api.get(`/car-offers/images/${offerId}`);
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

  createCarOrder: async (orderData: CreateCarOrderRequest) => {
    const { data } = await api.post<CarOrder>('/car-rental/api/car-orders', orderData);
    return data;
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