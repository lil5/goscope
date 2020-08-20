import {BrowserModule} from '@angular/platform-browser';
import {HttpClientModule} from '@angular/common/http';
import {NgModule} from '@angular/core';
import {RequestDetailsComponent} from './request-details/request-details.component';
import {RequestListComponent} from './request-list/request-list.component';
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
  providers: [

  ],
})
export class RequestModule {
}
