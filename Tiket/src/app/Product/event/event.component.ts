import { ApolloService } from 'src/app/Services/apollo.service';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-event',
  templateUrl: './event.component.html',
  styleUrls: ['./event.component.scss']
})
export class EventComponent implements OnInit {

  activities:any[]
  constructor(
    private apollo:ApolloService
  ) { }

  ngOnInit() {
    
    this.apollo.selectEventByCategory('activity').subscribe(async a=> {
      console.log(a)
    })

  }

}
