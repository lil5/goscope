import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {environment} from '../../environments/environment';
import {Observable} from 'rxjs';
import {ApplicationNameResponse} from './application-name-response';

@Injectable({
  providedIn: 'root'
})
export class AppdetailsService {
  private infoUrl = environment.apiApplicationNameUrl;

  constructor(private http: HttpClient) {
  }

  getApplicationName(): Observable<ApplicationNameResponse> {
    return this.http.get<ApplicationNameResponse>(this.infoUrl);
  }
}
