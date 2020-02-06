import { MatDialogRef } from '@angular/material';
import { ApolloService } from './../../../Services/apollo.service';
import { Component, OnInit } from '@angular/core';
import { AuthService } from "angularx-social-login";
import { Subscription } from 'rxjs';
import { FacebookLoginProvider, GoogleLoginProvider ,SocialUser} from "angularx-social-login";
import { HttpClient } from '@angular/common/http';
import { Query } from 'src/app/Models/query';
// import {Query} from '../../../Services/apollo.service'
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  private user: SocialUser;
  private loggedIn: boolean;
  private userInput:string;
  private nvalid:boolean = false;
  private message:string;
  private len:number 
  private data:any


  constructor( 
    private auth:AuthService,
    private apollo: ApolloService,
    private ref:MatDialogRef<LoginComponent>
    ) { }

  ngOnInit() {
    this.auth.authState.subscribe((user) => {
      this.user = user;
      this.loggedIn = (user != null);
    });
  }
  close():void{
    this.ref.close({
      fromModal : true,
      data : this.data
    })
  }
  err():void{
      if(this.message.length!=0){
        document.getElementById('sign').classList.add('err')
      }
      else{
        document.getElementById('sign').classList.remove('err')
      }
  }
  checkEmail():boolean{
    this.len = this.userInput.length
    let flag:number = 1000000
    for(let i:number = 0;i<this.len;i++){
      if(this.userInput.charAt(i)=='@'){
        flag=i
        console.log(flag)
      }
      if(flag < i && this.userInput.charAt(i) == "."){
        return (i < this.len+1)
      }
    }
    return false
  }
  validate():void{
    const regex = new RegExp('^[0-9]+$')

    if(regex.test(this.userInput)){
      if(this.userInput.length < 9){
        this.message = "nomor ponsel tidak sesuai"
        this.err()
      }
    }
    else{
      if(!this.checkEmail()){
        this.message = "format email harus seperti email@mail.com"
        this.err()
      }
    }
  }
  signInGoogle():void{
    this.auth.signIn(GoogleLoginProvider.PROVIDER_ID)
  }
  signInFB():void{
    this.auth.signIn(FacebookLoginProvider.PROVIDER_ID)

    console.log(this.user.facebook)
  }

  userInput$: Subscription

  signIn():void{
    if(this.message.length ==0 && this.userInput.length != 0){
      this.userInput=(<HTMLInputElement>document.getElementById('sign')).value;

      this.userInput$ = this.apollo.searchUserByEmailPhone(this.userInput).subscribe(async(Query) => {
          await this.tes(Query)
      })
      console.log(this.data)
      if(this.data == null){
        this.data = this.userInput
      }
      this.close()

    } 
  }
  tes(Query: Query) {
    
      this.data = Query.data.userbyemailphone
      
      console.log(this.data)  
  }
}
