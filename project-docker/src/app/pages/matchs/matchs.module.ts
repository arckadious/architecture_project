import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { IonicModule } from '@ionic/angular';

import { MatchsPageRoutingModule } from './matchs-routing.module';

import { MatchsPage } from './matchs.page';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    MatchsPageRoutingModule
  ],
  declarations: [MatchsPage]
})
export class MatchsPageModule {}
