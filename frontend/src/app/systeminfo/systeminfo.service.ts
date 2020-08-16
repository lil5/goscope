import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {SystemInfoDetailsResponse} from './systeminfodetails';
import {environment} from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class SysteminfoService {
  private infoUrl = environment.apiSysInfoUrl;

  constructor(private http: HttpClient) {
  }

  getSystemInfo(): Observable<SystemInfoDetailsResponse> {
    return this.http.get<SystemInfoDetailsResponse>(this.infoUrl);
  }

}
