import { LoginComponent } from './../LoginRegister/login/login.component';
import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogRef } from '@angular/material';
import { RegisterComponent } from '../LoginRegister/register/register.component';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {

  constructor(
    public dialog: MatDialog
  ) { }

  private refLogin:MatDialogRef<LoginComponent>
  private refRegister: MatDialogRef<RegisterComponent>
  lang:string = "../../../assets/Nav/flag-en.png"
  currency:string="IDR"
  private userData:number;
  ngOnInit() {
    this.userData = -1;
  }

  drop(event): void {
    let obj = event.target.nextSibling;
    console.log(obj);
    if (obj.className.search('show') == -1) {
      obj.className += ' show';
    } else {
      obj.className = 'drop';
    }
  }
  closeDropdown():void{
    //  var elements = Array.from(document.getElementsByClassName('show'))
    //  console.log('?')
    //  for (let element of elements){
    //    console.log(element)
    //    element.classList.remove('show');
    //  }
  }

  onLogin(){
    
    this.refLogin = this.dialog.open(LoginComponent)

    this.refLogin.afterClosed().subscribe(temp =>{
      if(temp.fromModal && !temp.logged){
        console.log(temp)
        this.refRegister =this.dialog.open(RegisterComponent,{data:temp} )
      }
      else if(temp.logged){
        this.userData = Number(sessionStorage.getItem("token"))
        console.log("loged user:"+ this.userData)
      }
    })
  }
  onRegis(){
    this.refRegister = this.dialog.open(RegisterComponent)
    
  }
}
