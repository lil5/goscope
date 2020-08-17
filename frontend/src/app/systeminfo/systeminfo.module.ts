import {BrowserModule} from '@angular/platform-browser';
import {DetailsComponent} from './details/details.component';
import {NgModule} from '@angular/core';


@NgModule({
  declarations: [DetailsComponent],
  imports: [
    BrowserModule,
  ],
  exports: [
    DetailsComponent
  ],
  providers: [

  ]
})
export class SysteminfoModule {
}

