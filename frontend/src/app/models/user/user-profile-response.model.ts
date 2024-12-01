import { AddressResponse } from '../common/address-response.model';
import { ImageResponse } from '../festival/festival.model';

export interface UserProfileResponse {
  username: string;
  email: string;
  role: string;
  firstName: string;
  lastName: string;
  dateOfBirth: string;
  phoneNumber: string;
  address?: AddressResponse | null;
  imageURL?: string | null;
}
