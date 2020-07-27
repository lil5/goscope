import {Component, OnInit} from '@angular/core';
import {RequestService} from "../request.service";
import {Requests} from "../requests";

@Component({
  selector: 'app-requests',
  templateUrl: './requests.component.html',
  styleUrls: ['./requests.component.scss']
})
export class RequestsComponent implements OnInit {
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
    this.requestService.getRequests().subscribe(requests => this.requests = requests.data);
    console.log(this.requests)
  }

}
