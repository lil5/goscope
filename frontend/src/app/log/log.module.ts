import {BrowserModule} from '@angular/platform-browser';
import {HttpClientModule} from '@angular/common/http';
import {LogDetailsComponent} from './log-details/log-details.component';
import {LogsComponent} from './logs/logs.component';
import {NgModule} from '@angular/core';
import {RouterModule} from '@angular/router';

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
    RouterModule,
    HttpClientModule,

  ],
  providers: [

  ],
})
export class LogModule {
}
