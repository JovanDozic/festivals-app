import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { ItemsResponse } from '../../models/festival/festival.model';

@Injectable({
  providedIn: 'root',
})
export class ItemService {
  private apiUrl = 'http://localhost:4000';

  constructor(private http: HttpClient) {}

  getTicketTypes(festivalId: number): Observable<any> {
    return this.http
      .get<ItemsResponse>(
        `${this.apiUrl}/organizer/festival/${festivalId}/item`
      )
      .pipe(map((response) => response.items));
  }
}
