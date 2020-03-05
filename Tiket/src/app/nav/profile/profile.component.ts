import { PhoneVerifyService } from './../../Services/phone-verify.service';
import { ApolloService } from 'src/app/Services/apollo.service';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.scss']
})
export class ProfileComponent implements OnInit {
  private userId:number
  private User:any

  constructor(
    private apollo: ApolloService,
    private phone:PhoneVerifyService
    ) { 
    this.userId = parseInt(sessionStorage.getItem("token"))

   
  }


  ngOnInit() {
    

    this.User = this.apollo.selectUserById(this.userId).subscribe(async a=>{
      await this.getUserData(a.data.userbyid)
    })
  }
  getUserData(a:any){
    this.User = a
  }
  change(a:any){
    console.log(a)
  }
  update(){
    this.phone.verifPhoneNumber(this.User.phone).subscribe( (a:any) => {
    
      if(a.success== false || a.valid == false){
        alert("Invalid Phone Number!")
      }
      else{
        this.apollo.updateUser(this.User.address, this.User.email,this.User.firstname,this.User.lastname,
          this.userId,this.User.kode_pos,this.User.phone,this.User.title,this.User.languange).subscribe(async a =>{
            console.log(a)
          })
      }
    })

    
  }
  subscribe(){
    this.apollo.subscribe(this.User.email).subscribe(a=>{
      console.log(a)
    })
  }

}