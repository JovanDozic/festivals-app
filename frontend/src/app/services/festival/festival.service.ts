import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import {
  CreateFestivalRequest,
  Festival,
  FestivalsResponse,
} from '../../models/festival/festival.model';

@Injectable({
  providedIn: 'root',
})
export class FestivalService {
  private apiUrl = 'http://localhost:4000';

  constructor(private http: HttpClient) {}

  getMyFestivals(): Observable<Festival[]> {
    return this.http
      .get<{ festivals: Festival[] }>(`${this.apiUrl}/organizer/festival`)
      .pipe(map((response) => response.festivals));
  }

  createFestival(festival: CreateFestivalRequest): Observable<any> {
    console.log('Creating a festival', festival);
    return this.http.post<any>(`${this.apiUrl}/festival`, festival);
  }

  deleteFestival(festivalId: number): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/festival/${festivalId}`);
  }

  publishFestival(festivalId: number): Observable<void> {
    return this.http.put<void>(
      `${this.apiUrl}/festival/${festivalId}/publish`,
      {}
    );
  }

  cancelFestival(festivalId: number): Observable<void> {
    return this.http.put<void>(
      `${this.apiUrl}/festival/${festivalId}/cancel`,
      {}
    );
  }

  completeFestival(festivalId: number): Observable<void> {
    return this.http.put<void>(
      `${this.apiUrl}/festival/${festivalId}/complete`,
      {}
    );
  }

  openFestivalStore(festivalId: number): Observable<void> {
    return this.http.put<void>(
      `${this.apiUrl}/festival/${festivalId}/store/open`,
      {}
    );
  }

  closeFestivalStore(festivalId: number): Observable<void> {
    return this.http.put<void>(
      `${this.apiUrl}/festival/${festivalId}/store/close`,
      {}
    );
  }
}