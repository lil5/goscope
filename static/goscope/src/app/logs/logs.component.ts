import {Component, OnInit} from '@angular/core';
import {Logs} from '../logs';
import {LogService} from "../log.service";

@Component({
  selector: 'app-logs',
  templateUrl: './logs.component.html',
  styleUrls: ['./logs.component.scss']
})
export class LogsComponent implements OnInit {
  logs: Logs[];

  constructor(private logService: LogService) {
  }

  ngOnInit(): void {
    this.getLogs()
  }

  selectedLog: Logs;
  onSelect(log: Logs): void {
    this.selectedLog = log;
  }

  getLogs(): void {
    this.logs = this.logService.getLogs();
  }

}
