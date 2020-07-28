import {Injectable} from '@angular/core';
import {DetailedRequestResponse, Requests, RequestsEndpointResponse} from './requests';
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
    const options = { params: new HttpParams().set('offset', String(offset)) }
    return this.http.get<RequestsEndpointResponse>(this.requestsUrl, options);
  }
  searchRequest(query: string): Observable<RequestsEndpointResponse> {
    return this.http.post<RequestsEndpointResponse>(this.searchRequestsUrl, {"query": query});
  }
  getRequest(id: string): Observable<DetailedRequestResponse> {
    return this.http.get<DetailedRequestResponse>(`${this.requestsUrl}/${id}`);
  }
}
