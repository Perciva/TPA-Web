// import { ChatServiceService } from './../../Services/chat-service.service';
// import { Component, OnInit } from '@angular/core';
// import { Router } from '@angular/router';

// @Component({
//   selector: 'app-chat',
//   templateUrl: './chat.component.html',
//   styleUrls: ['./chat.component.scss']
// })
// export class ChatComponent implements OnInit {

//   constructor(private route:Router, private chatServ:ChatServiceService) { }

//   chats:Object[]=[]
//   userId:number=-1

//   ngOnInit() {
//     this.getUserData()
//   }

//   getUserData(){
//     if(localStorage.getItem("chat-user")&&sessionStorage.getItem("User")!=null){
//       this.chatServ.getUserData(sessionStorage.getItem("User")).subscribe(async q=>{
//         // console.log(q.data.getUserData)
//         this.userId=q.data.getUserData.id
//         await this.getAllChat()
//       })
//     }
//     else if(localStorage.getItem("chat-admin")&&sessionStorage.getItem("Admin")!=null){
//       this.getAllChat()
//     }
//   }

//   getAllChat(){
//     this.chatServ.getAllChat().subscribe(async q=>{
//       this.chats=q.data.getAllChat
//       // await console.log(this.chats)
//       // await console.log(JSON.parse(`${this.chats[0].content}`))
//     })
//   }

//   backToHome(){
//     if(localStorage.getItem("chat-user")!=null&&sessionStorage.getItem("User")!=null){
//       localStorage.removeItem("chat-user")
//       this.route.navigate(['/'])
//     }
//     else if(localStorage.getItem("chat-admin")!=null&&sessionStorage.getItem("Admin")!=null){
//       localStorage.removeItem("chat-admin")
//       this.route.navigate(['/admin'])
//     }
//   }

//   validateUser1Id(idx:number){
//     return this.userId!=this.chats[idx].user1
//   }

//   validateUser2Id(idx:number){
//     return this.userId!=this.chats[idx].user2
//   }

//   openChat(idx:number){
//     var id=this.chats[idx].id
//     this.route.navigate(['/chat',id])
//   }

//   createNewChat(){
//     var friendId=prompt("Input Friend ID")
//     if(friendId==null){
//       return
//     }
//     this.chatServ.insertChat(this.userId,parseInt(friendId),"").subscribe(async q=>{
//       await alert("Success Create New Chat")
//       // console.log(q.data.insertNewChat)
//       await this.getAllChat()
//     })
//   }

//   canShow(idx:number){
//     if(this.chats[idx].user1!=this.userId&&this.chats[idx].user2!=this.userId){
//       return false
//     }
//     return true
//   }

// }
