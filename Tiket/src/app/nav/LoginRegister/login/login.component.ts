import { User } from './../../../Models/user';
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
  private log: boolean=false;
  private userInput:string;
  private nvalid:boolean = false;
  private message:string;
  private len:number 
  private data:any
  private img:boolean = false;
  regis$: Subscription;
  constructor( 
    private auth:AuthService,
    private apollo: ApolloService,
    private ref:MatDialogRef<LoginComponent>
    ) { }

  ngOnInit() {
    this.auth.authState.subscribe((user) => {
      this.user = user;

    });
    this.userInput=""
  }
  close():void{
    this.ref.close({
      fromModal : true,
      data : this.data,
      logged : this.nvalid
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
        // console.log(flag)
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
    this.loginFromSocMed()
  }
  signInFB():void{
    this.auth.signIn(FacebookLoginProvider.PROVIDER_ID)
    this.loginFromSocMed()
    // console.log(this.user.facebook)
  }
  loginFromSocMed(){
    console.log(this.user)
    if(this.user){
      var email:string = this.user.email
      this.userInput$=this.apollo.searchUserByEmailPhone(email).subscribe(async query => {
        await this.tes(query)
    })
    }
  }
  userInput$: Subscription
  
  signIn():void{
    if(this.message.length ==0 && this.userInput.length != 0&&this.log==false){
      this.log = true;
      this.userInput=(<HTMLInputElement>document.getElementById('sign')).value;
     
      this.userInput$ = this.apollo.searchUserByEmailPhone(this.userInput).subscribe(async query => {
          await this.tes(query)
      })
      
    } 
  }

  login():void{
    var pass:string = (<HTMLInputElement>document.getElementById('pass')).value;
    console.log("ooo: " + pass)
    if(pass == this.data[0].password){
      console.log('login')
      this.setSess(this.data[0].id)
    }
    else{
      this.message="incorrect email or password"
    }
  }

  togle():void{
    if(!this.img){
      (<HTMLImageElement>document.getElementById('img')).src = "../../../../assets/Nav/login/invisible.png";
      (<HTMLInputElement>document.getElementById('pass')).type = "text"
    }
    else{
      (<HTMLImageElement>document.getElementById('img')).src = "../../../../assets/Nav/login/visible.png";
      (<HTMLInputElement>document.getElementById('pass')).type = "password"
    }
    this.img = !this.img
  }
  ngOnDestroy(): void {
    this.userInput$.unsubscribe()
    //Called once, before the instance is destroyed.
    //Add 'implements OnDestroy' to the class.
    
  }
  tes(a: any) {
    console.log(a.data.userbyemailphone.length);
      this.data = a.data.userbyemailphone

      if(this.data.length === 0){
        if(this.user){
          var newUser:User = new User(this.user.firstName, this.user.lastName, this.user.email, "" ,"")
          // console.log(newUser)
          this.regis$ = this.apollo.createUser(newUser).subscribe(async (d:any) => {
            // console.log("registered"+ d)
            this.setSess(d.data.createuser.id);   
          })
          
          return;
        }
        else{
          this.data = this.userInput
          this.close()
        }
      }
      else{
        if(this.user){
          this.setSess(this.data[0].id);
          return;
        }
        document.getElementById('passrow').classList.remove('hid')
      }
  }
  setSess(val:number):void{
          this.nvalid = true;

    sessionStorage.setItem("logged","true")
    sessionStorage.setItem("token", val.toString());
    this.close();
  }
}
