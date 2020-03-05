import { NotificationComponent } from './../../General/notification/notification.component';
import { ChatServiceService } from './../../Services/chat-service.service';
import { ApolloService } from 'src/app/Services/apollo.service';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MatDialog, MatDialogRef } from '@angular/material';
import { InsertBlogComponent } from 'src/app/Manage/manage-blog/insert-blog/insert-blog.component';


@Component({
  selector: 'app-main-blog',
  templateUrl: './main-blog.component.html',
  styleUrls: ['./main-blog.component.scss']
})
export class MainBlogComponent implements OnInit {
  private refInsert:MatDialogRef<InsertBlogComponent>
  private refNotif:MatDialogRef<NotificationComponent>
  constructor(private chat:ChatServiceService, private apollo:ApolloService, private router:Router, private dialog: MatDialog) { }

  private allData:any
  ngOnInit() {
    this.getInitData()
    this.chat.listen('blog').subscribe(a =>{
      this.refNotif = this.dialog.open(NotificationComponent, {data:"New Blog!"})
      this.refNotif.afterClosed().subscribe(a =>{
        
        this.getInitData()
      })
    })
  }
  getInitData(){
    this.apollo.selectAllBlog().subscribe(async a =>{
      await this.getData(a.data.getallblog)
    })
  }
  getData(a:any){
    this.allData = a
  }

  ToDetail(b: any) {
    console.log(b)
    this.router.navigate(['/Blog/Detail'], {
      queryParams: {
        id: b.id,
      }
    });
  }
  createBlog(){
    this.refInsert = this.dialog.open(InsertBlogComponent)
    this.refInsert.afterClosed().subscribe(data =>{
      if(data.fromModal== true){
        console.log("Reload")
        this.getInitData()
      }
    })
  }

}
