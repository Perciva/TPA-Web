import { EventDetailComponent } from './Product/event-detail/event-detail.component';
import { EventComponent } from './Product/event/event.component';
import { ChatroomComponent } from './Product/chatroom/chatroom.component';
import { PlaneSearchComponent } from './Product/main-plane/plane-search/plane-search.component';
import { MainCarComponent } from './Product/main-car/main-car.component';

import { HotelSearchComponent } from './Product/hotel/hotel-search/hotel-search.component';
import { HotelDetailComponent } from './Product/hotel/hotel-detail/hotel-detail.component';
import { HotelMainComponent } from './Product/hotel/hotel.component';
import { ManageTrainComponent } from './Manage/manage-train/manage-train.component';
import { ManageFlightComponent } from './Manage/manage-flight/manage-flight.component';
import { HomePageComponent } from './Home/home-page/home-page.component';
import { ManageHotelComponent } from './Manage/manage-hotel/manage-hotel.component';
import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ManageBlogComponent } from './Manage/manage-blog/manage-blog.component';
import { ManageEventComponent } from './Manage/manage-event/manage-event.component';
import { MapComponent } from './Product/hotel/map/map.component';
import { TrainDetailComponent } from './Product/main-train/train-detail/train-detail.component';
import { MainTrainComponent } from './Product/main-train/main-train.component';
import { TrainSearchComponent } from './Product/main-train/train-search/train-search.component';
import { BlogDetailComponent } from './Product/main-blog/blog-detail/blog-detail.component';
import { MainBlogComponent } from './Product/main-blog/main-blog.component';
import { PromoComponent } from './Product/promo/promo.component';
import { CarSearchComponent } from './Product/main-car/car-search/car-search.component';
import { MainPlaneComponent } from './Product/main-plane/main-plane.component';
import { ProfileComponent } from './nav/profile/profile.component';


const routes: Routes = [
  {path:'Manage', children:[{
    path:'Hotel',
    component: ManageHotelComponent
  },{
    path:'Train',
    component: ManageTrainComponent
  },{
    path:'Blog',
    component: ManageBlogComponent
  },{
    path:'Event',
    component: ManageEventComponent
  },{
    path:'Flight',
    component: ManageFlightComponent
  },]},
  {path:'', component: HomePageComponent},
  {path:'Promo', component: PromoComponent},
  {path:'Profile', component: ProfileComponent},
  {path:'Chat', component: ChatroomComponent},
  {path:'Hotel', children:[{
    path:'Detail',
    component: HotelDetailComponent
  },{
    path:'',
    component: HotelMainComponent
  },{
    path:'Search',
    component: HotelSearchComponent
  },{
    path:'Map',
    component: MapComponent
  },]},
  {path:'Train', children:[{
    path:'Detail',
    component: TrainDetailComponent
  },{
    path:'',
    component: MainTrainComponent
  },{
    path:'Search',
    component: TrainSearchComponent
  }]},
  {path:'Blog', children:[{
    path:'Detail',
    component: BlogDetailComponent
  },{
    path:'',
    component: MainBlogComponent
  },]},
  {path:'Car', children:[{
    path:'Search',
    component: CarSearchComponent
  },{
    path:'',
    component: MainCarComponent
  },]},
  {path:'Plane', children:[{
    path:'Search',
    component: PlaneSearchComponent
  },{
    path:'',
    component: MainPlaneComponent
  },]},
  {path:'Event', children:[{
    path:'Detail',
    component: EventDetailComponent
  },{
    path:'',
    component: EventComponent
  },]}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule { }
