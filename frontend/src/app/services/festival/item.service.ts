import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import {
  CreateItemPriceRequest,
  CreateItemRequest,
  FestivalPropCountResponse,
  Item,
  ItemsResponse,
} from '../../models/festival/festival.model';

@Injectable({
  providedIn: 'root',
})
export class ItemService {
  private apiUrl = 'http://localhost:4000';

  constructor(private http: HttpClient) {}

  getTicketTypes(festivalId: number): Observable<any> {
    return this.http
      .get<ItemsResponse>(
        `${this.apiUrl}/organizer/festival/${festivalId}/item/ticket-type`
      )
      .pipe(map((response) => response.items));
  }

  getTicketTypesCount(festivalId: number): Observable<number> {
    return this.http
      .get<FestivalPropCountResponse>(
        `${this.apiUrl}/organizer/festival/${festivalId}/item/ticket-type/count`
      )
      .pipe(map((response) => response.count));
  }

  getTicketType(festivalId: number, itemId: number): Observable<Item> {
    return this.http.get<Item>(
      `${this.apiUrl}/organizer/festival/${festivalId}/item/ticket-type/${itemId}`
    );
  }

  createItem(
    festivalId: number,
    request: CreateItemRequest
  ): Observable<number> {
    return this.http
      .post<{ itemId: number }>(
        `${this.apiUrl}/organizer/festival/${festivalId}/item`,
        request
      )
      .pipe(map((response) => response.itemId));
  }

  createItemPrice(
    festivalId: number,
    request: CreateItemPriceRequest
  ): Observable<number> {
    console.log('REQUEST', request);
    return this.http
      .post<{ priceListItemId: number }>(
        `${this.apiUrl}/organizer/festival/${festivalId}/item/price`,
        request
      )
      .pipe(map((response) => response.priceListItemId));
  }
}
