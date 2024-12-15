import { Injectable } from '@angular/core';
import { Log } from '../../models/common/log.model';
import { Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class LogsService {
  private apiUrl = 'http://localhost:4000';

  constructor(private http: HttpClient) {}

  getLogs(): Observable<Log[]> {
    return this.http.get<Log[]>(`${this.apiUrl}/log`);
  }
}
