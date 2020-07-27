import {Component, OnInit} from '@angular/core';
import {RequestService} from "../request.service";
import {Requests} from "../requests";

@Component({
  selector: 'request-request-list',
  templateUrl: './request-list.component.html',
  styleUrls: ['./request-list.component.scss']
})
export class RequestListComponent implements OnInit {
  requests: Requests[]

  constructor(private requestService: RequestService) {
  }

  ngOnInit(): void {
    this.getRequests()
  }

  selectedRequest: Requests;

  onSelect(request: Requests): void {
    this.selectedRequest = request;
  }

  getRequests(): void {
    this.requests = this.requestService.getRequests();
  }

}
