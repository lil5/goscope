import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {LogsComponent} from "./log/logs/logs.component";
import {RequestListComponent} from "./request/request-list/request-list.component";
import {RequestDetailsComponent} from "./request/request-details/request-details.component";
import {RequestModule} from "./request/request.module";
import {LogModule} from "./log/log.module";

const routes: Routes = [
  {path: 'logs', component: LogsComponent},
  {
    path: 'requests', component: RequestListComponent, children: [
      {path: ':id', component: RequestDetailsComponent},
    ]
  },
  {path: '', redirectTo: '/requests', pathMatch: 'full'},
];

@NgModule({
  imports: [RouterModule.forRoot(routes), RequestModule, LogModule],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
