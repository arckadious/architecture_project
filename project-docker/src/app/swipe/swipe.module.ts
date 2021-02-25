import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { IonicModule } from '@ionic/angular';

import { SwipePageRoutingModule } from './swipe-routing.module';
import { SwipeCardLibModule } from 'ng-swipe-card';

import { SwipePage } from './swipe.page';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    SwipePageRoutingModule,
    SwipeCardLibModule,

  ],
  declarations: [SwipePage]
})
export class SwipePageModule {}
