import { RecHotelComponent } from '../Home/rec-hotel/rec-hotel.component';
import { PlaneHistoryComponent } from '../Home/QuickCard/plane/plane-history/plane-history.component';
import { HotelComponent } from '../Home/QuickCard/hotel/hotel.component';
import { TrainComponent } from '../Home/QuickCard/train/train.component';
import { EntertainmentComponent } from '../Home/QuickCard/entertainment/entertainment.component';
import { CarComponent } from '../Home/QuickCard/car/car.component';
import { PlaneComponent } from '../Home/QuickCard/plane/plane.component';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {MatRadioModule, MatDatepickerModule, MatNativeDateModule, MatFormFieldModule, MatInputModule, MatCheckboxModule, MatSidenavModule } from '@angular/material'
import { ReactiveFormsModule, FormsModule } from '@angular/forms';
import { NgxContentLoadingModule } from 'ngx-content-loading';

@NgModule({
  declarations: [
    PlaneComponent,
    CarComponent,
    EntertainmentComponent,
    TrainComponent,
    HotelComponent,
    PlaneHistoryComponent,
    RecHotelComponent
  ],
  imports: [
    CommonModule,
    MatRadioModule,
    MatDatepickerModule,
    MatNativeDateModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatCheckboxModule,
    MatSidenavModule,
    FormsModule,
    NgxContentLoadingModule

  ],
  exports:[
    PlaneComponent,
    CarComponent,
    EntertainmentComponent,
    TrainComponent,
    HotelComponent,
    MatRadioModule,
    MatDatepickerModule,
    MatNativeDateModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatCheckboxModule,
    FormsModule,
    PlaneHistoryComponent,
    MatSidenavModule,
    RecHotelComponent,
    NgxContentLoadingModule
  ]
})
export class QuickCardModule { }
