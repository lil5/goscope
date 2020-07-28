import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NavbarComponent } from './navbar/navbar.component';
import { HttpClientModule }    from '@angular/common/http';
import {RequestModule} from "./request/request.module";
import {LogModule} from "./log/log.module";


@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    RequestModule,
    LogModule,

  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
