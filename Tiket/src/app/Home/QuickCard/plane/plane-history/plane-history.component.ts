import { PlaneHistoryService } from './../../../../Services/plane-history.service';
import { Component, OnInit, Output, EventEmitter } from '@angular/core';


@Component({
  selector: 'app-plane-history',
  templateUrl: './plane-history.component.html',
  styleUrls: ['./plane-history.component.scss']
})
export class PlaneHistoryComponent implements OnInit {

  @Output() data: EventEmitter<number> = new EventEmitter<number>();
  private historyData:object[];
  constructor(private hist:PlaneHistoryService) { }

  ngOnInit() {
    this.historyData = JSON.parse(this.hist.previousHistory())
    console.log(this.historyData)

  }
  passData(event){
    console.log(event);
    this.data.emit(event);
  }
  rem(){
    console.log("AAAAAaa")
    this.hist.removeHistory();
  }

}
