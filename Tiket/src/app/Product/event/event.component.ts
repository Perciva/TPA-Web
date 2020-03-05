import { ApolloService } from 'src/app/Services/apollo.service';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-event',
  templateUrl: './event.component.html',
  styleUrls: ['./event.component.scss']
})
export class EventComponent implements OnInit {
  
  activities:any[]
  attractions:any[]
  events:any[]
  constructor(
    private apollo:ApolloService
  ) { }

  ngOnInit() {
    
    this.apollo.selectEventByCategory('activity').subscribe(async a=> {
      console.log(a)
      await this.fetch(a.data.geteventbycategory)
      //a.data.geteventbycategory
    })

  }
  fetch(a){
    this.activities = a

    this.apollo.selectEventByCategory('attraction').subscribe(async a=> {
      await this.fetch1(a.data.geteventbycategory)
      //a.data.geteventbycategory
    })
  }
  fetch1(a){
    this.attractions = a

    this.apollo.selectEventByCategory('events').subscribe(async a=> {
      await this.fetch2(a.data.geteventbycategory)
      //a.data.geteventbycategory
    })
  }
  fetch2(a){
    this.events = a
  }

}
