import {BrowserModule} from '@angular/platform-browser';
import {HIGHLIGHT_OPTIONS, HighlightModule} from 'ngx-highlightjs';
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
export class LogModule { }
