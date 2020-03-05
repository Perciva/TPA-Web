
import { TextEditorComponent } from './../text-editor/text-editor.component';
import { ApolloService } from 'src/app/Services/apollo.service';
import { Component, OnInit, KeyValueDiffers, ViewChild } from '@angular/core';
import { BlogData } from 'src/app/Interfaces/blog-interface';
import { MatDialogRef, MatDialog } from '@angular/material';
import { NotificationComponent } from 'src/app/General/notification/notification.component';

@Component({
  selector: 'app-insert-blog',
  templateUrl: './insert-blog.component.html',
  styleUrls: ['./insert-blog.component.scss']
})
export class InsertBlogComponent implements OnInit {

  @ViewChild(TextEditorComponent, {static: true})
  private editor:TextEditorComponent
  private arg:string
  private image:string
  private args:string
  private refNotif:MatDialogRef<NotificationComponent>
  constructor(private apollo:ApolloService,private ref:MatDialogRef<InsertBlogComponent>,private dialog: MatDialog) { }

  ngOnInit() {
    this.arg = "Insert Blog Title"

  }

  insert(){
    console.log(this.editor.getContent())
    if (this.editor.getContent() == "" || this.editor.getTitle() == ""){
      alert("Blog Title And Content Must Not Be Empty!")
    }
    else if(this.args == "" || this.image == ""){
      alert("Please Choose Blog Category and Provide Blog Image")
    }
    else{
      var b: BlogData = {
        category: this.args,
        content: this.editor.getContent(),
        image: this.image,
        title: this.editor.getTitle(),
        id: 0,
        userID: 0,
        viewCount: 0,
      };
      this.apollo.createBlog(b).subscribe(async (a:any)=>{
        console.log(a)
        await this.validate(a.data.createblog.id)
      })
    }
  }
  validate(b:number){
    console.log(b)
    if(b == -1){
      this.refNotif = this.dialog.open(NotificationComponent,{data: "Insert Failed, Content Must be at least 10 characters long"})
    }
    else{
      this.close()
    }
  }
  close():void{
    this.ref.close({
      fromModal:true
    })
  }

}
