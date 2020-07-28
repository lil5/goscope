import {Component, Input, OnInit} from '@angular/core';
import {Logs} from "../../log/logs";
import {Requests} from "../requests";

@Component({
  selector: 'request-request-details',
  templateUrl: './request-details.component.html',
  styleUrls: ['./request-details.component.scss']
})
export class RequestDetailsComponent implements OnInit {
  @Input() request: Requests;
  constructor() { }

  ngOnInit(): void {
  }

}
