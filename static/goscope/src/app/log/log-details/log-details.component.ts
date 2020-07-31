import {Component, OnInit} from '@angular/core';
import {DetailedLogsReponse} from '../logRecord';
import {ActivatedRoute} from '@angular/router';
import {Location} from '@angular/common';
import {LogService} from '../log.service';
import 'highlight.js/styles/dark.css';

@Component({
  selector: 'app-log-details',
  templateUrl: './log-details.component.html',
  styleUrls: ['./log-details.component.scss']
})

export class LogDetailsComponent implements OnInit {
  logDetails: DetailedLogsReponse;

  constructor(
    private route: ActivatedRoute,
    private logService: LogService,
    private location: Location
  ) {
  }

  ngOnInit(): void {
    this.getLog();
  }

  getLog(): void {
    const id = this.route.snapshot.paramMap.get('id');
    this.logService.getLog(id)
      .subscribe(log => {
        this.logDetails = log;
      });
  }

  goBack(): void {
    this.location.back();
  }

}
