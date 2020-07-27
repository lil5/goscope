import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NavbarComponent } from './navbar/navbar.component';
import { HttpClientModule }    from '@angular/common/http';
import { LogsComponent } from './logs/logs.component';
import { LogDetailsComponent } from './log-details/log-details.component';
import {RequestModule} from "./request/request.module";
// import {RequestListComponent} from "./request/request-list/request-list.component";
// import {RequestDetailsComponent} from "./request/request-details/request-details.component";

@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    LogsComponent,
    LogDetailsComponent,
    // RequestListComponent,
    // RequestDetailsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    RequestModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
