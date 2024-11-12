import { City } from './city.model';

export interface Address {
  id: number;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string | null;
  street: string;
  number: string;
  apartmentSuite?: string | null;
  cityId: number;
  city: City;
}
