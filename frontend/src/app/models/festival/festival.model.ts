import { AddressResponse } from '../common/address-response.model';
import { CreateAddressRequest } from '../common/create-address-request.model';

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
