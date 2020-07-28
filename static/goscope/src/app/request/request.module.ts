import {BrowserModule} from "@angular/platform-browser";
import {HIGHLIGHT_OPTIONS, HighlightModule} from 'ngx-highlightjs';
import {HttpClientModule} from "@angular/common/http";
import {NgModule} from "@angular/core";
import {RequestDetailsComponent} from "./request-details/request-details.component";
import {RequestListComponent} from "./request-list/request-list.component";
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
    HighlightModule,
  ],
  providers: [
    {
      provide: HIGHLIGHT_OPTIONS,
      useValue: {
        coreLibraryLoader: () => import('highlight.js/lib/highlight'),
        lineNumbersLoader: () => import('highlightjs-line-numbers.js'), // Optional, only if you want the line numbers
        languages: {
          json: () => import('highlight.js/lib/languages/json'),
        }
      }
    }
  ],
})
export class RequestModule {
}
