import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map, Observable, switchMap } from 'rxjs';
import {
  CreateItemPriceRequest,
  CreateItemRequest,
  FestivalPropCountResponse,
  Item,
  ItemsResponse,
  TransportAddonDTO,
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

  getPackageAddonsCount(
    festivalId: number,
    category: string
  ): Observable<number> {
    return this.http
      .get<FestivalPropCountResponse>(
        `${
          this.apiUrl
        }/organizer/festival/${festivalId}/item/package-addon/${category.toUpperCase()}/count`
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

  createPackageAddon(
    festivalId: number,
    request: CreateItemRequest,
    category: string
  ): Observable<number> {
    return this.http
      .post<{ itemId: number }>(
        `${this.apiUrl}/organizer/festival/${festivalId}/item`,
        request
      )
      .pipe(
        map((response) => response.itemId),
        switchMap((itemId) =>
          this.http.post<{ itemId: number }>(
            `${this.apiUrl}/organizer/festival/${festivalId}/item/package-addon`,
            { itemId, category }
          )
        ),
        map((response) => response.itemId)
      );
  }

  createItemPrice(
    festivalId: number,
    request: CreateItemPriceRequest
  ): Observable<number> {
    return this.http
      .post<{ priceListItemId: number }>(
        `${this.apiUrl}/organizer/festival/${festivalId}/item/price`,
        request
      )
      .pipe(map((response) => response.priceListItemId));
  }

  updateItem(festivalId: number, request: Item): Observable<any> {
    return this.http.put(
      `${this.apiUrl}/organizer/festival/${festivalId}/item/ticket-type/${request.id}`,
      request
    );
  }

  deleteItem(festivalId: number, itemId: number): Observable<any> {
    return this.http.delete(
      `${this.apiUrl}/organizer/festival/${festivalId}/item/ticket-type/${itemId}`
    );
  }

  addTransportConfig(festivalId: number, request: any): Observable<any> {
    return this.http.post(
      `${this.apiUrl}/organizer/festival/${festivalId}/item/package-addon/transport`,
      request
    );
  }

  addCampConfig(festivalId: number, request: any): Observable<any> {
    return this.http.post(
      `${this.apiUrl}/organizer/festival/${festivalId}/item/package-addon/camp`,
      request
    );
  }

  getTransportAddons(festivalId: number): Observable<TransportAddonDTO[]> {
    return this.http.get<TransportAddonDTO[]>(
      `${this.apiUrl}/organizer/festival/${festivalId}/item/package-addon/transport`
    );
  }
}
