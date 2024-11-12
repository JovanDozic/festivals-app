import { Injectable } from '@angular/core';
import { UserProfileResponse } from '../../models/user/user-profile-response.model';
import { Observable, tap } from 'rxjs';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  private apiUrl = 'http://localhost:4000';

  constructor(private http: HttpClient) {}

  getUserProfile(): Observable<UserProfileResponse> {
    console.log('getUserProfile');

    return this.http
      .get<UserProfileResponse>(`${this.apiUrl}/user/profile`)
      .pipe(
        tap((response) => {
          console.log('getUserProfile response', response);
        })
      );
  }
}
