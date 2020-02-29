import { Subscription } from 'rxjs';
import { ApolloService } from './../../Services/apollo.service';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-rec-hotel',
  templateUrl: './rec-hotel.component.html',
  styleUrls: ['./rec-hotel.component.scss']
})
export class RecHotelComponent implements OnInit {

  private rec$:Subscription
  private hotelList:any
  private path:string

  constructor(private apollo: ApolloService) { 
    navigator.geolocation.getCurrentPosition(curr => {
      // console.log("?")
      // console.log(curr)
      this.rec$ = this.apollo.getNearbyHotel(curr.coords.latitude, curr.coords.longitude).subscribe(
        async res => {
          console.log(res)
          await this.fetch(res)
          // this.hotelList = res.data.getnearbyhotel;???wai
          // this.hotelList = res.data.
        }
      )
    })
  }

  ngOnInit() {
    this.path= "../../../assets/Hotel/"
    this.hotelList=null
    console.log(this.hotelList)
    
  }
  ngOnDestroy(): void {
    //Called once, before the instance is destroyed.
    //Add 'implements OnDestroy' to the class.
    this.rec$.unsubscribe()
  }
  fetch(res:any){
    this.hotelList = res.data.getnearbyhotel
    console.log(this.hotelList)
  }

}
