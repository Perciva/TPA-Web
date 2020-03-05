import { ChatServiceService } from './../../Services/chat-service.service';
import { TextEditorComponent } from './../manage-blog/text-editor/text-editor.component';
import { Component, OnInit, ViewChild } from '@angular/core';
import { Subscription } from 'rxjs';
import { MatDialogRef, MatDialog, MatPaginator, MatSort, MatTableDataSource } from '@angular/material';
import { InsertEventComponent } from './insert-event/insert-event.component';
import { ConfirmDialogComponent } from '../confirm-dialog/confirm-dialog.component';
import { ApolloService } from 'src/app/Services/apollo.service';
import { UpdateEventComponent } from './update-event/update-event.component';
import { ElementSchemaRegistry } from '@angular/compiler';
import { EventData } from 'src/app/Interfaces/event-interface';

@Component({
  selector: 'app-manage-event',
  templateUrl: './manage-event.component.html',
  styleUrls: ['./manage-event.component.scss']
})
export class ManageEventComponent implements OnInit {

  private columns= ['Image', 'title','category','location','date','price','content','Edit','Update','Delete']
  private allEvent$:Subscription
  private path:string
  private Events:any
  private dataSource:any
  private refInsert:MatDialogRef<InsertEventComponent>
  private refDelete:MatDialogRef<ConfirmDialogComponent>
  private refEdit:MatDialogRef<UpdateEventComponent>
  constructor(private chat: ChatServiceService ,private apollo: ApolloService, private dialog: MatDialog) { }

  @ViewChild(MatPaginator, {static: true}) paginator: MatPaginator;
  @ViewChild(MatSort, {static: true}) sorter: MatSort;


  ngOnInit() {
    this.dataSource = null;
    this.getInitData()
  }
  ngOnDestroy(): void {
    //Called once, before the instance is destroyed.
    //Add 'implements OnDestroy' to the class.
    this.allEvent$.unsubscribe()
  }
  getInitData(){
    this.allEvent$ = this.apollo.selectAllEvent().subscribe(
      async res =>{
        await this.getAllEvent(res)
      } 
    )
  }
  getAllEvent(res:any){
    console.log(res)
    this.Events = res.data.getallentertainment
    this.dataSource = new MatTableDataSource(this.Events);
    this.dataSource.paginator = this.paginator
    this.dataSource.sort = this.sorter
    console.log(this.dataSource.data)
    
  }
  Insert(){
    this.refInsert = this.dialog.open(InsertEventComponent)
    this.refInsert.afterClosed().subscribe(data =>{
      if(data.fromModal== true){
        console.log("Reload")
        this.getInitData()
      }
    })
  }
  Edit(b:any){
    this.refEdit = this.dialog.open(UpdateEventComponent, {data:{
      title: b.title,
      desc: b.content
    }})
    this.refEdit.afterClosed().subscribe( res =>{
      if(res.fromModal){
        b.title = res.title,
        b.content = res.content
      }
    })
  }
  Update(b:any){
    console.log(new Date(b.date).toString())
    if(b.content.length == 0 || b.title.length == 0){
      alert("Event Content And Title must be filled!")
    }
    else if(isNaN(b.price)){
      alert("Price must not be empty")
    }
    else if(new Date(b.date).toString() == "Invalid Date"){
      alert("Invalid Event Date!")
    }
    else{
      var event: EventData = {
        type: b.type,
        date: b.date,
        image: b.image,
        latitude: b.latitude,
        location: b.location,
        longitude: b.longitude,
        price: b.price,
        title: b.title,
        description: b.content,
        id: b.id
      };
      this.apollo.updateEvent(event).subscribe()
    }
  }
  Delete(b:any){
    this.refDelete = this.dialog.open(ConfirmDialogComponent, {data:{
      message:"Are you sure you want to delete this event?",
      cancel: "No",
      accept: "Yes"
    }})

    this.refDelete.afterClosed().subscribe( res =>{
      if(res){
        this.apollo.deleteEvent(b.id).subscribe(async a =>{
          this.chat.emit('event',"reload")
          await this.getInitData()
        })
      }
    })
  }


}
