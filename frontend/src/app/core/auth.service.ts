import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, tap } from 'rxjs';

export interface AuthResponse {
  token: string;
}

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
    console.log('isLoggedIn', !!this.getToken());
    return !!this.getToken();
  }

  logout(): void {
    localStorage.removeItem('authToken');
  }

  constructor(private http: HttpClient) {}

  login(credentials: {
    username: string;
    password: string;
  }): Observable<AuthResponse> {
    console.log('trying login with ', credentials);
    return this.http
      .post<AuthResponse>(`${this.apiUrl}/user/login`, credentials)
      .pipe(
        tap((response) => {
          this.setToken(response.token);
        })
      );
  }
}
