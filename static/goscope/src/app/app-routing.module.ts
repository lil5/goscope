import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {LogsComponent} from "./logs/logs.component";
import {RequestsComponent} from "./requests/requests.component";
import {RequestDetailsComponent} from "./request-details/request-details.component";

const routes: Routes = [
  {path: 'logs', component: LogsComponent},
  {
    path: 'requests', component: RequestsComponent, children: [
      {path: ':id', component: RequestDetailsComponent},
    ]
  },
  {path: '', redirectTo: '/requests', pathMatch: 'full'},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
