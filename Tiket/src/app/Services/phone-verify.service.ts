import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class PhoneVerifyService {

  constructor(private http: HttpClient) { }

  verifPhoneNumber(phoneNum:string){
    return this.http.get("http://apilayer.net/api/validate?access_key=0c2dd173a184e07c091f347e7325d39a&number="+ phoneNum)
  }
}
