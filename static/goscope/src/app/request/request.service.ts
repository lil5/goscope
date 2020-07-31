import {Injectable} from '@angular/core';
import {DetailedRequestResponse, RequestsEndpointResponse} from './requests';
import {HttpClient, HttpParams} from '@angular/common/http';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class RequestService {
  private requestsUrl = 'http://localhost:7004/goscope/api/requests';
  private searchRequestsUrl = 'http://localhost:7004/goscope/api/search/requests';

  constructor(private http: HttpClient) {
  }

  getRequests(offset: number): Observable<RequestsEndpointResponse> {
    const options = {params: new HttpParams().set('offset', String(offset))};
    return this.http.get<RequestsEndpointResponse>(this.requestsUrl, options);
  }

  searchRequest(offset: number, query: string): Observable<RequestsEndpointResponse> {
    const options = {params: new HttpParams().set('offset', String(offset))};
    return this.http.post<RequestsEndpointResponse>(this.searchRequestsUrl, {query}, options);
  }


  getRequest(id: string): Observable<DetailedRequestResponse> {
    return this.http.get<DetailedRequestResponse>(`${this.requestsUrl}/${id}`);
  }
}
