import {Component, OnInit, ViewEncapsulation} from '@angular/core';
import {RequestService} from '../request.service';
import {Requests} from '../requests';
import {intervalToLevels} from '../../time-utils';

@Component({
  selector: 'app-request-list',
  templateUrl: './request-list.component.html',
  styleUrls: ['./request-list.component.scss'],
  encapsulation: ViewEncapsulation.None,
})
export class RequestListComponent implements OnInit {
  requests: Requests[];
  now: number;
  offset = 0;
  searchOffset = 0;
  searchModeEnabled = false;
  didGetNewContent = false;
  searchQuery: string;

  constructor(private requestService: RequestService) {
  }

  ngOnInit(): void {
    this.now = Math.round(new Date().getTime() / 1000);
    this.getRequests();
  }

  updateSearchQuery(textChanged: any): void {
    this.searchQuery = textChanged.target.value.trim();
  }

  searchButtonPressed(): void {
    if (this.searchQuery !== '' && this.searchQuery){
      this.searchModeEnabled = true;
      document.getElementById('search-cancel-button').style.display = 'flex';
      this.searchRequests();
    }
  }

  cancelSearchButtonPressed(): void {
    this.searchModeEnabled = false;
    document.getElementById('search-cancel-button').style.display = 'none';
    this.offset = 0;
    this.searchOffset = 0;
    this.getRequests();
  }


  searchRequests(): void {
    this.requestService.searchRequest(this.searchOffset, this.searchQuery).subscribe(requests => {
      this.requests = requests.data;
      this.didGetNewContent = true;
    });
  }

  getRequests(): void {
    this.requestService.getRequests(this.offset).subscribe(requests => {
      if (requests.data !== null && requests.data.length > 0) {
        this.requests = requests.data;
        this.didGetNewContent = true;
      } else {
        this.didGetNewContent = false;
      }
    });
  }

  previousPage(): void {
    if (!this.searchModeEnabled) {
      if (this.offset >= 50) {
        this.offset -= 50;
        this.getRequests();
        if (!this.didGetNewContent) {
          this.offset += 50;
        }
      }
    } else {
      if (this.searchOffset >= 50) {
        this.searchOffset -= 50;
        this.searchRequests();
        if (!this.didGetNewContent) {
          this.searchOffset += 50;
        }
      }
    }
  }

  nextPage(): void {
    if (!this.searchModeEnabled) {
      this.offset += 50;
      this.getRequests();
      if (!this.didGetNewContent) {
        this.offset -= 50;
      }
    } else {
      this.searchOffset += 50;
      this.searchRequests();
      if (!this.didGetNewContent) {
        this.searchOffset -= 50;
      }
    }
  }

  timeDiffToHuman(value: number): string {
    return intervalToLevels(value);
  }

  applyMethodColor(method: string): string {
    if (method === 'GET') {
      return `<span class="badge-secondary">${method}</span>`;
    } else if (method === 'POST') {
      return `<span class="badge-info">${method}</span>`;
    } else if (method === 'PUT') {
      return `<span class="badge-info">${method}</span>`;
    } else if (method === 'PATCH') {
      return `<span class="badge-turq">${method}</span>`;
    } else if (method === 'DELETE') {
      return `<span class="badge-danger">${method}</span>`;
    }
    return `<span class="badge-secondary">${method}</span>`;
  }

  applyStatusColor(status: number): string {
    if (status >= 200 && status < 300) {
      return `<span class="badge-success">${status}</span>`;
    } else if (status >= 300 && status < 400) {
      return `<span class="badge-info">${status}</span>`;
    } else if (status >= 400 && status < 500) {
      return `<span class="badge-warning">${status}</span>`;
    }
    return `<span class="badge-danger">${status}</span>`;
  }

}
