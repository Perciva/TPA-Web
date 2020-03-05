import { Component, OnInit, Inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material';

@Component({
  selector: 'app-notification',
  templateUrl: './notification.component.html',
  styleUrls: ['./notification.component.scss']
})
export class NotificationComponent implements OnInit {
  private message:string
  constructor(
    @Inject(MAT_DIALOG_DATA) public data: any,
    private ref:MatDialogRef<NotificationComponent>
  ) { }

  ngOnInit() {
    if(this.data){
      if(this.data.length != 0 && this.data != null){
        this.message = this.data
      }
    }
  }
  close(){
    this.ref.close()
  }

}
