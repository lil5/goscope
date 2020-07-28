import {Component, OnInit} from '@angular/core';
import {RequestService} from "../request.service";
import {Requests} from "../requests";
import {intervalToLevels} from "../../time-utils";

@Component({
  selector: 'request-request-list',
  templateUrl: './request-list.component.html',
  styleUrls: ['./request-list.component.scss']
})
export class RequestListComponent implements OnInit {
  requests: Requests[]
  now: number

  constructor(private requestService: RequestService) {
  }

  ngOnInit(): void {
    this.now = Math.round(new Date().getTime() / 1000)
    this.getRequests()
  }

  getRequests(): void {
    this.requestService.getRequests().subscribe(requests => this.requests = requests.data);
  }

  timeDiffToHuman(value: number): string {
    return intervalToLevels(value)
  }

  applyMethodColor(method: string): string {
    if (method === 'GET') {
      return `<span class="badge-secondary">${method}</span>`
    } else if (method === 'POST') {
      return `<span class="badge-info">${method}</span>`
    } else if (method === 'PUT') {
      return `<span class="badge-info">${method}</span>`
    } else if (method === 'PATCH') {
      return `<span class="badge-turq">${method}</span>`
    } else if (method === 'DELETE') {
      return `<span class="badge-danger">${method}</span>`
    }
    return `<span class="badge-secondary">${method}</span>`
  }

  applyStatusColor(status: number): string {
    if (status >= 200 && status < 300) {
      return `<span class="badge-success">${status}</span>`
    } else if (status >= 300 && status < 400) {
      return `<span class="badge-info">${status}</span>`
    } else if (status >= 400 && status < 500) {
      return `<span class="badge-warning">${status}</span>`
    }
    return `<span class="badge-danger">${status}</span>`
  }

}
