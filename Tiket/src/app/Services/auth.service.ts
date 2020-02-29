import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor() { }
  logout():void{
    sessionStorage.setItem("logged","false");
    sessionStorage.removeItem("token")
  }
}
