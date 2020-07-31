import {Component, OnInit} from '@angular/core';
import {LogRecord} from '../logRecord';
import {LogService} from '../log.service';
import {intervalToLevels} from '../../time-utils';

@Component({
  selector: 'app-logs',
  templateUrl: './logs.component.html',
  styleUrls: ['./logs.component.scss']
})
export class LogsComponent implements OnInit {
  logs: LogRecord[];
  now: number;
  offset = 0;
  didGetNewContent = false;
  searchOffset = 0;
  searchModeEnabled = false;
  searchQuery: string;

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
    if (!this.searchModeEnabled) {
      if (this.offset >= 50) {
        this.offset -= 50;
        this.getLogs();
        if (!this.didGetNewContent) {
          this.offset += 50;
        }
      }
    } else {
      if (this.searchOffset >= 50) {
        this.searchOffset -= 50;
        this.getLogs();
        if (!this.didGetNewContent) {
          this.searchOffset += 50;
        }
      }
    }
  }

  updateSearchQuery(textChanged: any): void {
    this.searchQuery = textChanged.target.value.trim();
  }

  searchButtonPressed(): void {
    if (this.searchQuery !== '' && this.searchQuery) {
      this.searchModeEnabled = true;
      document.getElementById('search-cancel-button').style.display = 'flex';
      this.searchLogs();
    }
  }

  cancelSearchButtonPressed(): void {
    this.searchModeEnabled = false;
    document.getElementById('search-cancel-button').style.display = 'none';
    this.offset = 0;
    this.searchOffset = 0;
    this.getLogs();
  }

  searchLogs(): void {
    this.logService.searchLog(this.searchOffset, this.searchQuery).subscribe(requests => {
      this.logs = requests.data;
      this.didGetNewContent = true;
    });
  }

  nextPage(): void {
    if (!this.searchModeEnabled) {
      this.offset += 50;
      this.getLogs();
      if (!this.didGetNewContent) {
        this.offset -= 50;
      }
    } else {
      this.searchOffset += 50;
      this.searchLogs();
      if (!this.didGetNewContent) {
        this.searchOffset -= 50;
      }
    }
  }
}
