import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class PlaneHistoryService {

  constructor() { }

  setHistory(val:object){
    var hist: object[] = JSON.parse(this.previousHistory())
    console.log("HI")
    if(hist == null){
      hist = [val]
    }
    else{
      hist.push(val)
      console.log(hist);
    }

    sessionStorage.setItem('planeHistory', JSON.stringify(hist))
    console.log(sessionStorage.getItem('planeHistory'))
  }


  previousHistory(){
    return sessionStorage.getItem('planeHistory')
  }
}
