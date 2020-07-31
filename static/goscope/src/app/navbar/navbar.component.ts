import {Component, OnInit} from '@angular/core';
import {AppdetailsService} from './appdetails.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.scss']
})
export class NavbarComponent implements OnInit {
  applicationName = '';

  constructor(private appdetailsService: AppdetailsService) {
  }

  ngOnInit(): void {
    this.getAppName();
  }

  getAppName(): void {
    this.appdetailsService.getApplicationName().subscribe(info => this.applicationName = info.applicationName);
  }

}
