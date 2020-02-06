import { Query } from './../../../Models/query';
import { User } from './../../../Models/user';
import { Subscription } from 'rxjs';
import { Component, OnInit, ComponentFactoryResolver, Inject } from '@angular/core';
import { ApolloService } from './../../../Services/apollo.service';
import { inject } from '@angular/core/testing';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material';
@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent implements OnInit {
  private img:boolean = false;//if false -> invisible
  private message:string = "tes";
  

  constructor(
    private apollo: ApolloService,
    @Inject(MAT_DIALOG_DATA) public data: any,
    private ref:MatDialogRef<RegisterComponent>
  ) { }

  ngOnInit() {
    console.log(this.data)
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
  err(id:string):void{
    if(this.message.length!=0){
      document.getElementById(id).classList.add('err')
    }
    else{
      document.getElementById(id).classList.remove('err')
    }
}
  validatePassword(password:string):boolean{
    if(password.length < 7){
      this.message = "Password must be at least 7 characters or more"
      this.err('pass')
    }
    else{
      var len:number = password.length
      var cap:boolean = false, sym:boolean=false,num:boolean=false
      const regex = new RegExp('^[a-zA-Z0-9]+$')
      if(!regex.test(password)){
        console.log('?')
        for (let i:number=0; i<len;i++){
          var temp:string = password.charAt(i)
          if(!cap && temp >= 'A' && temp<='Z'){
            cap = true
          }
          else if(!num && temp >= '0' && temp<='9')num=true

          if(cap&&num)return true
        }

      }
    }

    return false
  }
  validatePhone(phone:string):boolean{
    var len:number =phone.length
    for(let i:number = 0;i<len;i++){
      if(phone.charAt(i)<'0' || phone.charAt(i) >'9'){
        return false
      }
    }
    return true;
  }
  remo(id:string){
    this.message = "";
    this.err(id)
  }
  regis$:Subscription;

  register():void{
    var fname:string =(<HTMLInputElement>document.getElementById('fname')).value;
    var lname:string =(<HTMLInputElement>document.getElementById('lname')).value;
    var temp = <HTMLSelectElement>document.getElementById('select')
    var countrycode:string =temp.options[temp.selectedIndex].value;
    var pass:string =(<HTMLInputElement>document.getElementById('pass')).value;
    var phone:string =(<HTMLInputElement>document.getElementById('phone')).value;
    
    if(fname.length == 0){
      this.message = "first name should not be empty!"
      this.err('fname')
    }
    else if(phone.length < 7 || !this.validatePhone(phone)){
      this.message = "please enter valid phone number"
      this.err('phone')
    }
    else if(!this.validatePassword(pass)){
      this.message = "password not valid"
      this.err('pass')
    }
    else{
      if(this.message.length == 0){
        phone = countrycode.concat(phone)
        let user:User = new User(fname,lname,this.data.data,phone,pass);

        this.regis$ = this.apollo.createUser(user).subscribe(({data}) => {
          console.log(data)
          this.ref.close()
        })
      }
    }
  }

}
