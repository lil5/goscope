import {LogDetailsComponent} from "./log-details/log-details.component";
import {LogsComponent} from "./logs/logs.component";
import {BrowserModule} from "@angular/platform-browser";
import {HttpClientModule} from "@angular/common/http";
import {NgModule} from "@angular/core";

@NgModule({
  declarations: [
    LogDetailsComponent,
    LogsComponent,
  ],
  exports: [
    LogDetailsComponent,
    LogsComponent,
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
  ],
  providers: [],
})
export class LogModule { }
