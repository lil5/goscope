import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {LogsComponent} from './log/logs/logs.component';
import {RequestListComponent} from './request/request-list/request-list.component';
import {RequestModule} from './request/request.module';
import {LogModule} from './log/log.module';
import {RequestDetailsComponent} from './request/request-details/request-details.component';
import {LogDetailsComponent} from './log/log-details/log-details.component';
import {DetailsComponent} from "./systeminfo/details/details.component";

const routes: Routes = [
  {path: 'logs', component: LogsComponent},
  {path: 'logs/:id', component: LogDetailsComponent},
  {path: 'requests', component: RequestListComponent},
  {path: 'requests/:id', component: RequestDetailsComponent},
  {path: '', redirectTo: '/requests', pathMatch: 'full'},
  {path: 'info', component: DetailsComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes), RequestModule, LogModule],
  exports: [RouterModule, RequestModule, LogModule]
})
export class AppRoutingModule {
}
