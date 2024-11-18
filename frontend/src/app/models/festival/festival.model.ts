import { AddressResponse } from '../common/address-response.model';

export interface Festival {
  ID: number;
  Name: string;
  Description: string;
  StartDate: string;
  EndDate: string;
  Capacity: number;
  Status: string;
  StoreStatus: string;
  AddressID: number;
  Address?: AddressResponse | null;
}

export interface FestivalsResponse {
  Festivals: Festival[];
}
