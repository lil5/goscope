import {Component, OnInit} from '@angular/core';
import {LogRecord} from '../logRecord';
import {LogService} from '../log.service';
import {intervalToLevels} from '../../time-utils';

@Component({
  selector: 'log-logs',
  templateUrl: './logs.component.html',
  styleUrls: ['./logs.component.scss']
})
export class LogsComponent implements OnInit {
  logs: LogRecord[];
  now: number;
  offset: number = 0;
  didGetNewContent: boolean = false;

  constructor(private logService: LogService) {
  }

  ngOnInit(): void {
    this.now = Math.round(new Date().getTime() / 1000);
    this.getLogs();
  }

  getLogs(): void {
    this.logService.getLogs(this.offset).subscribe(logs => this.logs = logs.data);
  }

  timeDiffToHuman(value: number): string {
    return intervalToLevels(value);
  }
  previousPage(): void {
    if (this.offset >= 50) {
      this.offset -= 50
      this.getLogs();
      if (!this.didGetNewContent) {
        this.offset += 50
      }
    }
  }

  nextPage(): void {
    this.offset += 50
    this.getLogs();
    if (!this.didGetNewContent) {
      this.offset -= 50
    }
  }
}
