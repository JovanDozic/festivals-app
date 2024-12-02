import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map, Observable, switchMap } from 'rxjs';
import {
  AddCampConfigRequest,
  AddTransportConfigRequest,
  CampAddonDTO,
  CreateItemPriceRequest,
  CreateItemRequest,
  FestivalPropCountResponse,
  GeneralAddonDTO,
  Item,
  ItemCurrentPrice,
  ItemsResponse,
  TransportAddonDTO,
} from '../../models/festival/festival.model';

@Injectable({
  providedIn: 'root',
})
export class ItemService {
  private apiUrl = 'http://localhost:4000';

  constructor(private http: HttpClient) {}

  getTicketTypes(festivalId: number): Observable<ItemCurrentPrice[]> {
    return this.http
      .get<ItemsResponse>(
        `${this.apiUrl}/festival/${festivalId}/item/ticket-type`,
      )
      .pipe(map((response) => response.items));
  }

  getTicketTypesCount(festivalId: number): Observable<number> {
    return this.http
      .get<FestivalPropCountResponse>(
        `${this.apiUrl}/festival/${festivalId}/item/ticket-type/count`,
      )
      .pipe(map((response) => response.count));
  }

  getPackageAddonsCount(
    festivalId: number,
    category: string,
  ): Observable<number> {
    return this.http
      .get<FestivalPropCountResponse>(
        `${
          this.apiUrl
        }/festival/${festivalId}/item/package-addon/${category.toUpperCase()}/count`,
      )
      .pipe(map((response) => response.count));
  }

  getAllPackageAddonsCount(festivalId: number): Observable<number> {
    return this.http
      .get<FestivalPropCountResponse>(
        `${this.apiUrl}/festival/${festivalId}/item/package-addon/count`,
      )
      .pipe(map((response) => response.count));
  }

  getTicketType(festivalId: number, itemId: number): Observable<Item> {
    return this.http.get<Item>(
      `${this.apiUrl}/festival/${festivalId}/item/ticket-type/${itemId}`,
    );
  }

  createItem(
    festivalId: number,
    request: CreateItemRequest,
  ): Observable<number> {
    return this.http
      .post<{
        itemId: number;
      }>(`${this.apiUrl}/festival/${festivalId}/item`, request)
      .pipe(map((response) => response.itemId));
  }

  createPackageAddon(
    festivalId: number,
    request: CreateItemRequest,
    category: string,
  ): Observable<number> {
    return this.http
      .post<{
        itemId: number;
      }>(`${this.apiUrl}/festival/${festivalId}/item`, request)
      .pipe(
        map((response) => response.itemId),
        switchMap((itemId) =>
          this.http.post<{ itemId: number }>(
            `${this.apiUrl}/festival/${festivalId}/item/package-addon`,
            { itemId, category },
          ),
        ),
        map((response) => response.itemId),
      );
  }

  createItemPrice(
    festivalId: number,
    request: CreateItemPriceRequest,
  ): Observable<number> {
    return this.http
      .post<{
        priceListItemId: number;
      }>(`${this.apiUrl}/festival/${festivalId}/item/price`, request)
      .pipe(map((response) => response.priceListItemId));
  }

  updateItem(festivalId: number, request: Item): Observable<void> {
    return this.http.put<void>(
      `${this.apiUrl}/festival/${festivalId}/item/ticket-type/${request.id}`,
      request,
    );
  }

  deleteItem(festivalId: number, itemId: number): Observable<void> {
    return this.http.delete<void>(
      `${this.apiUrl}/festival/${festivalId}/item/ticket-type/${itemId}`,
    );
  }

  addTransportConfig(
    festivalId: number,
    request: AddTransportConfigRequest,
  ): Observable<void> {
    return this.http.post<void>(
      `${this.apiUrl}/festival/${festivalId}/item/package-addon/transport`,
      request,
    );
  }

  addCampConfig(
    festivalId: number,
    request: AddCampConfigRequest,
  ): Observable<void> {
    return this.http.post<void>(
      `${this.apiUrl}/festival/${festivalId}/item/package-addon/camp`,
      request,
    );
  }

  getTransportAddons(festivalId: number): Observable<TransportAddonDTO[]> {
    return this.http.get<TransportAddonDTO[]>(
      `${this.apiUrl}/festival/${festivalId}/item/package-addon/transport`,
    );
  }

  getGeneralAddons(festivalId: number): Observable<GeneralAddonDTO[]> {
    return this.http.get<GeneralAddonDTO[]>(
      `${this.apiUrl}/festival/${festivalId}/item/package-addon/general`,
    );
  }

  getCampAddons(festivalId: number): Observable<CampAddonDTO[]> {
    return this.http.get<CampAddonDTO[]>(
      `${this.apiUrl}/festival/${festivalId}/item/package-addon/camp`,
    );
  }
}
