import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import {
  CreatePackageOrderRequest,
  CreateTicketOrderRequest,
} from '../../models/festival/festival.model';

@Injectable({
  providedIn: 'root',
})
export class OrderService {
  private apiUrl = 'http://localhost:4000';

  constructor(private http: HttpClient) {}

  createTicketOrder(
    festivalId: number,
    request: CreateTicketOrderRequest,
  ): Observable<{ orderId: number }> {
    return this.http.post<{ orderId: number }>(
      `${this.apiUrl}/festival/${festivalId}/order/ticket`,
      request,
    );
  }

  createPackageOrder(
    festivalId: number,
    request: CreatePackageOrderRequest,
  ): Observable<{ orderId: number }> {
    return this.http.post<{ orderId: number }>(
      `${this.apiUrl}/festival/${festivalId}/order/package`,
      request,
    );
  }
}
