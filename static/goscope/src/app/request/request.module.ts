import {NgModule} from "@angular/core";
import {RequestListComponent} from "./request-list/request-list.component";
import {RequestDetailsComponent} from "./request-details/request-details.component";
import {BrowserModule} from "@angular/platform-browser";
import {HttpClientModule} from "@angular/common/http";
import {RouterModule} from '@angular/router';

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
    RouterModule,
    HttpClientModule,
  ],
  providers: [],
})
export class RequestModule {
}
