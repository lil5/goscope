import {Component, OnInit} from '@angular/core';
import {SysteminfoService} from '../systeminfo.service';
import {SystemInfoDetailsResponse} from "../systeminfodetails";
import {Location} from "@angular/common";
import 'highlight.js/styles/dark.css';

@Component({
  selector: 'app-details',
  templateUrl: './details.component.html',
  styleUrls: ['./details.component.scss']
})
export class DetailsComponent implements OnInit {
  systemInformation: SystemInfoDetailsResponse

  constructor(private systeminfoService: SysteminfoService, private location: Location) {

  }

  ngOnInit(): void {
    this.getSystemInformation()
  }

  getSystemInformation(): void {
    this.systeminfoService.getSystemInfo().subscribe(info => this.systemInformation = info)
  }
  goBack(): void {
    this.location.back();
  }
}
