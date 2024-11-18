import { AddressResponse } from '../common/address-response.model';

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
  address?: AddressResponse | null;
  images: ImageResponse[];
}

export interface FestivalsResponse {
  festivals: Festival[];
}

export interface ImageResponse {
  url: string;
}
