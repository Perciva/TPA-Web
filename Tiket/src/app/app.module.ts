import { RegisterComponent } from './nav/LoginRegister/register/register.component';
import { MapComponent } from './app/temp/map/map.component';
import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './nav/header/header.component';
import { FooterComponent } from './nav/footer/footer.component';
import { FooterTopComponent } from './nav/footer-top/footer-top.component';
import { LoginModule} from '../app/Modules/login-module/login.module'


import { BrowserAnimationsModule } from '@angular/platform-browser/animations'
import { LoginComponent } from './nav/LoginRegister/login/login.component';
import { GraphQLModule } from './graphql.module';
import { HttpClientModule } from '@angular/common/http';
import { MAT_DIALOG_DEFAULT_OPTIONS } from '@angular/material';
@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    FooterComponent,
    FooterTopComponent,
    LoginComponent,
    MapComponent,
    RegisterComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    LoginModule,
    GraphQLModule,
    HttpClientModule
  ],
  providers: [
    {provide: MAT_DIALOG_DEFAULT_OPTIONS, useValue: {hasBackdrop: false}}
  ],
  bootstrap: [AppComponent],
  entryComponents:[
    LoginComponent,
    RegisterComponent,
  ]
})
export class AppModule { }
