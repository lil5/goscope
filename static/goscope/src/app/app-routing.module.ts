import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {LogsComponent} from "./logs/logs.component";
import {RequestsComponent} from "./requests/requests.component";

const routes: Routes = [
  {path: 'logs', component: LogsComponent},
  {path: 'requests', component: RequestsComponent},
  {path: '', redirectTo: '/requests', pathMatch: 'full'},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
