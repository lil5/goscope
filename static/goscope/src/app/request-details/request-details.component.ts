import {Component, Input, OnInit} from '@angular/core';
import {Logs} from "../logs";
import {Requests} from "../requests";

@Component({
  selector: 'app-request-details',
  templateUrl: './request-details.component.html',
  styleUrls: ['./request-details.component.scss']
})
export class RequestDetailsComponent implements OnInit {
  @Input() request: Requests;
  constructor() { }

  ngOnInit(): void {
  }

}
