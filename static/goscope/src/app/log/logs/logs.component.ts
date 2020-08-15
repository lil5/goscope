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
  page = 1;
  didGetNewContent = false;
  searchPage = 0;
  searchModeEnabled = false;
  searchQuery: string;

  constructor(private logService: LogService) {
  }

  ngOnInit(): void {
    this.now = Math.round(new Date().getTime() / 1000);
    this.getLogs();
  }

  getLogs(): void {
    this.logService.getLogs(this.page).subscribe(logs => this.logs = logs.data);
  }

  timeDiffToHuman(value: number): string {
    return intervalToLevels(value);
  }

  previousPage(): void {
    if (!this.searchModeEnabled) {
      if (this.page > 1) {
        this.page--;
        this.getLogs();
      }
    } else {
      if (this.searchPage > 1) {
        this.searchPage--;
        this.getLogs();
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
      this.searchPage = 1;
      this.searchLogs();
    }
  }

  cancelSearchButtonPressed(): void {
    this.searchModeEnabled = false;
    document.getElementById('search-cancel-button').style.display = 'none';
    this.page = 1;
    this.searchPage = 1;
    this.getLogs();
  }

  searchLogs(): void {
    this.logService.searchLog(this.searchPage, this.searchQuery).subscribe(requests => {
      this.logs = requests.data;
      this.didGetNewContent = true;
    });
  }

  nextPage(): void {
    if (this.didGetNewContent) {
      if (!this.searchModeEnabled) {
        this.page++;
        this.getLogs();
      } else {
        this.searchPage++;
        this.searchLogs();
      }
    }
  }
}
