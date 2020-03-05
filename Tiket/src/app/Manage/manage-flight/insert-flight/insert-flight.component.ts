import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { ApolloService } from 'src/app/Services/apollo.service';
import { MatDialogRef } from '@angular/material';
import { FlightData } from 'src/app/Interfaces/flight-interface';

@Component({
  selector: 'app-insert-flight',
  templateUrl: './insert-flight.component.html',
  styleUrls: ['./insert-flight.component.scss']
})
export class InsertFlightComponent implements OnInit {
  private arrivalTimeFC:FormControl
  private departureTimeFC:FormControl
  private error:string
  // 
  constructor(private apollo: ApolloService,private ref:MatDialogRef<InsertFlightComponent> ) { }

  ngOnInit() {
    this.arrivalTimeFC = new FormControl()
    this.departureTimeFC = new FormControl()
  }

  err(a:string){
    this.error= a
  }
  insert(){
    this.error=""
    var name:string = (<HTMLInputElement>document.getElementById('name')).value
    var model:string = (<HTMLInputElement>document.getElementById('model')).value
    var arrStat:string = (<HTMLInputElement>document.getElementById('arriveAirport')).value
    var depStat:string = (<HTMLInputElement>document.getElementById('departureAirport')).value
    var transitStat:string = (<HTMLInputElement>document.getElementById('transitAirport')).value

    var arriveDate:Date = new Date((<HTMLInputElement>document.getElementById('arriveDate')).value)
    var departureDate:Date = new Date((<HTMLInputElement>document.getElementById('departureDate')).value)
    
    if(arriveDate.toString() == "Invalid Date" || departureDate.toString() == "Invalid Date"){
      this.err("Invalid Date!")
      return;
    }

    arriveDate.setHours(arriveDate.getHours() + 7);
    var arriveTime = arriveDate.toISOString().substr(0, 11) + this.arrivalTimeFC.value + ':00Z'

    departureDate.setHours(departureDate.getHours() + 7);
    var departureTime = departureDate.toISOString().substr(0, 11) + this.departureTimeFC.value + ':00Z'

    var price:number = parseFloat((<HTMLInputElement>document.getElementById('price')).value)

    
    if(name.length ==0){
      this.err("Flight Company Name Cannot Be Empty!")
    }
    else if(model.length == 0){
      this.err("Please Fill Flight's Model")
    }
    else if(isNaN(price)){
      this.err("Invalid Flight Price")
    }
    else if(arrStat.length==0){
      this.err("Arrival Station must be filled!")
    }
    else if(depStat.length == 0){
      this.err("Destination Station Must be filled")
    }
    else{
      var flight: FlightData = {
        id: 0,
        companyName: name,
        companyIcon: '',
        model: model,
        price: price,
        transit: transitStat,
        from: depStat,
        to: arrStat,
        departDate: departureTime,
        arriveDate: arriveTime,
        timeline: '',
        duration: ''
      };
      console.log(flight)
      this.apollo.createFlight(flight).subscribe(async (data:any)=>{
        
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
    console.log(res)
    this.close()
  }
}
