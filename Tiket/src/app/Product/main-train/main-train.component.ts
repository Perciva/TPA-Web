import { ChatServiceService } from './../../Services/chat-service.service';
import { MatDialog } from '@angular/material';
import { Component, OnInit } from '@angular/core';
import { NotificationComponent } from 'src/app/General/notification/notification.component';

@Component({
  selector: 'app-main-train',
  templateUrl: './main-train.component.html',
  styleUrls: ['./main-train.component.scss']
})
export class MainTrainComponent implements OnInit {
  refNotif: any;

  constructor(private dialog: MatDialog, private chat: ChatServiceService) {
    // this.getInitData()
    this.chat.listen('train').subscribe(a =>{
      this.refNotif = this.dialog.open(NotificationComponent, {data:"New Train!"})
      this.refNotif.afterClosed().subscribe(a =>{
        
        // this.getInitData()
      })
    })
   }

  ngOnInit() {
  }

}
