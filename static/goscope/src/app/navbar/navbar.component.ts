import {Component, OnInit} from '@angular/core';
import {RequestService} from '../request/request.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.scss']
})
export class NavbarComponent implements OnInit {
  applicationName = 'GoScope';

  constructor(private requestService: RequestService) {
  }

  ngOnInit(): void {
  }


}
