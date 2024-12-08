import { AddressResponse } from '../common/address-response.model';
import { CityRequest } from '../common/address.model';
import {
  CreateAddressRequest,
  UpdateAddressRequest,
} from '../common/create-address-request.model';
import { CreateProfileRequest } from '../user/user-profile-request.model';
import { UserProfileResponse } from '../user/user-profile-response.model';

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
  id: number;
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
export interface UpdateFestivalRequest {
  id: number;
  name: string;
  description: string;
  startDate: string;
  endDate: string;
  capacity: number;
  address: UpdateAddressRequest;
}

export interface EmployeesResponse {
  festivalId: number;
  employees: Employee[];
}

export interface ItemsResponse {
  festivalId: number;
  items: ItemCurrentPrice[];
}

export interface Employee {
  id: number;
  username: string;
  email: string;
  firstName: string;
  lastName: string;
  dateOfBirth: string;
  phoneNumber: string;
}

export interface CreateStaffRequest {
  username: string;
  password: string;
  email: string;
  userProfile: CreateProfileRequest;
}

export interface CreateStaffResponse {
  username: string;
  userId: number;
}

export interface ItemCurrentPrice {
  itemId: number;
  priceListItemId: number;
  name: string;
  description: string;
  availableNumber: number;
  remainingNumber: number;
  price: number;
  isFixed: boolean;
  dateFrom: string;
  dateTo: string;
}

export interface FestivalPropCountResponse {
  festivalId: number;
  count: number;
}

export interface CreateItemRequest {
  name: string;
  description: string;
  availableNumber: number;
  type: string;
}

export interface CreateItemPriceRequest {
  itemId: number;
  price: number;
  isFixed: boolean;
  dateFrom?: string;
  dateTo?: string;
}

export interface VariablePrice {
  price: number;
  dateFrom: Date;
  dateTo: Date;
}

export interface Item {
  id: number;
  name: string;
  description: string;
  type: string;
  availableNumber: number;
  remainingNumber: number;
  priceListItems: PriceListItem[];
}

export interface PriceListItem {
  id: number;
  price: number;
  isFixed: boolean;
  dateFrom?: string;
  dateTo?: string;
}

export interface AddTransportConfigRequest {
  itemId: number;
  transportType: string;
  departureCity: CityRequest;
  arrivalCity: CityRequest;
  departureTime: string;
  arrivalTime: string;
  returnDepartureTime: string;
  returnArrivalTime: string;
}

export interface AddCampConfigRequest {
  itemId: number;
  campName: string;
  imageURL: string;
  equipmentList: EquipmentDTO[];
}

export interface EquipmentDTO {
  name: string;
}

export interface TransportAddonDTO {
  priceListItemId: number;
  priceListId: number;
  itemId: number;
  itemName: string;
  itemDescription: string;
  itemType: string;
  itemAvailableNumber: number;
  itemRemainingNumber: number;
  dateFrom: Date | null;
  dateTo: Date | null;
  isFixed: boolean;
  price: number;
  packageAddonCategory: string;
  transportType: string;
  departureTime: Date;
  arrivalTime: Date;
  returnDepartureTime: Date;
  returnArrivalTime: Date;
  departureCityId: number;
  departureCityName: string;
  departurePostalCode: string;
  departureCountryISO3: string;
  departureCountryISO: string;
  departureCountryNiceName: string;
  arrivalCityId: number;
  arrivalCityName: string;
  arrivalPostalCode: string;
  arrivalCountryISO3: string;
  arrivalCountryISO: string;
  arrivalCountryNiceName: string;
}

export interface GeneralAddonDTO {
  priceListItemId: number;
  priceListId: number;
  itemId: number;
  itemName: string;
  itemDescription: string;
  itemType: string;
  itemAvailableNumber: number;
  itemRemainingNumber: number;
  dateFrom: Date | null;
  dateTo: Date | null;
  isFixed: boolean;
  price: number;
  packageAddonCategory: string;
}

export interface CampAddonDTO {
  priceListItemId: number;
  priceListId: number;
  itemId: number;
  itemName: string;
  itemDescription: string;
  itemType: string;
  itemAvailableNumber: number;
  itemRemainingNumber: number;
  dateFrom: Date | null;
  dateTo: Date | null;
  isFixed: boolean;
  price: number;
  packageAddonCategory: string;
  campName: string;
  imageUrl: string;
  equipmentNames: string;
}

export interface CreateTicketOrderRequest {
  ticketTypeId: number;
  totalPrice: number;
}

export interface TransportType {
  value: string;
  viewValue: string;
}

export interface CreatePackageOrderRequest {
  ticketTypeId: number;
  transportAddonId: number | null;
  campAddonId: number | null;
  generalAddonIds: number[];
  totalPrice: number;
}

export interface OrderDTO {
  orderId: number;
  orderType: string;
  timestamp: Date;
  totalPrice: number;
  ticket: ItemCurrentPrice;
  transportAddon?: TransportAddonDTO;
  campAddon?: CampAddonDTO;
  generalAddons: GeneralAddonDTO[];
  festival: Festival;
  attendee: UserProfileResponse;
  braceletStatus?: string;
  festivalTicketId: number;
  bracelet?: BraceletDTO;
}

export interface BraceletDTO {
  braceletId: number;
  barcodeNumber: string;
  balance: number;
  status: string;
  employee: UserProfileResponse;
}

export interface OrderPreviewDTO {
  orderId: number;
  orderType: string;
  timestamp: Date;
  totalPrice: number;
  festival: Festival;
  username: string; // this is attendee's username
  attendee: UserProfileResponse;
  braceletStatus?: string;
  braceletId?: number;
  festivalTicketId: number;
}

export interface IssueBraceletRequest {
  orderId: number;
  pin: string;
  barcodeNumber: string;
  festivalTicketId: number;
  attendeeUsername: string;
}

export interface IssueBraceletResponse {
  braceletId: number;
  shippingAddress: AddressResponse;
}

export interface ActivateBraceletRequest {
  braceletId: number;
  pin: string;
}

export interface TopUpBraceletRequest {
  braceletId: number;
  amount: number;
}

export interface ActivateBraceletHelpRequest {
  braceletId: number;
  barcodeNumberUser: string;
  pinUser: string;
  issueDescription: string;
  imageURL: string;
}

export interface ActivationHelpRequestDTO {
  activationHelpRequestId: number;
  userEnteredPIN: string;
  userEnteredBarcode: string;
  issueDescription: string;
  imageURL: string;
  status: string;
  bracelet: BraceletDTO;
  attendee: UserProfileResponse;
  employee: UserProfileResponse;
}
