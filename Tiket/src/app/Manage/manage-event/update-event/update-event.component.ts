import { NgxContentLoadingModule } from 'ngx-content-loading';
import { Component, OnInit, ViewChild, Inject } from '@angular/core';
import { TextEditorComponent } from '../../manage-blog/text-editor/text-editor.component';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material';

@Component({
  selector: 'app-update-event',
  templateUrl: './update-event.component.html',
  styleUrls: ['./update-event.component.scss']
})
export class UpdateEventComponent implements OnInit {

  @ViewChild(TextEditorComponent, {static: true})
  private editor:TextEditorComponent
  private arg:string
  private image:string
  private args:string
  private error:string
  constructor(private ref:MatDialogRef<UpdateEventComponent>, 
    @Inject(MAT_DIALOG_DATA) public data: any) { }

  ngOnInit() {
    this.editor.setTitle(this.data.title)
    this.editor.setContent(this.data.desc)
  }

  done(){
    if(this.editor.getTitle() == ""){
      this.error = "Title Must Not Be Empty!"
    }
    else if(this.editor.getContent()==""){
      this.error = "Event Content must not be empty!"
    }
    else{
      this.ref.close({
        fromModal:true,
        title : this.editor.getTitle(),
        content: this.editor.getContent()
      })

    }

  }

}
