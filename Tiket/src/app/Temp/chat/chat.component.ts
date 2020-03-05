import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { ChatServiceService } from 'src/app/Services/chat-service.service';

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.scss']
})
export class ChatComponent implements OnInit {

  messageControl = new FormControl();
  messageLists: Array<any> = [];

  constructor(private  chatService: ChatServiceService) { }

  ngOnInit() {
    this.chatService.listen('chat').subscribe(m => {
      this.messageLists.push(m);
    });
  }

  sendMessage() {
    this.chatService.emit('chat', this.messageControl.value);
    this.messageControl.setValue("");
  }

}
