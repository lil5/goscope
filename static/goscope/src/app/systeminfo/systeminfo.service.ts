import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {SystemInfoDetailsResponse} from './systeminfodetails';

@Injectable({
  providedIn: 'root'
})
export class SysteminfoService {
  private infoUrl = '/goscope/api/info';

  constructor(private http: HttpClient) {
  }

  getSystemInfo(): Observable<SystemInfoDetailsResponse> {
    return this.http.get<SystemInfoDetailsResponse>(this.infoUrl);
  }

}
