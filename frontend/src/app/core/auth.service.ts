import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, tap } from 'rxjs';
import { jwtDecode } from 'jwt-decode';
import { LoginResponse } from '../models/auth/login-response.model';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  private apiUrl = 'http://localhost:4000';

  setToken(token: string): void {
    localStorage.setItem('authToken', token);
  }

  getToken(): string | null {
    return localStorage.getItem('authToken');
  }

  isLoggedIn(): boolean {
    return !!this.getToken();
  }

  logout(): void {
    localStorage.removeItem('authToken');
    window.location.reload();
  }

  getUserRole(): string | null {
    const token = this.getToken();
    if (token) {
      const decodedToken: any = jwtDecode(token);
      return decodedToken.role;
    }
    return null;
  }

  constructor(private http: HttpClient) {}

  login(
    credentials: {
      username: string;
      password: string;
    },
    reload: boolean = true
  ): Observable<LoginResponse> {
    return this.http
      .post<LoginResponse>(`${this.apiUrl}/user/login`, credentials)
      .pipe(
        tap((response) => {
          this.setToken(response.token);
          if (reload) window.location.reload();
        })
      );
  }

  changePassword(oldPassword: string, newPassword: string): Observable<void> {
    return this.http
      .put<void>(`${this.apiUrl}/user/change-password`, {
        oldPassword,
        newPassword,
      })
      .pipe(tap(() => {}));
  }

  registerAttendee(credentials: {
    username: string;
    email: string;
    password: string;
  }): Observable<void> {
    return this.http.post<void>(
      `${this.apiUrl}/user/register-attendee`,
      credentials
    );
  }
}
