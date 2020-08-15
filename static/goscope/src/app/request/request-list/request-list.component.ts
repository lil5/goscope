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
  page = 1;
  searchPage = 1;
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
      this.searchPage = 1;
      this.searchRequests();
    }
  }

  cancelSearchButtonPressed(): void {
    this.searchModeEnabled = false;
    document.getElementById('search-cancel-button').style.display = 'none';
    this.page = 1;
    this.searchPage = 1;
    this.getRequests();
  }


  searchRequests(): void {
    this.requestService.searchRequest(this.searchPage, this.searchQuery).subscribe(requests => {
      this.requests = requests.data;
      this.didGetNewContent = true;
    });
  }

  getRequests(): void {
    this.requestService.getRequests(this.page).subscribe(requests => {
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
      if (this.page > 1) {
        this.page--;
        this.getRequests();
      }
    } else {
      if (this.searchPage > 1) {
        this.searchPage--;
        this.searchRequests();
      }
    }
  }

  nextPage(): void {
    if (this.didGetNewContent) {
      if (!this.searchModeEnabled) {
        this.page++;
        this.getRequests();
      } else {
        this.searchPage++;
        this.searchRequests();
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
