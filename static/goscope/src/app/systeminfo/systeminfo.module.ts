import {BrowserModule} from '@angular/platform-browser';
import {DetailsComponent} from './details/details.component';
import {HIGHLIGHT_OPTIONS, HighlightModule} from 'ngx-highlightjs';
import {NgModule} from '@angular/core';


@NgModule({
  declarations: [DetailsComponent],
  imports: [
    BrowserModule,
    HighlightModule,
  ],
  exports: [
    DetailsComponent
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
  ]
})
export class SysteminfoModule {
}

