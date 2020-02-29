import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material';
import { Component, OnInit, Input, Inject } from '@angular/core';

@Component({
  selector: 'app-confirm-dialog',
  templateUrl: './confirm-dialog.component.html',
  styleUrls: ['./confirm-dialog.component.scss']
})
export class ConfirmDialogComponent implements OnInit {
  private message:string
  private cancel:string
  private accept:string
  
  constructor(public ref: MatDialogRef<ConfirmDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: any) { }

  ngOnInit() {
    console.log(this.data)
    this.message = this.data.message
    this.cancel = this.data.cancel
    this.accept = this.data.accept
  }
  close(t:boolean){
    this.ref.close(t)
  }
  acc(){
    this.close(true)
  }
  cance(){
    this.close(false)
  }

}
