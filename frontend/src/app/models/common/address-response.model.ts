export interface AddressResponse {
  addressId?: number;
  street: string;
  number: string;
  apartmentSuite: string;
  city: string;
  postalCode: string;
  country: string;
  countryISO3: string;
  countryISO2: string;
  niceName?: string;
}
