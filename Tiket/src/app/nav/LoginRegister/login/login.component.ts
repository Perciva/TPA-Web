import { ApolloService } from './../../../Services/apollo.service';
import { Component, OnInit } from '@angular/core';
import { AuthService } from "angularx-social-login";
import { FacebookLoginProvider, GoogleLoginProvider ,SocialUser} from "angularx-social-login";
import {Query} from '../../../Services/apollo.service'
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  private user: SocialUser;
  private loggedIn: boolean;
  private userInput:string;

  constructor( 
    private auth:AuthService,
    private apollo: ApolloService
    ) { }

  ngOnInit() {
    this.auth.authState.subscribe((user) => {
      this.user = user;
      this.loggedIn = (user != null);
    });
  }
  signInGoogle():void{
    this.auth.signIn(GoogleLoginProvider.PROVIDER_ID)
  }
  signInFB():void{
    this.auth.signIn(FacebookLoginProvider.PROVIDER_ID)
  }

  signIn():void{
    this.userInput=(<HTMLInputElement>document.getElementById('sign')).value;
    this.apollo.searchUserByEmailPhone(this.userInput).subscribe(
      Query=>{
        this.userInput = Query.data.searchuseremailphone
      }
    )
  }
}
