import { PlaneHistoryService } from './../../../Services/plane-history.service';
import { Component, OnInit} from '@angular/core';
@Component({
  selector: 'app-plane',
  templateUrl: './plane.component.html',
  styleUrls: ['./plane.component.scss']
})
export class PlaneComponent implements OnInit {
  private radio:string="";
  private openHistory:boolean;
  constructor(private history: PlaneHistoryService) { }

  ngOnInit() {
    this.radio = "sj"
    this.openHistory = false;
  }


  searchPlane(){
    var bdate:Date = new Date((<HTMLInputElement>document.getElementById('berangkat')).value);
    // var pdate:Date = new Date((<HTMLInputElement>document.getElementById('pulang')).value);
    var from:string = (<HTMLInputElement>document.getElementById('from')).value
    var to:string = (<HTMLInputElement>document.getElementById('to')).value
    var pdate:Date = null
    var temp: String[] = ((<HTMLInputElement>document.getElementById('pclass')).value).split(' ')
    console.log(temp.length, temp[0], temp[1])
    if(temp.length !=2){
      alert("Mohon isi keterangan jumlah penumpang dan kelas kabin")
      return;
    }
    var people:number = parseInt(temp[0].toString());
    var clas:string = temp[1].toString();

    console.log(people)



    if(from == ""){
      alert("Mohon Isi kolom Kolom dari")
    }
    else if(to == ""){
      alert("Mohon Isi kolom Kolom Destinasi")
    }
    else if(isNaN(bdate.getFullYear())){
      alert("Mohon isi tangal keberangkatan");
    }
    else if(isNaN(people) || clas == ""){
      alert("Mohon isi Keterangan penumpang dan kelas kabin");
    }
    else{
      if(this.radio == 'pp'){
        pdate = new Date((<HTMLInputElement>document.getElementById('pulang')).value);
        if(isNaN(pdate.getFullYear())){
          alert("Mohon Isi tanggal pulang")
        }
      }
      console.log("Hey")
      let plane:object = {
        plane:{
          from: from,
          to: to,
          start: bdate,
          end:pdate,
          people: people,
          class: clas
        }
      }
      console.log(plane)
      this.history.setHistory(plane);
    }
  }  
  
  rcvData(event){
    var history = JSON.parse(this.history.previousHistory());
    console.log(history);
    console.log(event)
    var selectedHistory:any = history[event];
    console.log(selectedHistory.plane.end);
    (<HTMLInputElement>document.getElementById('berangkat')).value = selectedHistory.plane.start;
    // var pdate:Date = new Date((<HTMLInputElement>document.getElementById('pulang')).value);
    (<HTMLInputElement>document.getElementById('from')).value = selectedHistory.plane.from;
    (<HTMLInputElement>document.getElementById('to')).value = selectedHistory.plane.to;
    // var pdate:Date = null
    (<HTMLInputElement>document.getElementById('pclass')).value = selectedHistory.plane.people + " " + selectedHistory.plane.class;
    if(selectedHistory.plane.end != null){
      this.radio = "pp";
      (<HTMLInputElement>document.getElementById('pulang')).value = selectedHistory.plane.end
    }
    this.openHistory = false;
  }
}
