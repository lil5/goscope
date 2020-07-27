import {Injectable} from '@angular/core';
import {Requests} from "./requests";

@Injectable({
  providedIn: 'root'
})
export class RequestService {

  constructor() {
  }

  getRequests(): Requests[] {
    return [
      {
        uid: "test",
        time: 123,
        method: "POST",
        path: "/",
        responseStatus: 200,
      },
      {
        uid: "test",
        time: 123,
        method: "POST",
        path: "/",
        responseStatus: 200,
      },
    ];
  }
}
