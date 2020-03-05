import { ChatServiceService } from './../../Services/chat-service.service';
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-chatroom',
  templateUrl: './chatroom.component.html',
  styleUrls: ['./chatroom.component.scss']
})
export class ChatroomComponent implements OnInit {

  messageControl = new FormControl();
  messageLists: Array<any> = [];
  imgSrc: string

  constructor(private  chat: ChatServiceService, private router: Router) { }

  ngOnInit() {
    if(sessionStorage.getItem('token')==null)this.router.navigate(['/']);

    this.chat.listen('chat').subscribe(m => {
      this.messageLists.push(m);
    });
  }

  validate(message: String): Boolean{
    var getsplit = message.split("|");
    //console.log(getsplit[0]);
    //console.log(getsplit[1]);
    return sessionStorage.getItem('token')+""==getsplit[1];
  }

  sendMessage() {
    var sender = sessionStorage.getItem('token');
    this.chat.emit('chat', this.messageControl.value+"|"+sender);
    this.messageControl.setValue("");
  }

  sendImage(){
    console.log(this.imgSrc);
    var sender = sessionStorage.getItem('token');
    this.chat.emit('chat', this.imgSrc+"|"+sender);
    this.messageControl.setValue("");
  }

  onFileChange(e) {
    var file = e.dataTransfer ? e.dataTransfer.files[0] : e.target.files[0];
    var pattern = /image-*/;
    var reader = new FileReader();
    if (!file.type.match(pattern)) {
      alert('invalid format');
      return;
    }
    reader.onload = this._handleReaderLoaded.bind(this);
    reader.readAsDataURL(file);
  }

  _handleReaderLoaded(e) {
    let reader = e.target;
    this.imgSrc = reader.result;
    this.sendImage();
  }
}