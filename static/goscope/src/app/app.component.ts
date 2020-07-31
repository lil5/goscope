import {Component} from '@angular/core';
import {animate, query, style, transition, trigger} from '@angular/animations';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
  animations: [
    trigger('myAnimation', [
      transition('* => *', [
        query(
          ':enter',
          [style({opacity: 0})],
          {optional: true}
        ),
        query(
          ':leave',
          [style({opacity: 1}), animate('0.15s', style({opacity: 0}))],
          {optional: true}
        ),
        query(
          ':enter',
          [style({opacity: 0}), animate('0.15s', style({opacity: 1}))],
          {optional: true}
        )
      ])
    ])
  ]
})
export class AppComponent {
  title = 'goscope';
  appName = 'App Name';
}
