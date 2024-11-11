import { Injectable } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';

@Injectable({
  providedIn: 'root',
})
export class TokenService {
  private jwtHelper = new JwtHelperService();

  constructor() {}

  decodeToken(): any {
    const token = localStorage.getItem('token');
    if (token) {
      return this.jwtHelper.decodeToken(token);
    }
    return null;
  }

  getUserDetailsFromToken(): { name?: string; id?: number } {
    const decodedToken = this.decodeToken();
    if (decodedToken) {
      return {
        name: decodedToken[
          'http://schemas.xmlsoap.org/ws/2005/05/identity/claims/givenname'
        ],
        id: decodedToken[
          'http://schemas.xmlsoap.org/ws/2005/05/identity/claims/nameidentifier'
        ]
          ? parseInt(
              decodedToken[
                'http://schemas.xmlsoap.org/ws/2005/05/identity/claims/nameidentifier'
              ]
            )
          : undefined,
      };
    }
    return {};
  }
}
