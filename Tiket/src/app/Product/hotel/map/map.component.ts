import { Component, OnInit } from '@angular/core';
import * as L from 'leaflet';
import { Options, LabelType } from 'ng5-slider';
import { Subscription } from 'rxjs';
import { ActivatedRoute, Router } from '@angular/router';
import { ApolloService } from 'src/app/Services/apollo.service';
@Component({
  selector: 'app-map',
  templateUrl: './map.component.html',
  styleUrls: ['./map.component.scss']
})
export class MapComponent implements OnInit {
  private map:any 
  private userLocation: any;
  private min:number= 0
  private max:number=5000000;
  private options: Options = {
    floor: 0,
    ceil: 5000000,
    translate: (value: number, label: LabelType): string => {
      switch (label) {
        case LabelType.Low:
          return 'IDR ' + Math.floor(value / 1000) + 'K';
        case LabelType.High:
          return 'IDR ' + Math.floor(value / 1000) + 'K';
        default:
          return '';
      }
    }
  };
  private ratings:number[] = [1,2,3,4,5]
  private ratingsModel : boolean[] = []
  private p:any;
  private locations:any[] = [];
  private locationModel: boolean[] = [];
  private facility:string[] = ['AC','WiFi','Kulkas'];
  private facilityModel: boolean[] = [];
  private category:string[] = ['Hotel','Motel','Villa'];
  private categoryModel: boolean[] = [];
  private getLocation$:Subscription
  private getHotels$:Subscription
  private hotels
  private path:string;
  private SortBy:string
  private marker: L.marker[] = [];
  private showHot:any;
  markerLocations: any;
  filteredHotels: any;

  constructor(
    private activatedRoute: ActivatedRoute,
    private router: Router,
    private apollo: ApolloService
    ) { 
      this.activatedRoute.queryParams.subscribe(async p => {
        await this.getData(p);
      });
      
  }
  ngOnInit() {
    this.path = "../../../../assets/Hotel/"
    this.SortBy= "Recommended"
    this.filteredHotels = null
    navigator.geolocation.getCurrentPosition((succ) => {

      this.userLocation = succ;

      this.map = L.map('map', {zoomControl: false}).setView([this.userLocation.coords.latitude, this.userLocation.coords.longitude], 8);
      L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
      }).addTo(this.map);

      const temp = L.marker([this.userLocation.coords.latitude, this.userLocation.coords.longitude]);
      temp.addTo(this.map);
      this.getHotel()

      this.map.on('moveend', function(){

        let latitude = parseFloat(this.map.getCenter().lat)
        let longitude = parseFloat( this.map.getCenter().lng)
  
        this.apollo.getNearbyHotel(latitude,
          longitude).subscribe( async (result) => {
            this.getNewData(result.data.getnearbyhotel);
  
        })
      }.bind(this))

    },()=>{

    }) 
  }
  getNewData(result){
    console.log("len", result.length)
    this.marker.forEach(marker => {
      // this.map.removeLayer(marker)
      marker.remove()
    });
    var res:any[] = []
    if(this.filteredHotels != null){
      console.log("?")
      for(let z=0;z<result.length;z++){
        if(this.filteredHotels.includes(result[z].id)){
          console.log(result[z].id)
          res.push(result[z])
        }
      }
    }
    else{
      res = result
    }
    console.log(res)
    this.markerLocations = res
    console.log(this.markerLocations)
    
    let ctr =0
    for(let hot of this.markerLocations){
      this.marker[ctr] = new L.marker([hot.lat, hot.long]).
      bindPopup(String(hot.price)).openPopup();
        // bindToolTip(String(hot.price), {permanent:true})
      this.marker[ctr].on('click', (e) => {
          // alert(e.latlng.lat)
        this.apollo.getHotelByLatLong(e.latlng.lat, e.latlng.lng).subscribe(async q => {
          await this.changeHotel(q.data.hotelbylatlong);
        });
      });
      this.marker[ctr++].addTo(this.map)
    }
  }
  changeHotel(hot:any){
    console.log(hot)
    this.showHot = hot
  }
  log(){
    console.log(this.SortBy)
  }
  goToDetail(h:any){
    console.log(h)
  }
  changed(){
    var loc:number[]=[]
    var rat:number[]=[]
    var fac:string[] = []
    var pricemin:number = -1
    var pricemax:number = -1

    
    if(this.min!=0){
      pricemin = this.min
    }
    if(this.max!=5000000)pricemax = this.max

    for(let z=0;z<5;z++){
      if(this.ratingsModel[z]){
        rat.push(z+1)
      }
    }
    for(let z=0;z<this.locations.length;z++){
      if(this.locationModel[z]){
        loc.push(this.locations[z].id)
      }
    }

    this.getHotels$ = this.apollo.searchFilteredHotel(loc,rat, fac, pricemin,pricemax).subscribe(async a => {
      await this.getfilteredhotelData(a.data.getfilteredhotel)
    })


  }
  getfilteredhotelData(a:any){

    // console.log(a)
    this.filteredHotels = []
    for(let z=0;z<a.length;z++){
      this.filteredHotels.push(a[z].id)
    }
    console.log(this.filteredHotels)
    
  }
  sync(a:any){

  }
  getHotelData(a:any){
    console.log(a)
    this.hotels = a
    this.getMarker()
  }
  resetFilter(){
    this.min = 0
    this.max = 5000000
    this.facilityModel = [false]
    this.locationModel = [false]
    this.ratingsModel = [false]
    this.categoryModel = [false]
    this.filteredHotels = []

    this.changed()
    
  }
  getMarker(){
    let counter = 0

    for( let hot of this.hotels){
      this.marker[counter++] = new L.marker([hot.lat, hot.long])
    }
  }
  getHotel(){
    this.getHotels$ = this.apollo.getNearbyHotel(this.userLocation.coords.latitude,this.userLocation.coords.longitude).subscribe(async a => {
      await this.getHotelData(a.data.getnearbyhotel)
    })
  }
  
  getLoc(a:any){
    for(let z=0; z<a.length ; z++){
      this.locations.push(a[z])
    }
  }
  getLocData(){
    console.log(this.locations.length)
    this.getLocation$ = this.apollo.selectAllLocation().subscribe(
      async a=>{
        await this.getLoc(a.data.getallloc)
      }
    )
  }
  checkPhoto(path:any):boolean{
    return path.length != 0
  }
  getData(p:any){

    // if(p.from == undefined){
    //   this.router.navigateByUrl("/Hotel")
    // }
    // else{
    // }
    this.p = p 
  }
  
  
  search(){
    var bdate:Date = new Date((<HTMLInputElement>document.getElementById('berangkat')).value);
    var pdate:Date = new Date((<HTMLInputElement>document.getElementById('pulang')).value);
    var from:string = (<HTMLInputElement>document.getElementById('from')).value
    var temp: String[] = ((<HTMLInputElement>document.getElementById('pclass')).value).split(' ')
    if(temp.length !=2){
      alert("Mohon isi keterangan jumlah penumpang dan kelas kabin")
      return;
    }
    var people:number = parseInt(temp[0].toString());
    var room:string = temp[1].toString();

    if(from == ""){
      alert("Mohon Isi kolom Kolom dari")
    }
    else if(new Date(bdate).toString() == "Invalid Date"|| new Date(pdate).toString() == "Invalid Date"){
      alert("Tanggal Check In dan Check Out Invalid");
    }
    else if(isNaN(people) || room == ""){
      alert("Mohon isi Keterangan penumpang dan Jenis Ruangan");
    }
    else{
      this.router.navigate(["/Hotel/Search"], {
        queryParams:{
          'from': from,
          'berangkat': bdate.toString(),
          'pulang': pdate.toString(),
          'people': people,
          'room': room
        }
      })

    }
  }  
  toList(){
    this.router.navigate(['Hotel/Search/'])
  }

}
