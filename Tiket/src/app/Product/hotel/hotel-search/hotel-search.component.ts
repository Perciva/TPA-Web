import { MatDialog } from '@angular/material';
import { ChatServiceService } from './../../../Services/chat-service.service';
import { Subscription } from 'rxjs';
import { ApolloService } from 'src/app/Services/apollo.service';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Options, LabelType } from 'ng5-slider';
import { NotificationComponent } from 'src/app/General/notification/notification.component';

@Component({
  selector: 'app-hotel-search',
  templateUrl: './hotel-search.component.html',
  styleUrls: ['./hotel-search.component.scss']
})
export class HotelSearchComponent implements OnInit {
  private top = {
    a : 123,
    b: "asd"
  }
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
  refNotif: any;


  constructor(
    private activatedRoute: ActivatedRoute,
    private router: Router,
    private apollo: ApolloService,
    private chat: ChatServiceService,
    private dialog:MatDialog
    ) { 
      this.activatedRoute.queryParams.subscribe(async p => {
        await this.getData(p);
      });
      // this.getHotel()
      this.chat.listen('hotel').subscribe(a =>{
        this.refNotif = this.dialog.open(NotificationComponent, {data:"New Hotel!"})
        this.refNotif.afterClosed().subscribe(a =>{
          
          this.changed()
        })
      })
  }
  ngOnInit() {
    this.path = "../../../../assets/Hotel/"
    this.SortBy= "Recommended"
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
      await this.getHotelData(a.data.getfilteredhotel)
    })
  }
  resetFilter(){
    this.min = 0
    this.max = 5000000
    this.facilityModel = [false]
    this.locationModel = [false]
    this.ratingsModel = [false]
    this.categoryModel = [false]

    this.changed()
    
  }

  getHotel(){
    this.getHotels$ = this.apollo.selectAllHotel().subscribe(async a => {
      await this.getHotelData(a.data.allHotel)
    })
  }
  getHotelData(a:any){
    this.hotels = a
    console.log(this.hotels)
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
    console.log(p)
    console.log(p.berangkat)
    if(p.from == undefined){
      this.router.navigateByUrl("/Train")
    }
    else{
      this.p = p 
    }
    console.log(p)
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
  toMap(){
    this.router.navigate(['/Hotel/Map'])
  }
  

}
