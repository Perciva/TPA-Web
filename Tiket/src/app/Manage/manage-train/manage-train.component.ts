import { Component, OnInit, ViewChild } from '@angular/core';
import { Subscription } from 'rxjs';
import { ConfirmDialogComponent } from '../confirm-dialog/confirm-dialog.component';
import { InsertTrainComponent } from './insert-train/insert-train.component';
import { MatDialogRef, MatDialog, MatPaginator, MatSort, MatTableDataSource } from '@angular/material';
import { ApolloService } from 'src/app/Services/apollo.service';

@Component({
  selector: 'app-manage-train',
  templateUrl: './manage-train.component.html',
  styleUrls: ['./manage-train.component.scss']
})
export class ManageTrainComponent implements OnInit {

  private columns= ["name","arrival","arrivalto","departure","departfrom","class","price","seat","Update","Delete"]
  private allTrain$:Subscription
  private trains:any
  private dataSource:any
  private refInsert:MatDialogRef<InsertTrainComponent>
  private refDelete:MatDialogRef<ConfirmDialogComponent>
  constructor(private apollo: ApolloService, private dialog: MatDialog) { }

  @ViewChild(MatPaginator, {static: true}) paginator: MatPaginator;
  @ViewChild(MatSort, {static: true}) sorter: MatSort;

  ngOnInit() {
    this.dataSource = null
    this.getInitData()
  }
  ngOnDestroy(): void {
    //Called once, before the instance is destroyed.
    //Add 'implements OnDestroy' to the class.
    this.allTrain$.unsubscribe()
  }
  getInitData(){
    this.allTrain$ = this.apollo.selectAllTrain().subscribe(
      async res =>{
        await this.getAllTrain(res)
      } 
    )
  }
  getAllTrain(res:any){
    this.trains = res.data.getalltrain
    this.dataSource = new MatTableDataSource(this.trains);
    this.dataSource.paginator = this.paginator
    this.dataSource.sort = this.sorter
    console.log(this.dataSource.data)
  }

  Update(hot:any){
    // UpdateTrain(trainId:number, arrival:string, departure:string,price:number, seat:number)
    // console.log(new Date(hot.departureTime).toString())
    // var temp:any = hot.price
    if(isNaN(hot.seat) || isNaN(hot.price)){
      alert("Train Price and Rating must not be Empty!")
    }
    else if(hot.name.length == 0 || hot.code.length == 0){
      alert("Train Name and Code must be filled!")
    }
    else if(new Date(hot.arrivalTime).toString() == "Invalid Date" || new Date(hot.departureTime).toString() == "Invalid Date"){
      alert("Invalid Arrival Or Destination Date")
    }
    else{
      this.apollo.UpdateTrain(hot.Id,hot.arrivalTime,hot.departureTime,hot.price,hot.seat).subscribe()
      alert("Update Success!")
    }
    
  }
  Delete(hot:any){
    this.refDelete = this.dialog.open(ConfirmDialogComponent, {data:{
      message:"Are you sure you want to delete this Train?",
      cancel: "No",
      accept: "Yes"
    }})

    this.refDelete.afterClosed().subscribe( res =>{
      if(res){
        this.apollo.DeleteTrain(hot.Id).subscribe(async a =>{
          await this.getInitData()
        })
      }
    })
  }

  Insert(){
    this.refInsert = this.dialog.open(InsertTrainComponent)
    this.refInsert.afterClosed().subscribe(data =>{
      if(data.fromModal== true){
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
