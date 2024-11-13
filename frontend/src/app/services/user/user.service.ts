import { Injectable } from '@angular/core';
import { UserProfileResponse } from '../../models/user/user-profile-response.model';
import { map, Observable, tap } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { UpdateUserProfileRequest } from '../../models/user/update-user-profile-request.model';

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

  updateUserProfile(updatedProfile: UpdateUserProfileRequest) {
    console.log('Updating profile', updatedProfile);
    return this.http
      .put<void>(`${this.apiUrl}/user/profile`, {
        firstName: updatedProfile.firstName,
        lastName: updatedProfile.lastName,
        dateOfBirth: formatDate(updatedProfile.dateOfBirth),
        phoneNumber: updatedProfile.phoneNumber,
      })
      .pipe(
        tap(() => {
          console.log('Profile updated');
        })
      );
  }
}

function formatDate(date: Date): string {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
}
