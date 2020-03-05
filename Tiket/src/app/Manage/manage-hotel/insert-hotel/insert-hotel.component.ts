import { Hotel } from './../../../Models/hotel';
import { Subscription } from 'rxjs';
import { ApolloService } from './../../../Services/apollo.service';
import { HotelFacility } from './../../../Interfaces/hotel-facility';
import { HotelType } from './../../../Interfaces/hotel-type';
import { Component, OnInit } from '@angular/core';
import {MatChipInputEvent} from '@angular/material/chips';
import {COMMA, ENTER} from '@angular/cdk/keycodes';
import { MatDialogRef, MatDialog } from '@angular/material/dialog';
import { NotificationComponent } from 'src/app/General/notification/notification.component';

@Component({
  selector: 'app-insert-hotel',
  templateUrl: './insert-hotel.component.html',
  styleUrls: ['./insert-hotel.component.scss']
})

export class InsertHotelComponent implements OnInit {
  refNotif: MatDialogRef<NotificationComponent, any>;

  constructor(private apollo: ApolloService, 
              private ref:MatDialogRef<InsertHotelComponent>,
              private dialog: MatDialog) { }

  ngOnInit() {
  }
  visible = true;
  selectable = true;
  removable = true;
  addOnBlur = true;
  readonly separatorKeysCodes: number[] = [ENTER, COMMA];
  hotelTypes: HotelType[] = [];
  hotelFacilities: HotelFacility[] = [];

  error:string="";

  typeinsert$:Subscription
  hotelinsert$:Subscription
  imageinsert$:Subscription
  facinsert$:Subscription

  addType(event: MatChipInputEvent): void {
    const input = event.input;
    const value = event.value;

    // Add our fruit
    if ((value || '').trim()) {
      this.hotelTypes.push({name: value.trim(),path: value.trim()+".jpg"});
    }
    // Reset the input value
    if (input) {
      input.value = '';
    }
  }

  removeType(h:HotelType): void {
    const index = this.hotelTypes.indexOf(h);

    if (index >= 0) {
      this.hotelTypes.splice(index, 1);
    }
  }
  addFacility(event: MatChipInputEvent): void {
    const input = event.input;
    const value = event.value;

    // Add our fruit
    if ((value || '').trim()) {
      this.hotelFacilities.push({name: value.trim(),path: value.trim()+".jpg"});
    }

    // Reset the input value
    if (input) {
      input.value = '';
    }
  }

  removeFacility(h:HotelFacility): void {
    const index = this.hotelFacilities.indexOf(h);

    if (index >= 0) {
      this.hotelFacilities.splice(index, 1);
    }
  }

  err(a:string){
    this.error = a
  }
  close():void{
    this.ref.close({
      fromModal:true
    })
  }
  private image:string
  insert(){
    this.error=""
    var name:string = (<HTMLInputElement>document.getElementById('name')).value
    this.image= (<HTMLInputElement>document.getElementById('img')).value
    var location:string = (<HTMLInputElement>document.getElementById('city')).value
    var rating:number = parseFloat((<HTMLInputElement>document.getElementById('rating')).value)
    var lat:number = parseFloat((<HTMLInputElement>document.getElementById('lat')).value)
    var long:number = parseFloat((<HTMLInputElement>document.getElementById('long')).value)
    var price:number = parseFloat((<HTMLInputElement>document.getElementById('price')).value)
    var address:string = (<HTMLInputElement>document.getElementById('address')).value
    var desc:string = (<HTMLInputElement>document.getElementById('desc')).value
    
    if(name.length ==0){
      this.err("Hotel Name Cannot Be Empty!")
    }
    else if(location.length == 0){
      this.err("Please Fill Hotel's City For Location")
    }
    else if(isNaN(rating) || rating < 0 || rating>5){
      this.err("Invalid Hotel Rating")
    }
    else if(isNaN(lat) || lat < (-85) || lat>85){
      this.err("Invalid Hotel Latitude")
    }
    else if(isNaN(long) || long < (-180) || long>180){
      this.err("Invalid Hotel Langitude")
    }
    else if(isNaN(price)){
      this.err("Invalid Hotel Price")
    }
    else if(address.length==0){
      this.err("Address must be filled!")
    }
    else if(desc.length == 0){
      this.err("Please fill Hotel's description")
    }
    else{
      var h:Hotel = new Hotel()
      h.address = address
      h.desc = desc
      h.price= price
      h.lat = lat
      h.long = long
      h.rating = rating
      h.city = location
      h.name = name

      var hotelId:number
      this.hotelinsert$ = this.apollo.createHotel(h).subscribe(async (data:any)=>{
        
        await this.insertOther(data)
      })

    }
  }
  insertOther(h:any){
    if(h.data.createhotel.id == -1 || h.data.createhotel.id ==0){
      this.refNotif = this.dialog.open(NotificationComponent,{data: "Insert Failed, Hotel Description Must be at least 20 characters long"})
      return;
    }
    console.log(h)
    if(h.data.createhotel.id == 0){
      return
    }
    for (const hot of this.hotelTypes) {
      this.apollo.createHotelType(h.data.createhotel.id,hot.name, hot.path).subscribe()
      console.log(hot)
    }

    for (const hot of this.hotelFacilities) {
      this.apollo.createHotelFacility(h.data.createhotel.id,hot.name, hot.path).subscribe()
      console.log(hot)
    }
    console.log("image:", this.image)
    if(this.image !=''){
      this.apollo.createHotelImage(h.data.createhotel.id,this.image).subscribe()
      
    }
    this.close()
  }
}
