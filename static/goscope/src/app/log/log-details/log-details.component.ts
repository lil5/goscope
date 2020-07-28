import { Component, OnInit, Input } from '@angular/core';
import {Logs} from "../../logs";

@Component({
  selector: 'log-log-details',
  templateUrl: './log-details.component.html',
  styleUrls: ['./log-details.component.scss']
})

export class LogDetailsComponent implements OnInit {
  @Input() log: Logs;
  constructor() { }

  ngOnInit(): void {
  }

}
