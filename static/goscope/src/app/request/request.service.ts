import {Injectable} from '@angular/core';
import {DetailedRequestResponse, RequestsEndpointResponse} from './requests';
import {HttpClient, HttpParams} from '@angular/common/http';
import {Observable} from 'rxjs';
import {environment} from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class RequestService {
  private requestsUrl = environment.apiRequestsUrl;
  private searchRequestsUrl = environment.apiSearchRequestsUrl;

  constructor(private http: HttpClient) {
  }

  getRequests(page: number): Observable<RequestsEndpointResponse> {
    const offset: number = (page - 1) * 50;
    const options = {params: new HttpParams().set('offset', offset.toString())};
    return this.http.get<RequestsEndpointResponse>(this.requestsUrl, options);
  }

  searchRequest(page: number, query: string): Observable<RequestsEndpointResponse> {
    const offset: number = (page - 1) * 50;
    const options = {params: new HttpParams().set('offset', offset.toString())};
    return this.http.post<RequestsEndpointResponse>(this.searchRequestsUrl, {query}, options);
  }


  getRequest(id: string): Observable<DetailedRequestResponse> {
    return this.http.get<DetailedRequestResponse>(`${this.requestsUrl}/${id}`);
  }
}
