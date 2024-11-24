import { AddressResponse } from '../common/address-response.model';
import {
  CreateAddressRequest,
  UpdateAddressRequest,
} from '../common/create-address-request.model';

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

export interface Employee {
  id: number;
  username: string;
  email: string;
  firstName: string;
  lastName: string;
  dateOfBirth: string;
  phoneNumber: string;
}
