import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import {
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
}
