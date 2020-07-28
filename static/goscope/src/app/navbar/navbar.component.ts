import { Component, OnInit } from '@angular/core';
import {RequestService} from "../request/request.service";

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.scss']
})
export class NavbarComponent implements OnInit {
  applicationName = 'GoScope';
  searchQuery: string;
  constructor(private requestService: RequestService,) { }

  ngOnInit(): void {
  }

  updateSearchQuery(textChanged: any): void {
    this.searchQuery = textChanged.target.value;
  }
  searchButtonPressed(): void {
    console.log(this.searchQuery)
    this.requestService.searchRequest(this.searchQuery).subscribe(response => console.log(response));
  }
}
