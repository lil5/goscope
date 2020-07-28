import {Injectable} from '@angular/core';
import {DetailedLogsReponse, LogsEndpointResponse} from "./logRecord";
import {HttpClient} from '@angular/common/http';
import {Observable} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class LogService {
  private logsUrl = 'http://localhost:7004/goscope/api/logs';

  constructor(private http: HttpClient) {
  }

  getLogs(): Observable<LogsEndpointResponse> {
    return this.http.get<LogsEndpointResponse>(this.logsUrl);
  }

  getLog(id: string): Observable<DetailedLogsReponse> {
    return this.http.get<DetailedLogsReponse>(`${this.logsUrl}/${id}`)
  }
}
