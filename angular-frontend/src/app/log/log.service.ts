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


}
