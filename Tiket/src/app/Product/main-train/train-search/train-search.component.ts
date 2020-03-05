import { MatDialog, MatDialogRef } from '@angular/material';
import { ChatServiceService } from './../../../Services/chat-service.service';
import { Component, OnInit } from '@angular/core';
import { LabelType ,Options} from 'ng5-slider';
import { Subscription } from 'rxjs';
import { ActivatedRoute, Router } from '@angular/router';
import { ApolloService } from 'src/app/Services/apollo.service';
import { subscribe } from 'graphql';
import { NotificationComponent } from 'src/app/General/notification/notification.component';

@Component({
  selector: 'app-train-search',
  templateUrl: './train-search.component.html',
  styleUrls: ['./train-search.component.scss']
})
export class TrainSearchComponent implements OnInit {

  private p:any;
  private Classes:any[] = ['Economy','Business','Executive'];
  private ClassesModel: boolean[] = [];
  private hours:string[] = ['00:00 - 06:00','06:00 - 12:00','12:00 - 18:00','18:00 - 24:00'];
  private hoursModel: boolean[] = [];
  private trainName:any[] = [];
  private trainNameModel: boolean[] = [];
  private getLocation$:Subscription
  private getTrains$:Subscription
  private trains
  private path:string;
  private SortBy:string
  refNotif: MatDialogRef<NotificationComponent>;

  constructor(
    private activatedRoute: ActivatedRoute,
    private router: Router,
    private apollo: ApolloService,
    private chat:ChatServiceService,
    private dialog:MatDialog
    ) { 
      this.chat.listen('train').subscribe(a=>{
        this.refNotif = this.dialog.open(NotificationComponent, {data:"New Train!"})
      this.refNotif.afterClosed().subscribe(a =>{
        
        this.changed()
      })
      })
      this.activatedRoute.queryParams.subscribe(async p => {
        await this.getData(p);
      });
      this.getTrain()
  }
  ngOnInit() {
    this.path = "../../../../assets/Hotel/"
    this.SortBy= "Recommended"
  }
  log(){
    console.log(this.SortBy)
  }
  goToDetail(h:any){
    console.log(h)
  }
  changed(){
    let cls:string[] = []
    let trainName:string[] = []

    for(let z=0;z<this.ClassesModel.length;z++){
      if(this.ClassesModel[z]){
        cls.push(this.Classes[z])
      }
    }
    for(let z=0;z<this.trainNameModel.length;z++){
      if(this.trainNameModel[z]){
     
        trainName.push(this.trainName[z].name)
      }
    }
    console.log(trainName,cls)

    var arr = new Date(this.p.berangkat)
    arr.setHours(arr.getHours() + 7);
    var arri = arr.toISOString().substr(0, 11) + '00:00:00Z';

    this.getTrains$ = this.apollo.getFilteredTrain(this.p.from, this.p.to,arri,trainName,cls).subscribe(async (a:any) => {
      await this.getTrainData(a.data.getfilteredtrain)
    })
  }
  resetFilter(){
    this.trainNameModel = [false]
    this.ClassesModel = [false]
    this.hoursModel = [false]

    this.changed()
    
  }
  getTrain(){
    var arr = new Date(this.p.berangkat)
    arr.setHours(arr.getHours() + 7);
    var arri = arr.toISOString().substr(0, 11) + '00:00:00Z';
    this.getTrains$ = this.apollo.selectTrainByLocation(this.p.from, this.p.to,arri).subscribe(async a => {
      await this.getTrainData(a.data.gettrainbyloc)
    })
    this.apollo.selectTrainName().subscribe(async a=>{
      await this.getTrainName(a.data.gettrainnames)
    })
  }
  getTrainName(a:any){
    this.trainName = a
    // console.log(this.trainName)
  }
  getTrainData(a:any){
    this.trains = a
    console.log(this.trains)
  }
  getData(p:any){
    console.log(p)
    // if(p.from == undefined){
    //   this.router.navigateByUrl("/Hotel")
    // }
    // else{
      this.p = p 
    // }
  }
  
  
  search(){
    var bdate:Date = new Date((<HTMLInputElement>document.getElementById('berangkat')).value);
    var pdate:Date = new Date((<HTMLInputElement>document.getElementById('pulang')).value);
    var from:string = (<HTMLInputElement>document.getElementById('from')).value
    var to:string = (<HTMLInputElement>document.getElementById('dest')).value
    var temp: String[] = ((<HTMLInputElement>document.getElementById('pclass')).value).split(' ')
    if(temp.length !=2){
      alert("Mohon isi keterangan jumlah penumpang dan Dewasa/Senior/Bayi")
      return;
    }
    var people:number = parseInt(temp[0].toString());
    var room:string = temp[1].toString();

    if(from == ""){
      alert("Mohon Isi kolom Kolom dari")
    }
    else if(new Date(bdate).toString() == "Invalid Date"|| new Date(pdate).toString() == "Invalid Date"){
      alert("Tanggal Check In dan Check Out Invalid");
    }
    else if(isNaN(people) || room == ""){
      alert("Mohon isi Keterangan penumpang dan Jenis Penumpang");
    }
    else{
      this.router.navigate(["/Train/Search/"], {
        queryParams:{
          'from': from,
          'to':to,
          'berangkat': bdate.toString(),
          'pulang': pdate.toString(),
          'people': people,
          'jenis': room
        }
      })

    }
  } 
  

}



