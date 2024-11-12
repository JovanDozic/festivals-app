import { Country } from './country.model';

export interface City {
  id: number;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string | null;
  name: string;
  postalCode: string;
  countryId: number;
  country: Country;
}
