import { NotificationComponent } from './General/notification/notification.component';
import { ChatComponent } from './Temp/chat/chat.component';

import { ManageModule } from './Modules/manage.module';
import { HomePageComponent } from './Home/home-page/home-page.component';
import { QuickCardComponent } from './Home/QuickCard/quick-card/quick-card.component';
import { ImageSliderComponent } from './Home/image-slider/image-slider.component';
import { RegisterComponent } from './nav/LoginRegister/register/register.component';
import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './nav/header/header.component';
import { FooterComponent } from './nav/footer/footer.component';
import { FooterTopComponent } from './nav/footer-top/footer-top.component';
import { LoginModule} from './Modules/login.module'


import { BrowserAnimationsModule } from '@angular/platform-browser/animations'
import { LoginComponent } from './nav/LoginRegister/login/login.component';
import { GraphQLModule } from './graphql.module';
import { HttpClientModule } from '@angular/common/http';
import { MAT_DIALOG_DEFAULT_OPTIONS } from '@angular/material';
import { QuickCardModule } from './Modules/quick-card.module';
import { StarRatingComponent } from './General/star-rating/star-rating.component';
import {icon, Marker} from 'leaflet';
import { ProfileComponent } from './nav/profile/profile.component';
import { EventComponent } from './Product/event/event.component';
const iconRetinaUrl = 'assets/leaflet/images/marker-icon-2x.png';
const iconUrl = 'assets/leaflet/images/marker-icon.png';
const shadowUrl = 'assets/leaflet/images/marker-shadow.png';

const iconDefault = icon({
  iconRetinaUrl,
  iconUrl,
  shadowUrl,
  iconSize: [25, 41],
  iconAnchor: [12, 41],
  popupAnchor: [1, -34],
  tooltipAnchor: [16, -28],
  shadowSize: [41, 41]
});
Marker.prototype.options.icon = iconDefault;
@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    FooterComponent,
    FooterTopComponent,
    LoginComponent,
    RegisterComponent,
    ImageSliderComponent,
    QuickCardComponent,
    HomePageComponent,
    ChatComponent,
    NotificationComponent,
    EventComponent

    
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    LoginModule,
    GraphQLModule,
    HttpClientModule,
    QuickCardModule,
    ManageModule,
  ],
  providers: [
    // {provide: MAT_DIALOG_DEFAULT_OPTIONS, useValue: {hasBackdrop: false}}+++++++++++++++++
    
  ],
  bootstrap: [AppComponent],
  entryComponents:[
    LoginComponent,
    RegisterComponent,
   
  ]
})
export class AppModule { }
