import { InsertHotelComponent } from './insert-hotel/insert-hotel.component';
import { isNumber } from 'util';
import { Subscription } from 'rxjs';
import { ApolloService } from './../../Services/apollo.service';
import { Component, OnInit, ViewChild } from '@angular/core';
import { MatTableDataSource, MatTab, MatDialog, MatDialogRef, MatPaginator } from '@angular/material';
import { AttachSession } from 'protractor/built/driverProviders';
import { ConfirmDialogComponent } from '../confirm-dialog/confirm-dialog.component';


@Component({
  selector: 'app-manage-hotel',
  templateUrl: './manage-hotel.component.html',
  styleUrls: ['./manage-hotel.component.scss']
})
export class ManageHotelComponent implements OnInit {
  private columns= ['Image', 'Name','Description','Price','Rating','Facilities','Address','Type','Update','Delete']
  private allHotel$:Subscription
  private path:string
  private hotels:any
  private dataSource:any
  private refInsert:MatDialogRef<InsertHotelComponent>
  private refDelete:MatDialogRef<ConfirmDialogComponent>
  constructor(private apollo: ApolloService, private dialog: MatDialog) { }

  @ViewChild(MatPaginator, {static: true}) paginator: MatPaginator;
  ngOnInit() {
    this.path = "../../../assets/Hotel/"
    this.dataSource = null
    this.getInitData()
  }
  getInitData(){
    this.allHotel$ = this.apollo.selectAllHotel().subscribe(
      async res =>{
        await this.getAllHotel(res)
      } 
    )
  }
  getAllHotel(res:any){
    this.hotels = res.data.allHotel
    this.dataSource = new MatTableDataSource(this.hotels);
    this.dataSource.paginator = this.paginator
    console.log(this.dataSource.data)
    for(let h of this.dataSource.data){
      for(let i of h.type){
        console.log(i.name)
      }
    }
  }
  validateNumber(num:string):boolean{
    var len:number =num.length
    for(let i:number = 0;i<len;i++){
      if(num.charAt(i)<'0' || num.charAt(i) >'9'){
        return false
      }
    }
    return true;
  }
  Update(hot:any){
    
    var temp:any = hot.price
    if(isNaN(hot.rating) || isNaN(hot.price)){
      alert("Hotel Price and Rating must not be Empty!")
    }
    else if(hot.rating > 5 ||hot.rating < 0){
      alert("Invalid Hotel Rating")
    }
    else if(hot.name.length == 0 || hot.description.length == 0){
      alert("Hotel Name and Description must be filled!")
    }
    else{
      this.apollo.UpdateHotel(hot.rating,hot.description,hot.id,hot.name,hot.price).subscribe()
      alert("Update Success!")
    }
    
  }
  Delete(hot:any){
    this.refDelete = this.dialog.open(ConfirmDialogComponent, {data:{
      message:"Are you sure you want to delete this hotel?",
      cancel: "No",
      accept: "Yes"
    }})

    this.refDelete.afterClosed().subscribe( res =>{
      if(res){
        this.apollo.DeleteHotel(hot.id).subscribe(async a =>{
          await this.getInitData()
        })
      }
    })
  }

  Insert(){
    this.refInsert = this.dialog.open(InsertHotelComponent)
    this.refInsert.afterClosed().subscribe(data =>{
      if(data.fromModal== true){
        console.log("Reload")
        this.getInitData()
      }
    })
  }
  checkHotelPhoto(h:any):boolean{
    return h.photo.length > 0
  }
  applyFilter(val: string){
    val = val.trim()
    val = val.toLowerCase()
    this.dataSource.filter = val
  }

}