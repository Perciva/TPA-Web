import { Component, OnInit, ViewChild } from '@angular/core';
import { Subscription } from 'rxjs';
import { MatDialogRef, MatDialog, MatPaginator, MatSort, MatTableDataSource } from '@angular/material';
import { InsertBlogComponent } from './insert-blog/insert-blog.component';
import { ConfirmDialogComponent } from '../confirm-dialog/confirm-dialog.component';
import { ApolloService } from 'src/app/Services/apollo.service';
import { ChatServiceService } from 'src/app/Services/chat-service.service';

@Component({
  selector: 'app-manage-blog',
  templateUrl: './manage-blog.component.html',
  styleUrls: ['./manage-blog.component.scss']
})
export class ManageBlogComponent implements OnInit {

  private columns= ['Image', 'Title','Content','Category','Update','Delete']
  private allblog$:Subscription
  private path:string
  private blogs:any
  private dataSource:any
  private refInsert:MatDialogRef<InsertBlogComponent>
  private refDelete:MatDialogRef<ConfirmDialogComponent>
  constructor(private chat:ChatServiceService, private apollo: ApolloService, private dialog: MatDialog) { }

  @ViewChild(MatPaginator, {static: true}) paginator: MatPaginator;
  @ViewChild(MatSort, {static: true}) sorter: MatSort;


  ngOnInit() {
    this.dataSource = null;
    this.getInitData()
  }
  getInitData(){
    this.allblog$ = this.apollo.selectAllBlog().subscribe(
      async res =>{
        await this.getAllblog(res)
      } 
    )
  }
  getAllblog(res:any){
    console.log(res)
    this.blogs = res.data.getallblog
    this.dataSource = new MatTableDataSource(this.blogs);
    this.dataSource.paginator = this.paginator
    this.dataSource.sort = this.sorter
    console.log(this.dataSource.data)
    
  }
  Insert(){
    this.refInsert = this.dialog.open(InsertBlogComponent)
    this.refInsert.afterClosed().subscribe(data =>{
      if(data.fromModal== true){
        this.chat.emit("blog","reload dong")
        this.getInitData()
      }
    })
  }
  Update(b:any){
    if(b.content.length == 0 || b.category.length == 0){
      alert("Blog Content And Category must be filled!")
    }
    else{
      this.apollo.UpdateBlog(b.id,b.content,b.image,b.category).subscribe()
    }
  }
  Delete(b:any){
    this.refDelete = this.dialog.open(ConfirmDialogComponent, {data:{
      message:"Are you sure you want to delete this hotel?",
      cancel: "No",
      accept: "Yes"
    }})

    this.refDelete.afterClosed().subscribe( res =>{
      if(res){
        this.apollo.DeleteBlog(b.id).subscribe(async a =>{
          await this.getInitData()
        })
      }
    })
  }

}
