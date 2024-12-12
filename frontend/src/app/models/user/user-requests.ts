export interface CreateUpdateUserProfileRequest {
  firstName: string;
  lastName: string;
  dateOfBirth: Date;
  phoneNumber: string;
}

export interface CreateProfileRequest {
  firstName: string;
  lastName: string;
  dateOfBirth: string;
  phoneNumber: string;
}

export interface UpdateStaffProfileRequest {
  username: string;
  firstName: string;
  lastName: string;
  dateOfBirth: Date;
  phoneNumber: string;
}

export interface UpdateStaffEmailRequest {
  username: string;
  email: string;
}
