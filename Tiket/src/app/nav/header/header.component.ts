import { LoginComponent } from './../LoginRegister/login/login.component';
import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {

  constructor(
    public dialog: MatDialog
  ) { }

  ngOnInit() {

  }
  lang:string = "../../../assets/Nav/flag-en.png"
  currency:string="IDR"

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
    
    this.dialog.open(LoginComponent)
  }
}
