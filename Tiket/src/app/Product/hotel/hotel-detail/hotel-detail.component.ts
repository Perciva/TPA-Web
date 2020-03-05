import { Subscription } from 'rxjs';
import { Component, OnInit } from '@angular/core';
import { ApolloService } from 'src/app/Services/apollo.service';
import { ActivatedRoute } from '@angular/router';
import * as L from 'leaflet';

@Component({
  selector: 'app-hotel-detail',
  templateUrl: './hotel-detail.component.html',
  styleUrls: ['./hotel-detail.component.scss']
})
export class HotelDetailComponent implements OnInit {
  private hotelId:number
  private Hotel$:Subscription
  private hotel:any
  private path:string
  private radius:number[] = []
  private hotelList:any
  private tes:number
  userLocation: Position;
  map: any;
  constructor(
    private activatedRoute: ActivatedRoute,
    private apollo: ApolloService
    ) { 
      this.path="../../../../assets/Hotel/"
      this.activatedRoute.queryParams.subscribe(async p => {
        await this.getParam(p.id);
    });
    
  }

  getParam(p:number){
    this.hotelId = p
  }

  ngOnInit() {
    this.Hotel$ = this.apollo.searchHotelById(this.hotelId).subscribe(async a=>{
      await this.getHotel(a.data.hotelbyid)
    })
    

      
 
  }
  fetch(res:any){
    this.hotelList = res.data.getnearbyhotel
    console.log(this.hotelList)
  }
  getHotel(a:any){
    console.log(a)
    this.hotel = a

    this.apollo.getNearbyHotel(this.hotel.lat, this.hotel.long).subscribe(
      async res => {
        console.log(res)
        await this.fetch(res)
        // this.hotelList = res.data.getnearbyhotel;???wai
        // this.hotelList = res.data.
      }
    )

    this.map = L.map('map', {zoomControl: false, draggable: false}).setView([this.hotel.lat, this.hotel.long], 8);
      L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
      }).addTo(this.map);

      const temp = L.marker([this.hotel.lat, this.hotel.long]);
      temp.addTo(this.map);
      this.apollo.getNearbyHotel(this.hotel.lat,
        this.hotel.long).subscribe( async (result) => {
          this.getNearData(result.data.getnearbyhotel);
      })
      // this.map.setState({ draggable: false }
  }
  getDistanceFromLatLonInKm(lat1, lon1, lat2, lon2) {
    var R = 6371; // Radius of the earth in kilometers
    var dLat = this.deg2rad(lat2 - lat1); // deg2rad below
    var dLon = this.deg2rad(lon2 - lon1);
    var a =
      Math.sin(dLat / 2) * Math.sin(dLat / 2) +
      Math.cos(this.deg2rad(lat1)) * Math.cos(this.deg2rad(lat2)) *
      Math.sin(dLon / 2) * Math.sin(dLon / 2);
    var c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));
    var d = R * c; // Distance in KM
    return d;
  }
  
  deg2rad(deg) {
    return deg * (Math.PI / 180)
  }
  getNearData(a:any){
    let index=0
    for(let hot of a){
      console.log(hot.name)
      this.radius[index++]=this.getDistanceFromLatLonInKm(hot.lat, hot.long, this.hotel.lat, this.hotel.long)
      
      // this.radius[index++]=Math.acos(Math.sin(hot.lat)*Math.sin(this.hotel.lat)+Math.cos(hot.lat)*Math.cos(this.hotel.lat)*
      //       Math.cos(hot.long-this.hotel.long)
      //       )*6371
    }
    console.log("rad",this.radius)
  }

  fb() {
    window.open('http://www.facebook.com/sharer.php?u=localhost:4200/Hotel/Detail?id=' + this.hotelId, 'facebookShare', 'width=626,height=436');
  }

  wa() {
    window.open('https://api.whatsapp.com/send?text=localhost:4200/Hotel/Detail?id=' + this.hotelId)
  }

  copy() {
    navigator.clipboard.writeText('localhost:4200/Hotel/Detail?id=' + this.hotelId);
  }

  line() {
    window.open('http://line.me/R/msg/text/?Hotel/Detail?id=' + this.hotelId);
  }

}
