import { Component, OnInit } from '@angular/core';
import { EventData } from 'src/app/Interfaces/event-interface';
import { FormControl } from '@angular/forms';
import { ApolloService } from 'src/app/Services/apollo.service';
import { MatDialogRef, MatDialog } from '@angular/material';
import { InsertFlightComponent } from '../../manage-flight/insert-flight/insert-flight.component';
import { NotificationComponent } from 'src/app/General/notification/notification.component';

@Component({
  selector: 'app-insert-event',
  templateUrl: './insert-event.component.html',
  styleUrls: ['./insert-event.component.scss']
})
export class InsertEventComponent implements OnInit {
  private error:string
  private category:string
  refNotif: MatDialogRef<NotificationComponent, any>;
  // 
  constructor(private dialog:MatDialog,private apollo: ApolloService,private ref:MatDialogRef<InsertFlightComponent> ) { }

  ngOnInit() {
    
  }

  err(a:string){
    this.error= a
  }
  insert(){
    // console.log(this.category)
    // return
    this.error=""
    var name:string = (<HTMLInputElement>document.getElementById('name')).value
    var img:string = (<HTMLInputElement>document.getElementById('img')).value
    var city:string = (<HTMLInputElement>document.getElementById('city')).value
    var lat:number = parseFloat((<HTMLInputElement>document.getElementById('lat')).value)
    var long:number = parseFloat((<HTMLInputElement>document.getElementById('long')).value)
    var price:number = parseFloat((<HTMLInputElement>document.getElementById('price')).value)
    var date:string = (<HTMLInputElement>document.getElementById('date')).value

    var desc:string = (<HTMLInputElement>document.getElementById('desc')).value
    if(name.length ==0){
      this.err("Please Fill Event Name")
    }
    else if(this.category.length == 0){
      this.err("Please Fill Event Category")
    }
    else if(isNaN(lat) || lat < (-85) || lat>85){
      this.err("Invalid Event Latitude")
    }
    else if(isNaN(long) || long < (-180) || long>180){
      this.err("Invalid Event Langitude")
    }
    else if(isNaN(price)){
      this.err("Invalid Event Price")
    }
    else if(new Date(date).toString() == "Invalid Date"){
      this.err("Invalid Date!")
    }
    else if(city.length == 0){
      this.err("Please fill Event's City Location")
    }
    else if(desc.length == 0){
      this.err("Please fill Event's Description")
    }
    else{
      // YYYY-MM-DDTHH:MM:SSZ
      var event: EventData = {
      type: this.category,
      date: new Date(date).toISOString().substr(0, 11) + '00:00:00Z',
      image: img,
      latitude: lat,
      location: city,
      longitude: long,
      price: price,
      title: name,
      description: desc,
      id: 0
    };
      console.log(event)
      this.apollo.CreateEvent(event).subscribe(async (data:any)=>{
        
        await this.after(data)
      })

    }
  }

  close():void{
    this.ref.close({
      fromModal:true
    })
  }

  after(res:any){
    if(res.data.createevent.id == -1 || res.data.createevent.id ==0){
      this.refNotif = this.dialog.open(NotificationComponent,{data: "Insert Failed, Event Description Must be at least 20 characters long"})
      return;
    }
    this.close()
  }
}
