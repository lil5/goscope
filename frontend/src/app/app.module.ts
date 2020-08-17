import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {NavbarComponent} from './navbar/navbar.component';
import {HttpClientModule} from '@angular/common/http';
import {RequestModule} from './request/request.module';
import {LogModule} from './log/log.module';
import {SysteminfoModule} from './systeminfo/systeminfo.module';
import { FooterComponent } from './footer/footer.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {HighlightService} from "./highlight.service";

@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    FooterComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    RequestModule,
    SysteminfoModule,
    BrowserAnimationsModule,
    LogModule,
  ],
  providers: [
    HighlightService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule {
}
