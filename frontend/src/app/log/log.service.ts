import {Injectable} from '@angular/core';
import {DetailedLogsReponse, LogsEndpointResponse} from './logRecord';
import {HttpClient, HttpParams} from '@angular/common/http';
import {Observable} from 'rxjs';
import {environment} from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class LogService {
  private logsUrl = environment.apiLogsUrl;
  private searchLogsUrl = environment.apiSearchLogsUrl;

  constructor(private http: HttpClient) {
  }

  getLogs(page: number): Observable<LogsEndpointResponse> {
    const offset: number = (page - 1) * 50;
    const options = {params: new HttpParams().set('offset', offset.toString())};
    return this.http.get<LogsEndpointResponse>(this.logsUrl, options);
  }

  searchLog(page: number, query: string): Observable<LogsEndpointResponse> {
    const offset: number = (page - 1) * 50;
    const options = {params: new HttpParams().set('offset', offset.toString())};
    return this.http.post<LogsEndpointResponse>(this.searchLogsUrl, {query}, options);
  }

  getLog(id: string): Observable<DetailedLogsReponse> {
    return this.http.get<DetailedLogsReponse>(`${this.logsUrl}/${id}`);
  }
}