import { Injectable } from '@angular/core';
import { UserProfileResponse } from '../../models/user/user-profile-response.model';
import { map, Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import {
  CreateUpdateUserProfileRequest,
  UpdateStaffEmailRequest,
  UpdateStaffProfileRequest,
} from '../../models/user/user-profile-request.model';
import {
  CreateAddressRequest,
  UpdateAddressRequest,
} from '../../models/common/create-address-request.model';
import {
  CreateStaffRequest,
  CreateStaffResponse,
} from '../../models/festival/festival.model';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  private apiUrl = 'http://localhost:4000';

  constructor(private http: HttpClient) {}

  getUserProfile(): Observable<UserProfileResponse> {
    return this.http
      .get<{ userProfile: UserProfileResponse }>(`${this.apiUrl}/user/profile`)
      .pipe(map((response) => response.userProfile));
  }

  updateUserProfilePhoto(imageURL: string) {
    return this.http.put<void>(`${this.apiUrl}/user/profile/photo`, {
      imageURL,
    });
  }

  updateUserProfile(updatedProfile: CreateUpdateUserProfileRequest) {
    return this.http.put<void>(`${this.apiUrl}/user/profile`, {
      firstName: updatedProfile.firstName,
      lastName: updatedProfile.lastName,
      dateOfBirth: formatDate(updatedProfile.dateOfBirth),
      phoneNumber: updatedProfile.phoneNumber,
    });
  }

  updateUserEmail(email: string) {
    return this.http.put<void>(`${this.apiUrl}/user/email`, {
      email,
    });
  }

  updateUserAddress(updatedAddress: UpdateAddressRequest) {
    return this.http.put<void>(
      `${this.apiUrl}/user/profile/address`,
      updatedAddress,
    );
  }

  createUserProfile(profile: CreateUpdateUserProfileRequest) {
    return this.http.post<void>(`${this.apiUrl}/user/profile`, {
      firstName: profile.firstName,
      lastName: profile.lastName,
      dateOfBirth: formatDate(profile.dateOfBirth),
      phoneNumber: profile.phoneNumber,
    });
  }

  createAddress(address: CreateAddressRequest) {
    return this.http.post<void>(`${this.apiUrl}/user/profile/address`, address);
  }

  registerEmployee(
    employee: CreateStaffRequest,
  ): Observable<CreateStaffResponse> {
    return this.http
      .post<{
        employee: CreateStaffResponse;
      }>(`${this.apiUrl}/organizer/employee`, employee)
      .pipe(map((response) => response.employee));
  }

  updateStaffProfile(updatedProfile: UpdateStaffProfileRequest) {
    return this.http.put<void>(`${this.apiUrl}/organizer/employee`, {
      username: updatedProfile.username,
      firstName: updatedProfile.firstName,
      lastName: updatedProfile.lastName,
      dateOfBirth: formatDate(updatedProfile.dateOfBirth),
      phoneNumber: updatedProfile.phoneNumber,
    });
  }

  updateStaffEmail(updatedProfile: UpdateStaffEmailRequest) {
    return this.http.put<void>(
      `${this.apiUrl}/organizer/employee/email`,
      updatedProfile,
    );
  }
}

function formatDate(date: Date): string {
  console.log(date);
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
}
