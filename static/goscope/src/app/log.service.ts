import { Injectable } from '@angular/core';
import {Logs} from "./logs";

@Injectable({
  providedIn: 'root'
})
export class LogService {

  constructor() { }
  getLogs(): Logs[] {
    return [
      {
        uid: "test",
        time: 123,
        error: "testing error"
      },
      {
        uid: "test2",
        time: 1234,
        error: "testing error2"
      }];
  }
}
