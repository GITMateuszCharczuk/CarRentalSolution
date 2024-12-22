// Common Types
export interface ApiResponse {
  success: boolean;
  message: string;
}

export interface PaginatedResponse<T> extends ApiResponse {
  current_page: number;
  page_size: number;
  total_items: number;
  total_pages: number;
  items: T[];
}

// Auth Types
export interface LoginRequest {
  email: string;
  password: string;
}

export interface LoginResponse extends ApiResponse {
  token: string;
  refresh_token: string;
  roles: string[];
}

export interface RegisterRequest {
  email_address: string;
  password: string;
  name: string;
  surname: string;
  address: string;
  city: string;
  postal_code: string;
  phone_number: string;
}

// Car Types
export interface CarOffer {
  id: string;
  heading: string;
  shortDescription: string;
  urlHandle: string;
  featuredImageUrl: string;
  imageUrls: string[];
  visible: boolean;
  tags: string[];
  engineDetails: string;
  gearboxDetails: string;
  driveDetails: string;
  horsepower: string;
  yearOfProduction: number;
  oneNormalDayPrice: number;
  oneWeekendDayPrice: number;
  oneWeekPrice: number;
  oneMonthPrice: number;
  publishedDate?: string;
}

export interface CreateCarOfferRequest {
  heading: string;
  shortDescription: string;
  urlHandle: string;
  featuredImageUrl?: string;
  imageUrls?: string[];
  visible?: boolean;
  tags?: string[];
  engineDetails?: string;
  gearboxDetails?: string;
  driveDetails?: string;
  horsepower?: string;
  yearOfProduction?: number;
  oneNormalDayPrice?: number;
  oneWeekendDayPrice?: number;
  oneWeekPrice?: number;
  oneMonthPrice?: number;
}

export interface UpdateCarOfferRequest extends CreateCarOfferRequest {
  publishedDate?: string;
}

export interface CarOrder {
  id: string;
  userId: string;
  carOfferId: string;
  startDate: string;
  endDate: string;
  deliveryLocation: string;
  returnLocation: string;
  numOfDrivers: number;
  totalCost: number;
  status: string;
}

export interface CreateCarOrderRequest {
  carOfferId: string;
  startDate: string;
  endDate: string;
  deliveryLocation?: string;
  returnLocation?: string;
  numOfDrivers?: number;
  totalCost?: number;
  status?: string;
}

export interface UpdateCarOrderRequest extends CreateCarOrderRequest {
  userId: string;
}

// Blog Types
export interface BlogPost {
  id: string;
  heading: string;
  pageTitle: string;
  content: string;
  shortDescription: string;
  featuredImageUrl: string;
  urlHandle: string;
  author: string;
  publishedDate: string;
  visible: boolean;
  tags: string[];
}

export interface CreateBlogPostRequest {
  heading: string;
  pageTitle: string;
  content: string;
  shortDescription: string;
  urlHandle: string;
  featuredImageUrl?: string;
  tags?: string[];
  visible?: boolean;
}

export interface UpdateBlogPostRequest extends CreateBlogPostRequest {
  publishedDate: string;
  jwtToken: any; // Using any for now as the API spec doesn't provide details
}

export interface BlogComment {
  id: string;
  blogPostId: string;
  userId: string;
  description: string;
  createdAt: string;
}

export interface CreateBlogPostCommentRequest {
  description: string;
}

// User Types
export interface UserInfo {
  id: string;
  email_address: string;
  name: string;
  surname: string;
  address: string;
  city: string;
  postal_code: string;
  phone_number: string;
  roles: string[];
}

export interface ModifyUserRequest {
  user_id?: string;
  email_address?: string;
  name: string;
  surname: string;
  address: string;
  city: string;
  postal_code: string;
  phone_number: string;
  roles?: string[];
}

// Email Types
export interface SendEmailRequest {
  from: string;
  subject: string;
  body: string;
}

export interface SendInternalEmailRequest {
  to: string;
  subject: string;
  body: string;
}

// Query Parameters
export interface PaginationParams {
  page_size?: number;
  current_page?: number;
}

export interface SortParams {
  sort_fields?: string[];
}

export interface CarOffersQueryParams extends PaginationParams, SortParams {
  ids?: string[];
  date_time_from?: string;
  date_time_to?: string;
  tags?: string[];
  visible?: boolean;
}

export interface CarOrdersQueryParams extends PaginationParams, SortParams {
  start_date?: string;
  end_date?: string;
  user_id?: string;
  car_offer_id?: string;
  statuses?: string[];
  date_filter_type?: string;
}

export interface BlogPostsQueryParams extends PaginationParams, SortParams {
  ids?: string[];
  'date-time-from'?: string;
  'date-time-to'?: string;
  'author-ids'?: string[];
  tags?: string[];
  visible?: boolean;
}

export interface BlogCommentsQueryParams extends PaginationParams, SortParams {
  ids?: string[];
  user_ids?: string[];
  date_time_from?: string;
  date_time_to?: string;
} 