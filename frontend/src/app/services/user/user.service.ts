import { Injectable } from '@angular/core';
import { UserProfileResponse } from '../../models/user/user-profile-response.model';
import { map, Observable, tap } from 'rxjs';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  private apiUrl = 'http://localhost:4000';

  constructor(private http: HttpClient) {}

  getUserProfile(): Observable<UserProfileResponse> {
    return this.http
      .get<{ userProfile: UserProfileResponse }>(`${this.apiUrl}/user/profile`)
      .pipe(
        tap((response) => {
          console.log('getUserProfile response', response);
        }),
        // Map the response to only extract the userProfile property
        map((response) => response.userProfile)
      );
  }
}
