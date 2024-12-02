import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import {
  CreateFestivalRequest,
  Employee,
  EmployeesResponse,
  Festival,
  FestivalPropCountResponse,
  UpdateFestivalRequest,
} from '../../models/festival/festival.model';

@Injectable({
  providedIn: 'root',
})
export class FestivalService {
  private apiUrl = 'http://localhost:4000';

  constructor(private http: HttpClient) {}

  getMyFestivals(): Observable<Festival[]> {
    return this.http
      .get<{ festivals: Festival[] }>(`${this.apiUrl}/festival/organizer`)
      .pipe(map((response) => response.festivals));
  }

  getEmployeeCount(festivalId: number): Observable<number> {
    return this.http
      .get<FestivalPropCountResponse>(
        `${this.apiUrl}/festival/${festivalId}/employee/count`,
      )
      .pipe(map((response) => response.count));
  }

  getEmployees(festivalId: number): Observable<Employee[]> {
    return this.http
      .get<EmployeesResponse>(`${this.apiUrl}/festival/${festivalId}/employee`)
      .pipe(map((response) => response.employees));
  }

  getEmployeesNotOnFestival(festivalId: number): Observable<Employee[]> {
    return this.http
      .get<EmployeesResponse>(
        `${this.apiUrl}/festival/${festivalId}/employee/available`,
      )
      .pipe(map((response) => response.employees));
  }

  employEmployee(festivalId: number, employeeId: number): Observable<void> {
    return this.http.put<void>(
      `${this.apiUrl}/festival/${festivalId}/employee/${employeeId}/employ`,
      {},
    );
  }

  fireEmployee(festivalId: number, employeeId: number): Observable<void> {
    return this.http.delete<void>(
      `${this.apiUrl}/festival/${festivalId}/employee/${employeeId}/fire`,
      {},
    );
  }

  createFestival(festival: CreateFestivalRequest): Observable<{ id: number }> {
    return this.http.post<{ id: number }>(`${this.apiUrl}/festival`, festival);
  }

  deleteFestival(festivalId: number): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/festival/${festivalId}`);
  }

  publishFestival(festivalId: number): Observable<void> {
    return this.http.put<void>(
      `${this.apiUrl}/festival/${festivalId}/publish`,
      {},
    );
  }

  cancelFestival(festivalId: number): Observable<void> {
    return this.http.put<void>(
      `${this.apiUrl}/festival/${festivalId}/cancel`,
      {},
    );
  }

  completeFestival(festivalId: number): Observable<void> {
    return this.http.put<void>(
      `${this.apiUrl}/festival/${festivalId}/complete`,
      {},
    );
  }

  openFestivalStore(festivalId: number): Observable<void> {
    return this.http.put<void>(
      `${this.apiUrl}/festival/${festivalId}/store/open`,
      {},
    );
  }

  closeFestivalStore(festivalId: number): Observable<void> {
    return this.http.put<void>(
      `${this.apiUrl}/festival/${festivalId}/store/close`,
      {},
    );
  }

  updateFestival(festival: UpdateFestivalRequest): Observable<void> {
    return this.http.put<void>(
      `${this.apiUrl}/festival/${festival.id}`,
      festival,
    );
  }

  getFestival(festivalId: number): Observable<Festival> {
    return this.http
      .get<{ festival: Festival }>(`${this.apiUrl}/festival/${festivalId}`)
      .pipe(map((response) => response.festival));
  }
}
