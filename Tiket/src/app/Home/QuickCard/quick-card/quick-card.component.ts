import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-quick-card',
  templateUrl: './quick-card.component.html',
  styleUrls: ['./quick-card.component.scss']
})
export class QuickCardComponent implements OnInit {
  private highlight:boolean = false;
  private previousID:string = "";
  constructor() { }

  ngOnInit() {
    this.previousID = "imgPlane"
    document.getElementById('imgPlane').classList.add('blue')
  }

  clickHead(id:string){
    if(this.highlight){
      document.getElementById(this.previousID).classList.remove('blue')
      this.highlight=false;
    }
    document.getElementById(id).classList.add('blue')
    this.highlight=true;
    this.previousID=id;
  }

}
