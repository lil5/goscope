import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {LogsComponent} from "./logs/logs.component";
import {RequestListComponent} from "./request/request-list/request-list.component";
import {RequestDetailsComponent} from "./request-details/request-details.component";
import {RequestModule} from "./request/request.module";

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
  imports: [RouterModule.forRoot(routes), RequestModule],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
