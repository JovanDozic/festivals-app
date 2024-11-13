import { Injectable } from '@angular/core';
import { UserProfileResponse } from '../../models/user/user-profile-response.model';
import { map, Observable, tap } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { CreateUpdateUserProfileRequest } from '../../models/user/user-profile-request.model';
import { CreateAddressRequest } from '../../models/user/create-address-request.model';

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

  updateUserProfile(updatedProfile: CreateUpdateUserProfileRequest) {
    return this.http
      .put<void>(`${this.apiUrl}/user/profile`, {
        firstName: updatedProfile.firstName,
        lastName: updatedProfile.lastName,
        dateOfBirth: formatDate(updatedProfile.dateOfBirth),
        phoneNumber: updatedProfile.phoneNumber,
      })
      .pipe(tap(() => {}));
  }

  createUserProfile(profile: CreateUpdateUserProfileRequest) {
    return this.http
      .post<void>(`${this.apiUrl}/user/profile`, {
        firstName: profile.firstName,
        lastName: profile.lastName,
        dateOfBirth: formatDate(profile.dateOfBirth),
        phoneNumber: profile.phoneNumber,
      })
      .pipe(tap(() => {}));
  }

  createAddress(address: CreateAddressRequest) {
    return this.http
      .post<void>(`${this.apiUrl}/user/profile/address`, address)
      .pipe(tap(() => {}));
  }
}

function formatDate(date: Date): string {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
}
