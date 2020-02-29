import { HomePageComponent } from './Home/home-page/home-page.component';
import { ManageHotelComponent } from './Manage/manage-hotel/manage-hotel.component';
import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';


const routes: Routes = [
  {path:'Manage', children:[{
    path:'Hotel',
    component: ManageHotelComponent
  }]},
  {path:'', component: HomePageComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
