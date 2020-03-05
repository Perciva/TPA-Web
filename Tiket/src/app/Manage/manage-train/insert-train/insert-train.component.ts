import { ApolloService } from 'src/app/Services/apollo.service';
import { TrainData } from './../../../Interfaces/train-interface';
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { MatDialogRef } from '@angular/material';

@Component({
  selector: 'app-insert-train',
  templateUrl: './insert-train.component.html',
  styleUrls: ['./insert-train.component.scss']
})
export class InsertTrainComponent implements OnInit {
  private arrivalTimeFC:FormControl
  private departureTimeFC:FormControl
  private error:string
  
  constructor(private apollo: ApolloService,private ref:MatDialogRef<InsertTrainComponent> ) { }

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
    var code:string = (<HTMLInputElement>document.getElementById('code')).value
    var arrStat:string = (<HTMLInputElement>document.getElementById('arriveStation')).value
    var depStat:string = (<HTMLInputElement>document.getElementById('departureStation')).value
    var transitStat:string = (<HTMLInputElement>document.getElementById('transitStation')).value
    var trclass:string = (<HTMLInputElement>document.getElementById('class')).value

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
    var seat:number = parseFloat((<HTMLInputElement>document.getElementById('seat')).value)

    
    if(name.length ==0){
      this.err("Train Name Cannot Be Empty!")
    }
    else if(code.length == 0){
      this.err("Please Fill Train's Code")
    }
    else if(isNaN(price)){
      this.err("Invalid Train Price")
    }
    else if(isNaN(seat)){
      this.err("Please Insert Number of Seat in train")
    }
    else if(arrStat.length==0){
      this.err("Arrival Station must be filled!")
    }
    else if(depStat.length == 0){
      this.err("Destination Station Must be filled")
    }
    else if(trclass.length == 0){
      this.err("Train Class Must Be Filled!")
    }
    else{
      var h:TrainData = {
        departureName: depStat,
        arrivalName: arrStat,
        price: price,
        nameCode: name + ',' + code,
        transit: transitStat,
        class: trclass,
        seat: seat,
        id: 0,
        departureTime: departureTime,
        arrivalTime: arriveTime
      };

      this.apollo.createTrain(h).subscribe(async (data:any)=>{
        
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
