import {NgModule} from "@angular/core";
import {AppComponent} from "../app.component";
import {NavbarComponent} from "../navbar/navbar.component";
import {LogsComponent} from "../log/logs/logs.component";
import {LogDetailsComponent} from "../log/log-details/log-details.component";
import {RequestListComponent} from "./request-list/request-list.component";
import {RequestDetailsComponent} from "./request-details/request-details.component";
import {BrowserModule} from "@angular/platform-browser";
import {AppRoutingModule} from "../app-routing.module";
import {HttpClientModule} from "@angular/common/http";

@NgModule({
  declarations: [
    RequestListComponent,
    RequestDetailsComponent,
  ],
  exports: [
    RequestListComponent,
    RequestDetailsComponent,
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
  ],
  providers: [],
})
export class RequestModule { }
