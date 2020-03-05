import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { ApolloService } from 'src/app/Services/apollo.service';

@Component({
  selector: 'app-event-detail',
  templateUrl: './event-detail.component.html',
  styleUrls: ['./event-detail.component.scss']
})
export class EventDetailComponent implements OnInit {
  private currEvent:any

  constructor(private activatedRoute: ActivatedRoute,
    private router: Router,
    private apollo:ApolloService) { 

    this.activatedRoute.queryParams.subscribe(async a=>{
    await this.getParam(a)
    })
  }
  getParam(a:any){
    console.log(a)
    this.apollo.selectEventById(parseInt(a.id)).subscribe(async a=>{
      console.log(a)
    })
  }

  ngOnInit() {
  }

}
