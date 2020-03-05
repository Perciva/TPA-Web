import { Component, OnInit, ViewChild } from '@angular/core';
import { Subscription } from 'rxjs';
import { MatDialogRef, MatDialog, MatPaginator, MatSort, MatTableDataSource } from '@angular/material';
import { ConfirmDialogComponent } from '../confirm-dialog/confirm-dialog.component';
import { ApolloService } from 'src/app/Services/apollo.service';
import { InsertFlightComponent } from './insert-flight/insert-flight.component';
import { ChatServiceService } from 'src/app/Services/chat-service.service';

@Component({
  selector: 'app-manage-flight',
  templateUrl: './manage-flight.component.html',
  styleUrls: ['./manage-flight.component.scss']
})
export class ManageFlightComponent implements OnInit {

  private columns= ["companyIcon", "companyName",  "arrival","departure","model","from","to", "price", "transit", "Update", "Delete"];
  private allflight$:Subscription
  private flights:any
  private dataSource:any
  private path:string
  private refInsert:MatDialogRef<InsertFlightComponent>
  private refDelete:MatDialogRef<ConfirmDialogComponent>
  constructor(private chat:ChatServiceService,private apollo: ApolloService, private dialog: MatDialog) { }

  @ViewChild(MatPaginator, {static: true}) paginator: MatPaginator;
  @ViewChild(MatSort, {static: true}) sorter: MatSort;

  ngOnInit() {
    this.path = "../../../assets/Flight/"
    this.dataSource = null
    this.getInitData()
  }
  ngOnDestroy(): void {
    //Called once, before the instance is destroyed.
    //Add 'implements OnDestroy' to the class.
    this.allflight$.unsubscribe()
  }
  getInitData(){
    this.allflight$ = this.apollo.selectAllFlight().subscribe(
      async res =>{
        await this.getAllflight(res)
      } 
    )
  }
  getAllflight(res:any){
    this.flights = res.data.getallflight
    this.dataSource = new MatTableDataSource(this.flights);
    this.dataSource.paginator = this.paginator
    this.dataSource.sort = this.sorter
    console.log(this.dataSource.data)
  }

  Update(hot:any){
    // UpdateFlight flightId:number, from:string,to:string, transit:string, arrival:string ,departure:string,model:string, price:number
    if(isNaN(hot.price)){
      alert("flight Price must not be Empty!")
    }
    else if(hot.company.name.length == 0 || hot.model.length == 0){
      alert("flight Company Name and Model must be filled!")
    }
    else if(new Date(hot.arrivalTime).toString() == "Invalid Date" || new Date(hot.departureTime).toString() == "Invalid Date"){
      alert("Invalid Arrival Or Departure Date")
    }
    else if(hot.fromAirport.code.length != 3 || hot.toAirport.code.length != 3){
      alert("Air Port Code must be Filled!")
    }
    else{
      this.apollo.UpdateFlight(hot.Id,hot.fromAirport.code,hot.toAirport.code,hot.transit.code,hot.arrivalTime, hot.departureTime, hot.model,hot.price).subscribe()
      alert("Update Success!")
    }
    
  }
  Delete(hot:any){
    this.refDelete = this.dialog.open(ConfirmDialogComponent, {data:{
      message:"Are you sure you want to delete this flight?",
      cancel: "No",
      accept: "Yes"
    }})

    this.refDelete.afterClosed().subscribe( res =>{
      if(res){
        this.apollo.DeleteFlight(hot.Id).subscribe(async a =>{
          await this.getInitData()
        })
      }
    })
  }

  Insert(){
    this.refInsert = this.dialog.open(InsertFlightComponent)
    this.refInsert.afterClosed().subscribe(data =>{
      if(data.fromModal== true){
        this.chat.emit("flight","reload dong")
        this.getInitData()
      }
    })
  }
  
  applyFilter(val: string){
    val = val.trim()
    val = val.toLowerCase()
    this.dataSource.filter = val
  }

}
