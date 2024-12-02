import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, tap } from 'rxjs';
import { jwtDecode } from 'jwt-decode';
import { LoginResponse } from '../../models/auth/login-response.model';

export interface JwtPayload {
  username: string;
  role: string;
  exp?: number;
  iat?: number;
}

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  private apiUrl = 'http://localhost:4000';

  constructor(private http: HttpClient) {}

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

  // Use the JwtPayload interface with jwtDecode
  private getDecodedToken(): JwtPayload | null {
    const token = this.getToken();
    if (token) {
      try {
        return jwtDecode<JwtPayload>(token);
      } catch (error) {
        console.error('Invalid token', error);
        return null;
      }
    }
    return null;
  }

  getUserRole(): string | null {
    const decodedToken = this.getDecodedToken();
    return decodedToken ? decodedToken.role : null;
  }

  getUsername(): string | null {
    const decodedToken = this.getDecodedToken();
    return decodedToken ? decodedToken.username : null;
  }

  login(
    credentials: {
      username: string;
      password: string;
    },
    reload = true,
  ): Observable<LoginResponse> {
    return this.http
      .post<LoginResponse>(`${this.apiUrl}/user/login`, credentials)
      .pipe(
        tap((response) => {
          this.setToken(response.token);
          if (reload) window.location.reload();
        }),
      );
  }

  changePassword(oldPassword: string, newPassword: string): Observable<void> {
    return this.http.put<void>(`${this.apiUrl}/user/change-password`, {
      oldPassword,
      newPassword,
    });
  }

  registerAttendee(credentials: {
    username: string;
    email: string;
    password: string;
  }): Observable<void> {
    return this.http.post<void>(
      `${this.apiUrl}/user/register-attendee`,
      credentials,
    );
  }
}
