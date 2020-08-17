import {AfterViewChecked, Component, OnInit} from '@angular/core';
import {SysteminfoService} from '../systeminfo.service';
import {SystemInfoDetailsResponse} from '../systeminfodetails';
import {Location} from '@angular/common';
import {HighlightService} from '../../highlight.service';

@Component({
  selector: 'app-details',
  templateUrl: './details.component.html',
  styleUrls: ['./details.component.scss']
})
export class DetailsComponent implements OnInit, AfterViewChecked {
  systemInformation: SystemInfoDetailsResponse;

  constructor(private systeminfoService: SysteminfoService, private location: Location,  private highlightService: HighlightService) {

  }
  ngAfterViewChecked(): void {
    this.highlightService.highlightAll();
  }

  ngOnInit(): void {
    this.getSystemInformation();

  }

  getSystemInformation(): void {
    this.systeminfoService.getSystemInfo().subscribe(info => this.systemInformation = info);
  }
  goBack(): void {
    this.location.back();
  }
}
