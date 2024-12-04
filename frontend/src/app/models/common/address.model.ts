export interface CityRequest {
  name: string;
  postalCode: string;
  countryISO3: string;
}

export interface CountryResponse {
  id: number;
  niceName: string;
  iso: string;
  iso3: string;
}
