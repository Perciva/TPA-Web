import { PromoComponent } from './../Product/promo/promo.component';
import { NotificationComponent } from './../General/notification/notification.component';
import { InsertBlogComponent } from './../Manage/manage-blog/insert-blog/insert-blog.component';
import { BlogDetailComponent } from './../Product/main-blog/blog-detail/blog-detail.component';
import { MainBlogComponent } from './../Product/main-blog/main-blog.component';
import { MainTrainComponent } from './../Product/main-train/main-train.component';
import { HotelSortPipe } from './../Pipes/hotel-sort.pipe';
import { HotelMainComponent } from './../Product/hotel/hotel.component';
import { RecHotelComponent } from '../Home/rec-hotel/rec-hotel.component';
import { PlaneHistoryComponent } from '../Home/QuickCard/plane/plane-history/plane-history.component';
import { HotelComponent } from '../Home/QuickCard/hotel/hotel.component';
import { TrainComponent } from '../Home/QuickCard/train/train.component';
import { EntertainmentComponent } from '../Home/QuickCard/entertainment/entertainment.component';
import { CarComponent } from '../Home/QuickCard/car/car.component';
import { PlaneComponent } from '../Home/QuickCard/plane/plane.component';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {MatRadioModule, MatDatepickerModule, MatNativeDateModule, MatFormFieldModule, MatInputModule, MatCheckboxModule, MatSidenavModule, MatExpansionModule, MatProgressBarModule, MatDividerModule, MatButtonModule, MatSelectModule } from '@angular/material'
import { ReactiveFormsModule, FormsModule } from '@angular/forms';
import { NgxContentLoadingModule } from 'ngx-content-loading';
import { StarRatingComponent } from '../General/star-rating/star-rating.component';
import { MapComponent } from '../Product/hotel/map/map.component';
import { HotelSearchComponent } from '../Product/hotel/hotel-search/hotel-search.component';
import { HotelDetailComponent } from '../Product/hotel/hotel-detail/hotel-detail.component';
import { Ng5SliderModule } from 'ng5-slider';
import { TrainDetailComponent } from '../Product/main-train/train-detail/train-detail.component';
import { TrainSearchComponent } from '../Product/main-train/train-search/train-search.component';
import { TrainSortPipe } from '../Pipes/train-sort.pipe';

const decs=[
  PlaneComponent,
  CarComponent,
  EntertainmentComponent,
  TrainComponent,
  HotelComponent,
  PlaneHistoryComponent,
  RecHotelComponent,
  HotelMainComponent,
  StarRatingComponent,
  MapComponent,
  HotelSearchComponent,
  HotelDetailComponent,
  HotelSortPipe,
  TrainDetailComponent,
  TrainSearchComponent,
  MainTrainComponent,
  TrainSortPipe,
  MainBlogComponent,
  BlogDetailComponent,
  PromoComponent
  
]
const imps = [
  MatRadioModule,
  MatDatepickerModule,
  MatNativeDateModule,
  ReactiveFormsModule,
  NgxContentLoadingModule,
  FormsModule,
  MatInputModule,
  MatFormFieldModule,
  MatCheckboxModule,
  MatExpansionModule,
  MatProgressBarModule,
  MatDividerModule,
  Ng5SliderModule,
  MatButtonModule,
  MatSelectModule
]
@NgModule({
  declarations: [
    
    decs
  ],
  imports: [
    CommonModule,
    
    imps
  ],
  exports:[
    decs,imps
  ],
  entryComponents:[
    InsertBlogComponent,
    NotificationComponent
  ]
  
})
export class QuickCardModule { }
