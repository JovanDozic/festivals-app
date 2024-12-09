import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import {
  ActivateBraceletHelpRequest,
  ActivateBraceletRequest,
  ActivationHelpRequestDTO,
  CreatePackageOrderRequest,
  CreateTicketOrderRequest,
  IssueBraceletRequest,
  IssueBraceletResponse,
  OrderDTO,
  OrderPreviewDTO,
  TopUpBraceletRequest,
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

  getMyOrders(): Observable<OrderPreviewDTO[]> {
    return this.http.get<OrderPreviewDTO[]>(`${this.apiUrl}/order/attendee`);
  }

  getFestivalOrders(festivalId: number): Observable<OrderPreviewDTO[]> {
    return this.http.get<OrderPreviewDTO[]>(
      `${this.apiUrl}/festival/${festivalId}/order`,
    );
  }

  issueBracelet(
    request: IssueBraceletRequest,
  ): Observable<IssueBraceletResponse> {
    return this.http.post<IssueBraceletResponse>(
      `${this.apiUrl}/bracelet`,
      request,
    );
  }

  getMyBracelets(): Observable<OrderDTO[]> {
    return this.http.get<OrderDTO[]>(`${this.apiUrl}/bracelet/attendee`);
  }

  activateBracelet(request: ActivateBraceletRequest): Observable<void> {
    return this.http.put<void>(
      `${this.apiUrl}/bracelet/${request.braceletId}/activate`,
      request,
    );
  }

  topUpBracelet(request: TopUpBraceletRequest): Observable<void> {
    return this.http.put<void>(
      `${this.apiUrl}/bracelet/${request.braceletId}/top-up`,
      request,
    );
  }

  sendHelpRequest(request: ActivateBraceletHelpRequest): Observable<void> {
    return this.http.post<void>(
      `${this.apiUrl}/bracelet/${request.braceletId}/activate/help`,
      request,
    );
  }

  getHelpRequest(braceletId: number): Observable<ActivationHelpRequestDTO> {
    return this.http.get<ActivationHelpRequestDTO>(
      `${this.apiUrl}/bracelet/${braceletId}/activate/help`,
    );
  }

  handleHelpRequest(status: string, braceletId: number): Observable<void> {
    return this.http.put<void>(
      `${this.apiUrl}/bracelet/${braceletId}/activate/help/${status}`,
      null,
    );
  }
}
