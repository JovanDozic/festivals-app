import { AddressResponse } from '../common/address-response.model';
import { CityRequest } from '../common/address.model';
import {
  CreateAddressRequest,
  UpdateAddressRequest,
} from '../common/create-address-request.model';
import { CreateProfileRequest } from '../user/user-profile-request.model';

export interface Festival {
  id: number;
  name: string;
  description: string;
  startDate: string;
  endDate: string;
  capacity: number;
  status: string;
  storeStatus: string;
  addressID: number;
  address: AddressResponse | null;
  images: ImageResponse[];
}

export interface FestivalsResponse {
  festivals: Festival[];
}

export interface ImageResponse {
  url: string;
}
export interface CreateFestivalRequest {
  name: string;
  description: string;
  startDate: string;
  endDate: string;
  capacity: number;
  address: CreateAddressRequest;
}
export interface UpdateFestivalRequest {
  id: number;
  name: string;
  description: string;
  startDate: string;
  endDate: string;
  capacity: number;
  address: UpdateAddressRequest;
}

export interface EmployeesResponse {
  festivalId: number;
  employees: Employee[];
}

export interface ItemsResponse {
  festivalId: number;
  items: ItemCurrentPrice[];
}

export interface Employee {
  id: number;
  username: string;
  email: string;
  firstName: string;
  lastName: string;
  dateOfBirth: string;
  phoneNumber: string;
}

export interface CreateStaffRequest {
  username: string;
  password: string;
  email: string;
  userProfile: CreateProfileRequest;
}

export interface CreateStaffResponse {
  username: string;
  userId: number;
}

export interface ItemCurrentPrice {
  itemId: number;
  priceListItemId: number;
  name: string;
  description: string;
  availableNumber: number;
  remainingNumber: number;
  price: number;
  isFixed: boolean;
  dateFrom: string;
  dateTo: string;
}

export interface FestivalPropCountResponse {
  festivalId: number;
  count: number;
}

export interface CreateItemRequest {
  name: string;
  description: string;
  availableNumber: number;
  type: string;
}

export interface CreateItemPriceRequest {
  itemId: number;
  price: number;
  isFixed: boolean;
  dateFrom?: string;
  dateTo?: string;
}

export interface VariablePrice {
  price: number;
  dateFrom: Date;
  dateTo: Date;
}

export interface Item {
  id: number;
  name: string;
  description: string;
  type: string;
  availableNumber: number;
  remainingNumber: number;
  priceListItems: PriceListItem[];
}

export interface PriceListItem {
  id: number;
  price: number;
  isFixed: boolean;
  dateFrom?: string;
  dateTo?: string;
}

export interface AddTransportConfigRequest {
  itemId: number;
  transportType: string;
  departureCity: CityRequest;
  arrivalCity: CityRequest;
  departureTime: string;
  arrivalTime: string;
  returnDepartureTime: string;
  returnArrivalTime: string;
}
