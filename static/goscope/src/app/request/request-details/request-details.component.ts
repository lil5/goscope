import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import {Location} from '@angular/common';
import {RequestService} from '../request.service';
import {DetailedRequest, DetailedResponse} from '../requests';
import 'highlight.js/styles/dark.css';

@Component({
  selector: 'request-request-details',
  templateUrl: './request-details.component.html',
  styleUrls: ['./request-details.component.scss']
})
export class RequestDetailsComponent implements OnInit {
  requestDetails: DetailedRequest;
  responseDetails: DetailedResponse;

  constructor(
    private route: ActivatedRoute,
    private requestService: RequestService,
    private location: Location
  ) {

  }

  ngOnInit(): void {
    this.getRequest();
  }


  getRequest(): void {
    const id = this.route.snapshot.paramMap.get('id');
    this.requestService.getRequest(id)
      .subscribe(request => {
        this.requestDetails = request.data.request;
        this.responseDetails = request.data.response;
      });
  }

  goBack(): void {
    this.location.back();
  }

}
