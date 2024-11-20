export interface CreateAddressRequest {
  street: string;
  number: string;
  apartmentSuite?: string | null;
  city: string;
  postalCode: string;
  countryISO3: string;
}
