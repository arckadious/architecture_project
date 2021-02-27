import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { MatchsPage } from './matchs.page';

const routes: Routes = [
  {
    path: '',
    component: MatchsPage
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class MatchsPageRoutingModule {}
