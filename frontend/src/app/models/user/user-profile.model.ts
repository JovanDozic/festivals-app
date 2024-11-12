import { Address } from '../common/address.model';
import { Image } from '../common/image';
import { User } from './user.model';

export interface UserProfile {
  id: number;
  firstName: string;
  lastName: string;
  dateOfBirth: string;
  phoneNumber: string;
  userId: number;
  user: User;
  addressId?: number | null;
  address?: Address | null;
  imageId?: number | null;
  image?: Image | null;
}
